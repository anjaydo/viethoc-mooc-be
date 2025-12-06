package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitDB connects to Supabase (Postgres)
func InitDB(databaseURL string) *gorm.DB {
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// We use GORM for the MVP to speed up CRUD operations.
	// In a high-perf scenario, we might switch to raw 'pgx' later.
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // Ensure table names match Supabase (e.g., "users" not "user")
		},
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to Supabase Postgres")
	return db
}
