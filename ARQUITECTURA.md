# Arquitectura del Proyecto - Sabáticos Mid

## Descripción General

`sabaticos_mid` es un API middleware desarrollado en Go utilizando el framework Beego. Este servicio actúa como una capa intermedia entre el frontend y los servicios CRUD de backend, orquestando la lógica de negocio para la gestión de solicitudes de año sabático en la Universidad Distrital.

## Estructura de Carpetas

### 📁 `/clients`

**Propósito**: Contiene los clientes HTTP que se comunican con servicios externos (CRUD y otros microservicios).

Esta capa actúa como una abstracción de las llamadas HTTP a servicios externos, encapsulando la lógica de comunicación y transformación de datos. Cada cliente se especializa en interactuar con un recurso específico del sistema.

**Archivos principales**:
- **`solicitud_client.go`**: Cliente principal para operaciones CRUD de solicitudes. Maneja la creación y gestión del recurso principal del sistema.
- **`tercero_client.go`**: Cliente para validar y consultar información de terceros (personas que realizan solicitudes)
- **`historial_solicitud_client.go`**: Cliente para registrar y consultar el historial de cambios y transiciones de estado de las solicitudes
- Y otros clientes para recursos complementarios del sistema

**Características importantes**:
- **Transformación de datos**: Los clientes convierten los modelos internos a las estructuras esperadas por los servicios CRUD
- **Referencias de ID**: Las entidades relacionadas se envían como objetos con estructura `{ "Id": valor }` en lugar de IDs simples, permitiendo que el backend procese correctamente las relaciones
- **Manejo de errores**: Captura y propaga errores de comunicación HTTP de manera consistente
- **Abstracción**: Ocultan los detalles de implementación de las llamadas HTTP al resto de la aplicación

**Patrón de Diseño**: Repository Pattern - cada cliente encapsula la comunicación HTTP con un servicio específico, proporcionando una interfaz limpia para el resto de la aplicación y facilitando el testing mediante mock.

---

### 📁 `/conf`

**Propósito**: Configuración de la aplicación.

**Archivos**:
- **`app.conf`**: Archivo de configuración principal que define:
  - Puerto HTTP de la aplicación
  - Modo de ejecución (dev/prod)
  - URLs de servicios externos (terceroService, sabaticosService)
  - Configuraciones de seguridad (XSRF)
  - Configuración de documentación (Swagger)

---

### 📁 `/controllers`

**Propósito**: Capa de controladores que maneja las peticiones HTTP entrantes y devuelve respuestas.

**Archivos**:
- **`solicitud_controller.go`**: Controlador principal que expone endpoints REST para:
  - `POST /v1/solicitud/` - Crear nueva solicitud
  - `GET /v1/solicitud/` - Obtener todas las solicitudes
  - `GET /v1/solicitud/:id` - Obtener solicitud por ID
  - `PUT /v1/solicitud/:id` - Actualizar solicitud
  - `DELETE /v1/solicitud/:id` - Eliminar solicitud

**Responsabilidades**:
- Validar datos de entrada
- Llamar a la capa de servicio
- Formatear respuestas HTTP
- Manejo de errores HTTP

---

### 📁 `/enums`

**Propósito**: Define constantes y enumeraciones para valores predefinidos del sistema.

**Archivos**:
- **`estado_solicitud.enum.go`**: Estados posibles de una solicitud (ENVIADA, APROBADA, RECHAZADA, etc.)
- **`tipo_solicitud.enum.go`**: Tipos de solicitud disponibles
- **`tipo_tercero.enum.go`**: Tipos de terceros (docente, estudiante, etc.)

**Ventajas**: 
- Previene errores de typos en strings
- Facilita refactorización
- Documenta valores válidos del sistema

---

### 📁 `/helpers`

**Propósito**: Funciones auxiliares reutilizables en toda la aplicación.

**Archivos**:
- **`response_helper.go`**: Funciones para formatear respuestas HTTP de manera consistente
- **`values_helper.go`**: Funciones para extraer y transformar datos de las respuestas de APIs externas

**Funciones típicas**:
- Formateo de respuestas de éxito/error
- Extracción de datos de respuestas JSON
- Transformaciones de datos comunes

---

### 📁 `/models`

**Propósito**: Define las estructuras de datos (DTOs y entidades) utilizadas en la aplicación.

Los modelos se dividen en **cuatro tipos principales**, cada uno con un propósito específico en el flujo de datos:

#### 1. **Modelos Base** (sufijo `_model.go`):
Representan las entidades de base de datos tal como están almacenadas en el sistema CRUD. Se utilizan principalmente para **lectura** y mapeo directo de respuestas del backend.

