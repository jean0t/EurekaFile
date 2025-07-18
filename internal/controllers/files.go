package controllers

import (
	"fmt"
	"os"
	"io"
	"time"
	"strings"
	"net/http"
	"html/template"
	"path/filepath"
	"mime/multipart"
	
	"github.com/jean0t/EurekaFile/internal/database"
	"github.com/jean0t/EurekaFile/internal/auth"

	"gorm.io/gorm"
)

const basePath string = "internal/views"

func formatDate(t time.Time) string {
	return t.Format("02 Jan 2006 15:04")
}

var templ = template.Must(template.New("").Funcs(template.FuncMap{
	"formatDate": formatDate,
}).ParseFiles(
	filepath.Join(basePath, "files.tmpl"),
	filepath.Join(basePath, "upload.tmpl"),
	filepath.Join(basePath, "navbar.tmpl"),
))

type FilesViewData struct {
	Files []database.File
}

func Files(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		db *gorm.DB
		files []database.File
		data FilesViewData
	)

	db, err = database.ConnectToDB()
	if err != nil {
		fmt.Println("[!] Error connecting to database")
		http.Error(w, "<h1>Internal Server Error</h1>", http.StatusInternalServerError)
		return
	}

	files, err  = database.GetAllFiles(db)
	if err != nil {
		fmt.Println("[!] Error fetching all files from the database")
		http.Error(w, "<h1>Internal Server Error</h1>", http.StatusInternalServerError)
		return
	}

	data = FilesViewData{Files: files}
	err = templ.ExecuteTemplate(w, "Files", data)
	if err != nil {
		fmt.Println("[!] Error executing template for /files")
		http.Error(w, "<h1>Internal Server Error</h1>", http.StatusInternalServerError)
		return
	}
}



type UploadPageData struct {
	Message string
}

func Upload(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		uploadPageData UploadPageData = UploadPageData{Message: ""}
	)

	if r.Method == http.MethodPost {
		var (
			file multipart.File
			handler *multipart.FileHeader
			saveDir string = "./uploaded_files"
			destination string
			destinationFile *os.File
			
			db *gorm.DB
			user database.User
			claims *auth.Claims = auth.GetUser(r)
		)

		db, err = database.ConnectToDB()
		if err != nil {
			fmt.Println("[!] Error connecting to database")
			http.Error(w, "<h1>Internal Server Error</h1>", http.StatusInternalServerError)
			return
		}

		user, err = database.QueryUser(db, claims.Username)
		if err != nil {
			fmt.Println("[!] Failed to query the user, can't save file to database")
			templ.ExecuteTemplate(w, "Upload", uploadPageData)
			return
		}


		err = r.ParseMultipartForm(5 << 20) // 5MB
		if err != nil {
			fmt.Println("[!] Parse Multipart gave an error")
			return
		}


		file, handler, err = r.FormFile("file")
		if err != nil {
			fmt.Println("[!] Error parsing the uploaded file")
			return
		}
		defer file.Close()


		if err = os.MkdirAll(saveDir, os.ModePerm); err != nil {
			http.Error(w, "Unable to create save directory: "+err.Error(), http.StatusInternalServerError)
			return
		}

		destination = filepath.Join(saveDir, handler.Filename)

		destinationFile, err = os.Create(destination)
		if err != nil {
			fmt.Println("[!] Error saving uploaded file")
			return
		}
		defer destinationFile.Close()


		if _, err = io.Copy(destinationFile, file); err != nil {
			fmt.Println("[!] Error coping file data to server")
			return
		}
		
		if err = database.RegisterFile(db, user, handler.Filename); err != nil {
			fmt.Println("[!] Error registering file data to database")
			uploadPageData.Message = "File Upload Failed"
			templ.ExecuteTemplate(w, "Upload", uploadPageData)
			return
		}


		uploadPageData.Message = "File Uploaded Successfully"
		templ.ExecuteTemplate(w, "Upload", uploadPageData)
		return
	}

	err = templ.ExecuteTemplate(w, "Upload", uploadPageData)
	if err != nil {
		fmt.Println("[!] Error executing template for /upload")
	}
}


func DownloadFile(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		fileName string
		filePath string
		filesDir string = "./uploaded_files"
	)

	fileName = strings.TrimPrefix(r.URL.Path, "/files/")
	if fileName == "" {
		http.Redirect(w, r, "/files", http.StatusSeeOther)
		return
	}

	fileName = filepath.Clean(fileName)
	if strings.Contains(fileName, "..") {
		http.Error(w, "Invalid File Name", http.StatusBadRequest)
		return
	}

	filePath = filepath.Join(filesDir, fileName)
	if _, err = os.Stat(filePath); err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
	http.Redirect(w, r, "/files", http.StatusSeeOther)

}
