package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/KONICCO/go-industrious-boilerplate/internal/config"
	"github.com/KONICCO/go-industrious-boilerplate/internal/handler"
	"github.com/KONICCO/go-industrious-boilerplate/internal/repository"
	"github.com/KONICCO/go-industrious-boilerplate/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Run() error {
	cfg := config.New()

	// Initialize PostgreSQL connection
	// dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Test the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	repo := repository.NewPostgresRepository(db)
	svc := service.New(repo)
	h := handler.New(svc)

	r := gin.Default()
	h.RegisterRoutes(r)

	log.Printf("Server is running on %s", cfg.Server.Address)
	return r.Run(cfg.Server.Address)
}
