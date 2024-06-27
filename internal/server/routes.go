package server

import (
	"local-storage-server/internal/database"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Serve the static website that will communicate to the API
	r.Use(static.Serve("/", static.LocalFile("./internal/frontend", true)))

	// db methods
	r.GET("/health", s.healthHandler)

	// API methods
	r.GET("/api/files", s.listFilesHandler)
	r.GET("/api/files/:id", s.getFileHandler)
	r.POST("/api/files", s.uploadFileHandler)
	r.DELETE("/api/files/:id", s.deleteFileHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) uploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is nil"})
		return
	} else {
		var path string
		path, err = filepath.Abs("store")

		path = filepath.Join(path, file.Filename)

		toAdd := database.FileToAdd{
			Name: file.Filename,
			Path: path,
			Type: file.Header.Get("Content-Type"),
			Size: file.Size,
		}

		err = s.db.AddFile(toAdd)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}

}

func (s *Server) getFileHandler(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := s.db.GetFile(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
		return
	}

	c.FileAttachment(file.Path, file.Name)
	c.JSON(http.StatusOK, gin.H{})
	return
}

func (s *Server) deleteFileHandler(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := s.db.GetFile(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
	}

	err = s.db.DeleteFile(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
	}

	err = os.RemoveAll(file.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (s *Server) listFilesHandler(c *gin.Context) {
	files, err := s.db.ListFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}
