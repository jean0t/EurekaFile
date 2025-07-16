package controllers

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"html/template"
	"path/filepath"
	"mime/multipart"
	
	"github.com/jean0t/EurekaFile/internal/database"

	"gorm.io/gorm"
)

const basePath string = "internal/views"

var templ = template.Must(template.ParseFiles(
	filepath.Join(basePath, "files.tmpl"),
	filepath.Join(basePath, "upload.tmpl"),
	filepath.Join(basePath, "navbar.tmpl"),
))

func Files(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		db *gorm.DB
		files []database.File
	)

	db, err = database.ConnectToDB()
	if err != nil {
		fmt.Println("[!] Error connecting to database")
	}

	files, err  = database.GetAllFiles(db)
	if err != nil {
		fmt.Println("[!] Error fetching all files from the database")
	}

	err = templ.ExecuteTemplate(w, "Files", files)
	if err != nil {
		fmt.Println("[!] Error executing template for /files")
	}
}


type UploadPageData struct {
	Message string
}

func Upload(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		file multipart.File
		handler *multipart.FileHeader
		saveDir string
		destination string
		destinationFile *os.File
		uploadPageData UploadPageData = UploadPageData{Message: ""}
	)

	if r.Method == http.MethodPost {
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

		saveDir = "./uploaded_files"
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

		uploadPageData.Message = "File Uploaded Successfully"
		templ.ExecuteTemplate(w, "Upload", uploadPageData)
		return
	}

	templ.ExecuteTemplate(w, "Upload", uploadPageData)
	if err != nil {
		fmt.Println("[!] Error executing template for /upload")
	}
}
