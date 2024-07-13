# revision_cumplidos_proveedores_crud

Repositorio de revisión cumplidos proveedores crud que permite gestionar la revisión, aprobación o rechazo de los cumplidos


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# parametros de api
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_HTTP_PORT=[Puerto de exposición del API]
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGUSER=[Usuario de BD]
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGADMIN=[Contraseña del usaurio de BD]
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_GPORT=[Puerto de la BD]
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGDB=[Nombre de Base de Datos]
REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con CUMPLIDOS_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/revision_cumplidos_proveedores_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/revision_cumplidos_proveedores_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
CUMPLIDOS_CRUD_PORT=8080 CUMPLIDOS_CRUD_SOME_VARIABLE bee run


### Ejecución Dockerfile
```shell
# docker build --tag=cumplidos_crud . --no-cache
# docker run -p 80:80 cumplidos_crud
```

### Ejecución docker-compose
```shell
# 1. Obtener repositorio
git clone https://github.com/udistrital/revision_cumplidos_proveedores_crud.git
# 2. Ir a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/revision_cumplidos_proveedores_crud
# 3. Cambiar a la rama develop
git checkout develop
# 4. Crear red back_end
docker network create back_end
# 5. Ejecutar docker compose
docker-compose up --build --remove-orphans


### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/revision_cumplidos_proveedores_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/revision_cumplidos_proveedores_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/revision_cumplidos_proveedores_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/revision_cumplidos_proveedores_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/revision_cumplidos_proveedores_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/revision_cumplidos_proveedores_crud) |


## Licencia

This file is part of cumplidos_crud.

cumplidos_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cumplidos_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.