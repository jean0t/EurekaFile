# EurekaFile

EurekaFile is a simple file server web application written in Go. It allows users to upload and download files securely, with user authentication and session management. The application provides a clean Bootstrap-based UI and is structured following the MVC (Model-View-Controller) pattern.

## Features

- **User Accounts & Sessions:**  
  Users can log in to access the service. New users are automatically created upon their first login. Sessions are managed to ensure security and privacy.

- **File Upload & Download:**  
  Authenticated users can upload files, view a list of all uploaded files, and download them.

- **Responsive UI:**  
  The web interface uses Bootstrap for a modern and responsive appearance.

## Supported Routes

- `/login`  
  Displays the login page (username & password). If the user does not exist, they are created; otherwise, the user is logged in. Sessions are created upon successful login.

- `/upload`  
  Presents the file upload form. Only available to authenticated users. Allows users to select and upload files.

- `/files`  
  Shows a table of all uploaded files, including file name, author, upload date, and a download button.

- `/logout`  
  Logs the user out and destroys their session.

## Project Structure

```
internal
├── controllers      # HTTP handlers (controllers) for each route
├── database         # Database/model logic for users and files
├── middleware       # Custom middleware, e.g. logging
├── router           # Route registration
└── views            # HTML templates with Bootstrap styling
```

## Getting Started

1. **Clone the repository:**
    ```sh
    git clone github.com/jean0t/EurekaFile # or user JMFern01
    cd EurekaFile
    ```

2. **Build and run:**
    ```sh
    go build -o eurekafile cmd/main.go
    ./eurekafile -s
    ```

3. **Open in your browser:**  
   Visit [http://localhost:8080/login](http://localhost:8080/login) to start.

## License

> MIT

