package api

import (
	"go-note/internal/note"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize dependencies
	noteRepo := note.NewRepository(db)
	noteService := note.NewService(noteRepo)
	noteHandler := NewNoteHandler(noteService)

	// Routes
	api := r.Group("/api")
	{
		api.POST("/notes", noteHandler.CreateNote)
		api.GET("/notes", noteHandler.GetNotes)
		api.GET("/notes/:id", noteHandler.GetNote)
		api.PUT("/notes/:id", noteHandler.UpdateNote)
		api.DELETE("/notes/:id", noteHandler.DeleteNote)
	}

	return r
}
