# sabaticos_mid
:heavy_check_mark: Check: Repositorio API MID para el sistema de sabáticos.


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
# parametros de api
APPNAME=[Nombre de la aplicación]
SABATICOS_MID_HTTPPORT=[Puerto de exposición del API]
SABATICOS_MID_RUNMODE=[Modo de ejecución: dev/prod]
SABATICOS_MID_AUTORENDER=[Habilitar autorender]
SABATICOS_MID_COPYREQUESTBODY=[Copiar cuerpo de la petición]
SABATICOS_MID_ENABLEDOCS=[Habilitar documentación]
SABATICOS_MID_ENABLEXSRF=[Habilitar protección XSRF]
SABATICOS_MID_XRAY=[Habilitar AWS X-Ray]

# servicios consumidos
SABATICOS_MID_TERCEROS=[URL del API de terceros]
SABATICOS_MID_SABATICOS_CRUD=[URL del API CRUD de sabáticos]
SABATICOS_MID_GESTORDOCUMENTAL=[URL del API de gestor documental]
SABATICOS_MID_DOCUMENTOS_CRUD=[URL del API CRUD de documentos]
SABATICOS_MID_ACADEMICA_JBPM=[URL del servicio académica JBPM]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas acorde a los lineamientos de definición de variables.


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/sabaticos_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/sabaticos_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
APPNAME=api_mid_sabaticos SABATICOS_MID_HTTPPORT=8081 SABATICOS_MID_RUNMODE=dev bee run
```

### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/sabaticos_mid

#2. Moverse a la carpeta del repositorio
cd sabaticos_mid

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas
```shell
# Ejecutar todas las pruebas del proyecto
go test ./...
```

Pruebas unitarias
```shell
# Ejecutar las pruebas unitarias de los servicios
go test ./tests/services/...

# Ejecutar una prueba específica por nombre
go test ./tests/services/... -run TestCambiarEstado

# Ejecutar las pruebas con reporte de cobertura
go test ./tests/services/... -cover
```
## Estado CI


| Develop | Relese 0.0.1 | Master | Sonar |
| -- | -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_mid/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_mid/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_mid) | [![Quality Gate Status](https://sonar.portaloas.udistrital.edu.co/api/project_badges/measure?project=sabaticos_mid&metric=alert_status&token=sqb_8cd46b629cded43e15ff3052983c0dccfc91bca8)](https://sonar.portaloas.udistrital.edu.co/dashboard?id=sabaticos_mid) |


## Licencia

This file is part of sabaticos_mid.

sabaticos_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

sabaticos_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with sabaticos_mid. If not, see https://www.gnu.org/licenses/.
