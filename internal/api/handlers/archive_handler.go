package handlers

import (
	"net/http"
	"rbp/internal/models"
	"rbp/internal/service"
	"rbp/pkg/validator"

	"github.com/gin-gonic/gin"
)

type ArchiveHandler struct {
	archiveService *service.ArchiveService
}

func NewArchiveHandler(archiveService *service.ArchiveService) *ArchiveHandler {
	return &ArchiveHandler{archiveService: archiveService}
}

func (h *ArchiveHandler) AddToArchive(c *gin.Context) {
	var archive models.Archive
	if err := c.ShouldBindJSON(&archive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.Validate(archive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.archiveService.AddToArchive(c, &archive); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added to archive successfully"})
}

func (h *ArchiveHandler) GetArchivesByUserID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	archives, err := h.archiveService.GetArchivesByUserID(c, userID.(int32))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, archives)
}
