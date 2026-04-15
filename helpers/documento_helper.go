package helpers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

// ArchivoBase64 structure that contains the file name and base64-encoded content
type ArchivoBase64 struct {
	Nombre    string `json:"nombre"`
	Contenido string `json:"contenido"`
}

// ConvertirArchivoABase64 converts a multipart file to base64 with its name
func ConvertirArchivoABase64(fileHeader *multipart.FileHeader) (*ArchivoBase64, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the entire file content
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Encode to base64
	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	return &ArchivoBase64{
		Nombre:    fileHeader.Filename,
		Contenido: base64String,
	}, nil
}

// ConvertirArchivosABase64 converts multiple files to base64
func ConvertirArchivosABase64(files []*multipart.FileHeader) ([]ArchivoBase64, error) {
	var archivosBase64 []ArchivoBase64

	for _, fileHeader := range files {
		archivo, err := ConvertirArchivoABase64(fileHeader)
		if err != nil {
			return nil, err
		}
		archivosBase64 = append(archivosBase64, *archivo)
	}

	return archivosBase64, nil
}
