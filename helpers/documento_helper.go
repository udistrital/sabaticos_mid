package helpers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

// ArchivoBase64 estructura que contiene el nombre y contenido en base64 del archivo
type ArchivoBase64 struct {
	Nombre    string `json:"nombre"`
	Contenido string `json:"contenido"`
}

// ConvertirArchivoABase64 convierte un archivo multipart a base64 con su nombre
func ConvertirArchivoABase64(fileHeader *multipart.FileHeader) (*ArchivoBase64, error) {
	// Abrir el archivo
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo: %v", err)
	}
	defer file.Close()

	// Leer todo el contenido del archivo
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo: %v", err)
	}

	// Codificar a base64
	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	return &ArchivoBase64{
		Nombre:    fileHeader.Filename,
		Contenido: base64String,
	}, nil
}

// ConvertirArchivosABase64 convierte múltiples archivos a base64
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
