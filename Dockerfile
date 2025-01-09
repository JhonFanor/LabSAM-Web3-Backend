# Etapa de compilación
FROM golang:1.23-alpine as builder

# Instalar herramientas necesarias (opcional)
RUN apk add --no-cache git

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de go.mod y go.sum para aprovechar la caché de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación, especificando la ruta correcta para main.go
RUN go build -o main cmd/main.go

# Etapa final: imagen ligera para producción
FROM alpine:latest

# Establecer el directorio de trabajo
WORKDIR /root/

# Copiar el binario compilado desde la etapa de construcción
COPY --from=builder /app/main .

# Copiar el archivo .env (si es necesario)
COPY .env ./

# Exponer el puerto en el que la aplicación escucha
EXPOSE 8080

# Configurar permisos para ejecutar el binario
RUN chmod +x ./main

# Comando para ejecutar la aplicación
CMD ["./main"]

