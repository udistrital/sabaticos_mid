package models

import "encoding/xml"

type SecretariaAcademicaResponse struct {
	XMLName xml.Name `xml:"secretaria"`
	Persona *Persona `xml:"persona"`
}

type Persona struct {
	Apellidos         string `xml:"apellidos"`
	Estado            string `xml:"estado"`
	Identificacion    string `xml:"identificacion"`
	Dependencia       string `xml:"dependencia"`
	CodigoDependencia string `xml:"codigo_dependencia"`
	Nombres           string `xml:"nombres"`
}