**Características**:
- Incluyen tags ORM para mapeo con base de datos
- Usan tipos primitivos para IDs (int, string)
- JSON tags en `snake_case` (ejemplo: `tercero_id`)
- Se usan en operaciones GET y respuestas del CRUD

**Ejemplos**: `solicitud_model.go`, `historial_solicitud_model.go`, `sabatico_model.go`

---

#### 2. **Modelos Request** (sufijo `_request.model.go`):
Estructuras DTO que **reciben datos desde el frontend**. Validan y transforman la entrada del usuario antes de procesarla en los servicios.

**Características**:
- JSON tags en `PascalCase` o `snake_case` según convención del frontend
- Validaciones de entrada
- Campos opcionales u obligatorios según la operación
- Se usan en el body de peticiones POST/PUT desde el cliente web

**Ejemplo**: `solicitud_request.model.go` - recibe la solicitud inicial del usuario con datos básicos.

---

#### 3. **Modelos CreateRequest** (sufijo `_create_request.model.go`):
Estructuras DTO especializadas para **enviar datos de creación al CRUD**. La diferencia clave con los modelos base es que utilizan referencias de ID como objetos en lugar de IDs simples.

**Características**:
- JSON tags en `PascalCase`
- **Relaciones como objetos**: Usan `IdReference` en lugar de `int` para foreign keys
- Estructura esperada por el backend CRUD para procesar correctamente las relaciones
- Solo contienen campos necesarios para creación (sin IDs autogenerados, fechas de sistema, etc.)

**Ejemplo**:
```go
type SolicitudCreateRequest struct {
    TerceroId       int         `json:"TerceroId"`
    Activo          bool        `json:"Activo"`
    TipoSolicitudId IdReference `json:"TipoSolicitudId"` // Objeto, no int
}
```

**¿Por qué existen?**: El servicio CRUD backend requiere que las relaciones se envíen como objetos `{ "Id": valor }` para procesarlas correctamente en las transacciones de base de datos.

---

#### 4. **Modelos Response** (sufijo `_response.model.go`):
Estructuras DTO que **envían datos procesados al frontend**. Combinan y transforman información de múltiples fuentes para presentar una respuesta completa al cliente.

**Características**:
- Agregación de datos de múltiples modelos
- Formato optimizado para el frontend
- Pueden incluir campos calculados o derivados
- Excluyen información sensible o innecesaria

**Ejemplo**: `solicitud_response.model.go` - puede incluir datos de la solicitud + información del tercero + estado actual.

---

#### Modelos Auxiliares:
- **`id_reference.go`**: Estructura especial para representar referencias de ID como objetos:
  ```go
  type IdReference struct {
      Id int `json:"Id"`
  }
  ```
  Utilizado en todos los `CreateRequest` para campos relacionales.

---

**Flujo de Transformación de Datos**:
```
Frontend → Request → Service → CreateRequest → CRUD Backend
                                     ↓
Frontend ← Response ← Service ← Modelo Base ← CRUD Backend
```

**Ventajas de esta arquitectura**:
- **Separación de responsabilidades**: Cada tipo de modelo tiene un propósito claro
- **Flexibilidad**: Permite transformar datos sin afectar las entidades de DB
- **Validación**: Cada capa puede validar datos en su formato específico
- **Compatibilidad**: Los CreateRequest adaptan los datos al formato esperado por el CRUD

---

### 📁 `/routers`

**Propósito**: Define las rutas HTTP y las mapea a los controladores correspondientes.

**Archivos**:
- **`router.go`**: Configuración central de rutas usando el sistema de namespaces de Beego
  - Versión de API: `/v1`
  - Namespaces: `/user`, `/solicitud`, `/sabatico`

**Características**:
- Versionado de API
- Agrupación lógica de endpoints
- Documentación automática con Swagger

---

### 📁 `/service`

**Propósito**: Capa de lógica de negocio que orquesta operaciones complejas.

**Archivos**:
- **`solicitud_service.go`**: Servicio principal que orquesta la creación de solicitudes
  - Valida tercero
  - Crea solicitud
  - Registra historial y formulario en paralelo usando goroutines
  - Maneja transacciones lógicas entre múltiples servicios

**Responsabilidades**:
- Orquestación de múltiples llamadas a clientes
- Lógica de negocio compleja
- Validaciones de negocio
- Manejo de concurrencia

**Patrón**: Service Layer - separa la lógica de negocio de los controladores

---

## Convenciones de Nombres

- **Archivos**: `snake_case` con sufijos descriptivos (`_client.go`, `_model.go`, `_service.go`)
- **Tipos**: `PascalCase` (`SolicitudController`, `IdReference`)
- **Funciones**: `PascalCase` para exportadas, `camelCase` para privadas
- **JSON Tags**: 
  - Modelos de DB: `snake_case` (`tercero_id`)
  - Modelos de Request: `PascalCase` (`TerceroId`)

---