package models

type EstadoSabatico struct {
	Id                int    `json:"Id"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
	NombreEstado      string `json:"NombreEstado"`
	Activo            bool   `json:"Activo"`
}
