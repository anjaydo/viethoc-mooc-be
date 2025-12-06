package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"mooc-platform/internal/config"
	"mooc-platform/internal/middleware"
	"mooc-platform/pkg/database"
	// In a real modular app, you would import module routes here
	// "mooc-platform/internal/modules/content"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database (Supabase)
	db := database.InitDB(cfg.DatabaseURL)
	// Pass 'db' to your modules later...
	// Log the database connection
	log.Println("Database connected successfully")
	// log first row of the database
	row := db.Table("users").Limit(1).Row()
	var user string
	if err := row.Scan(&user); err != nil {
		log.Fatalf("Failed to get row: %v", err)
	}
	log.Println(user)

	// 3. Setup Router (Gin)
	r := gin.Default()

	// 4. Global Middleware
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// 5. API Routes Group
	api := r.Group("/api/v1")
	{
		// Health Check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "db": "connected"})
		})

		// --- PUBLIC ROUTES (No Auth) ---
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Login implementation pending"})
			})
		}

		// --- PROTECTED ROUTES (Requires Tenant ID) ---
		// We apply the Tenant Middleware to this group
		protected := api.Group("/")
		protected.Use(middleware.RequireTenant())
		{
			// Example: Content Module Routes
			content := protected.Group("/courses")
			{
				content.GET("", func(c *gin.Context) {
					// Example of accessing the Tenant ID set by middleware
					tenantID, _ := c.Get(middleware.TenantIDKey)

					// In a real handler, you would pass 'db' and 'tenantID' to a Service
					c.JSON(http.StatusOK, gin.H{
						"module":    "content",
						"tenant_id": tenantID,
						"data":      "List of courses for this tenant",
					})
				})

				content.POST("", func(c *gin.Context) {
					c.JSON(http.StatusCreated, gin.H{"message": "Course created"})
				})
			}

			// Example: Assessment Module Routes (Manual Grading)
			grading := protected.Group("/assessments")
			{
				grading.POST("/submit", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Exam submitted for manual grading"})
				})
			}
		}
	}

	// 6. Start Server
	log.Printf("🚀 Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
