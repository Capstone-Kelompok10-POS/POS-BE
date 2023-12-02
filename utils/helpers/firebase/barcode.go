package firebase

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"image/png"
	"io"
	"net/http"
)

func GenerateBarcodeAndUploadToFirebase(ctx echo.Context, CodeMember string) (string, error) {
	// Generate the barcode
	barcodeEncode, err := code128.Encode(CodeMember)
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error generating barcode: %v", err))
	}
	barcode, err := barcode.Scale(barcodeEncode, 500, 150)
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error scaling barcode: %v", err))
	}

	// Encode the barcode image to PNG format and buffer it
	var buf bytes.Buffer
	if err := png.Encode(&buf, barcode); err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error encoding barcode to buffer: %v", err))
	}

	// Set the path to your service account JSON file
	serviceAccountKeyPath := "credentials.json"

	// Initialize Google Cloud Storage client
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error creating storage client: %v", err))
	}
	defer client.Close()

	// Specify the name of your Firebase Storage bucket
	bucketName := "qbils-d46b3.appspot.com"

	// Set the appropriate MIME type
	contentType := "image/png"

	// Set the destination path in Firebase Storage
	storagePath := "barcode/" + CodeMember + ".png"

	// Upload the file to Firebase Storage with the determined content type
	object := client.Bucket(bucketName).Object(storagePath)
	wc := object.NewWriter(context.Background())
	wc.ContentType = contentType

	if _, err := io.Copy(wc, &buf); err != nil {
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
	attrs, err := object.Attrs(context.Background())
	if err != nil {
		return "", ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file attributes: %v", err))
	}

	// Return the read-only URL to the client
	url := attrs.MediaLink

	return url, nil
}
