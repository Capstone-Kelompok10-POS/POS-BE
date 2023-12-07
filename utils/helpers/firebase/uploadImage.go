package firebase

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func UploadImageProduct(ctx echo.Context) (string, error) {

	// Set the path to your service account JSON file
	serviceAccountKeyPath := "credentials.json"

	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error initializing app: %v", err))
	}

	// Generate a new UUID
	newUUID := uuid.New()

	// Convert the UUID to a string
	uuidString := newUUID.String()

	// Set the destination path in Firebase Storage
	storagePath := "product/" + uuidString + ".png"

	// Open the uploaded file
	file, err := ctx.FormFile("image")
	if err != nil {
		return "", ctx.String(http.StatusBadRequest, fmt.Sprintf("Error reading uploaded file: %v", err))
	}

	src, err := file.Open()
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error opening uploaded file: %v", err))
	}
	defer src.Close()

	// Initialize Google Cloud Storage client
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error creating storage client: %v", err))
	}
	defer client.Close()

	// Specify the name of your Firebase Storage bucket
	bucketName := "qbills-casier.appspot.com"

	// Set the appropriate MIME type based on the file extension
	fileExtension := strings.TrimLeft(filepath.Ext(file.Filename), ".")
	var contentType string
	switch fileExtension {
	case "jpg", "jpeg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	default:
		return "", ctx.String(http.StatusBadRequest, fmt.Sprintf("Unsupported file format: %s", fileExtension))
	}

	// Upload the file to Firebase Storage with the determined content type
	object := client.Bucket(bucketName).Object(storagePath)
	wc := object.NewWriter(context.Background())
	wc.ContentType = contentType

	if _, err := io.Copy(wc, src); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error copying file to Firebase Storage: %v", err))
	}
	if err := wc.Close(); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error closing writer: %v", err))
	}

	// Set ACL for public read access after creating the object
	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error setting ACL: %v", err))
	}

	// Get the download URL
	_, err = object.Attrs(context.Background())
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file attributes: %v", err))
	}

	// Return the read-only URL to the client
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucketName, url.QueryEscape(storagePath))

	return url, nil
}
