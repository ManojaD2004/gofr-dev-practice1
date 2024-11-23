package googledrive

import (
    "context"
    "fmt"
    "log"
    "os"
    "google.golang.org/api/drive/v3"
    "google.golang.org/api/option"
)

func Googledrive(CredentialsPath string, UploadFilePath string ) {
     
    serviceAccount := CredentialsPath
    filename := UploadFilePath

    if serviceAccount == "" || filename == "" {
        log.Fatal("SERVICE_ACCOUNT_JSON or UPLOAD_FILE_PATH not set in .env file")
    }
    ctx := context.Background()
    srv, err := drive.NewService(ctx, option.WithCredentialsFile(serviceAccount), option.WithScopes(drive.DriveScope))
    if err != nil {
        log.Fatalf("Unable to create Drive client: %v", err)
    }
    file, err := os.Open(filename)
    if err != nil {
        log.Fatalf("Unable to open file: %v", err)
    }
    defer file.Close()
    info, err := file.Stat()
    if err != nil {
        log.Fatalf("Unable to get file info: %v", err)
    }
    f := &drive.File{
        Name:    info.Name(),
        MimeType: "text/plain", 
   
    }

    uploadedFile, err := srv.Files.Create(f).
        Media(file).
        ProgressUpdater(func(now, size int64) { fmt.Printf("Upload progress: %d/%d bytes\r", now, size) }).
        Do()
    if err != nil {
        log.Fatalf("Unable to upload file: %v", err)
    }

    permission := &drive.Permission{
        Type: "anyone",
        Role: "reader",
    }
    _, err = srv.Permissions.Create(uploadedFile.Id, permission).Do()
    if err != nil {
        log.Fatalf("Failed to set file permissions: %v", err)
    }


    shareableLink := fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", uploadedFile.Id)
    fmt.Printf("File uploaded successfully! Access it here: %s\n", shareableLink)
}
