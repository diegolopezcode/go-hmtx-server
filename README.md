# Ejemplo de Uso de HTMX en Go

Este es un ejemplo de una aplicación Go que utiliza [HTMX](https://htmx.org/) para crear aplicaciones web altamente interactivas y dinámicas sin escribir mucho código JavaScript. En este proyecto, hemos implementado un contador web simple utilizando el framework web Gin y HTMX.

## Requisitos Previos

Asegúrate de tener Go instalado en tu sistema. Puedes descargar Go desde [golang.org](https://golang.org/dl/).

## Cómo Ejecutar el Servidor

1. Clona este repositorio en tu máquina local:

   ```bash
   git clone https://github.com/tuusuario/turepositorio.git
   cd turepositorio
2. Instala las dependencias:

    ```bash
   go mod tidy
    ```
3. Ejecuta el servidor:

    ```bash
    go run main.go
    ```
4. Abre tu navegador web y navega a [http://localhost:8080/api/v1/](http://localhost:8080/api/v1/).

## Cómo Funciona la Aplicación
La aplicación es un contador web simple que muestra un número y ofrece botones para incrementar y decrementar el contador. Utiliza HTMX para actualizar el valor del contador en tiempo real sin necesidad de recargar la página completa.

La lógica del contador se encuentra en el archivo __'handler/handler.go'__. Hemos creado un único objeto __"CounterImpl"__ que se utiliza para mantener el estado del contador entre las solicitudes.

## Uso de HTMX en Go
HTMX es una poderosa herramienta que permite actualizar partes específicas de una página web sin necesidad de escribir mucho código JavaScript. En esta aplicación, hemos utilizado HTMX para enviar solicitudes al servidor Go cuando se hacen clic en los botones de incremento y decremento. El servidor Go procesa estas solicitudes y actualiza el valor del contador en la página web en tiempo real.

Puedes aprender más sobre HTMX visitando el sitio web oficial: [htmx.org](https://htmx.org/).

## Contribuciones
¡Siéntete libre de contribuir a este proyecto! Si tienes ideas para mejorar la aplicación, agregar nuevas características o solucionar problemas, no dudes en abrir un issue o enviar una solicitud de extr