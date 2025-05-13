package controllers

import (
	"fmt"
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/services"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UploadControllerParams struct {
	fx.In
	NewsService services.NewsService
	Nginx       *config.NginxConfig
}

type UploadController struct {
	service services.NewsService
	nginx   *config.NginxConfig
}

func NewUploadController(p UploadControllerParams) *UploadController {
	return &UploadController{
		service: p.NewsService,
		nginx:   p.Nginx,
	}
}

func (n *UploadController) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo no encontrado"})
		return
	}

	folder := c.PostForm("folder")
	if folder == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'folder' es obligatorio"})
		return
	}

	uploadPath := filepath.Join("/uploads", folder)

	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la carpeta"})
			return
		}
	}

	fullPath := filepath.Join(uploadPath, file.Filename)

	if _, err := os.Stat(fullPath); err == nil {
		ext := filepath.Ext(file.Filename)
		base := file.Filename[:len(file.Filename)-len(ext)]
		newFilename := fmt.Sprintf("%s_%d%s", base, time.Now().Unix(), ext)
		fullPath = filepath.Join(uploadPath, newFilename)
	}

	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar archivo"})
		return
	}

	fullPath = filepath.Join(n.nginx.NGINX_URL, fullPath)

	c.JSON(http.StatusOK, gin.H{"message": "Archivo subido con éxito", "path": fullPath})
}

func (n *UploadController) DeleteFile(c *gin.Context) {
	filePath := c.PostForm("path")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'path' es obligatorio"})
		return
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "El archivo no existe"})
		return
	}

	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el archivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Archivo eliminado con éxito"})
}
