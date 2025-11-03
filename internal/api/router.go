package api

import (
	"go-note/configs"
	"go-note/internal/note"
	"go-note/internal/todo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	config := configs.GetConfig()

	r := gin.Default()

	// CORS middleware (basic implementation)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.CORSOrigins)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

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
		api.GET("/notes/search", noteHandler.SearchNotes)
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
		api.GET("/todos/completed", todoHandler.GetCompletedTodos)
		api.GET("/todos/pending", todoHandler.GetPendingTodos)
		api.GET("/todos/search/:text", todoHandler.GetTodosByTitle)
	}

	return r
}
