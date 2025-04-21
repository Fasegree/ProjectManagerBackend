package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite" // Используем драйвер без CGO
)

var DB *sqlx.DB

func InitDatabase() {
	var err error

	if err := os.MkdirAll("./data", os.ModePerm); err != nil {
		log.Fatalf("Error creating 'data' directory: %v", err)
	}

	// ✅ Заменено "sqlite3" на "sqlite"
	DB, err = sqlx.Open("sqlite", "./data/myapp.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		field TEXT NOT NULL,
		description TEXT,
		deadline TEXT NOT NULL,
		experience TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS vacancies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		field TEXT,
		country TEXT,
		experience TEXT,
		FOREIGN KEY (project_id) REFERENCES projects(id)
	);`

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	log.Println("Database initialized with projects and vacancies tables")
}

func CloseDatabase() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing the database: %v", err)
	}
	log.Println("Database connection closed")
}
