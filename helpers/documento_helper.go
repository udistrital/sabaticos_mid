package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// GuardarDocumentos guarda múltiples archivos en el sistema de archivos
// Retorna las rutas de los archivos guardados y un error si ocurre
func GuardarDocumentos(files []*multipart.FileHeader, directorio string) ([]string, error) {
	var documentosGuardados []string

	// Crear directorio si no existe
	uploadDir := filepath.Join("uploads", directorio)
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		return nil, fmt.Errorf("error creando carpeta de carga: %v", err)
	}

	// Procesar cada archivo
	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			return nil, fmt.Errorf("error abriendo archivo %s: %v", header.Filename, err)
		}

		// Generar nombre único para el archivo
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(header.Filename))
		filePath := filepath.Join(uploadDir, fileName)

		// Crear archivo destino
		dst, err := os.Create(filePath)
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("error creando archivo destino: %v", err)
		}

		// Copiar contenido
		_, err = io.Copy(dst, file)
		file.Close()
		dst.Close()

		if err != nil {
			return nil, fmt.Errorf("error guardando archivo %s: %v", header.Filename, err)
		}

		documentosGuardados = append(documentosGuardados, filePath)
	}

	return documentosGuardados, nil
}
