package api

import (
	"go-note/internal/note"
	"go-note/internal/todo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize dependencies
	noteRepo := note.NewRepository(db)
	noteService := note.NewService(noteRepo)
	noteHandler := NewNoteHandler(noteService)

	todoRepo := todo.NewRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := NewTodoHandler(todoService)

	// Routes
	api := r.Group("/api")
	{
		// Note routes
		api.POST("/notes", noteHandler.CreateNote)
		api.GET("/notes", noteHandler.GetNotes)
		api.GET("/notes/:id", noteHandler.GetNote)
		api.PUT("/notes/:id", noteHandler.UpdateNote)
		api.DELETE("/notes/:id", noteHandler.DeleteNote)

		// Todo routes
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos", todoHandler.GetTodos)
		api.GET("/todos/:id", todoHandler.GetTodo)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
		api.PATCH("/todos/:id/toggle", todoHandler.ToggleTodoComplete)
	}

	return r
}
