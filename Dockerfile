FROM golang:1.23-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./

RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación, especificando la ruta correcta para main.go
RUN go build -o main cmd/main.go

# Exponer el puerto en el que la aplicación escucha
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
