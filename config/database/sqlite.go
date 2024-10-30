package database

import (
	"database/sql"

	"github.com/ahmetilboga2004/go-blog/pkg/utils"
	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "blog.db")
	if err != nil {
		utils.Log(utils.ERROR, "Database connection failed: %v", err)
	}
	if err = db.Ping(); err != nil {
		utils.Log(utils.ERROR, "Database ping failed: %v", err)
	}

	utils.Log(utils.INFO, "Database connection successfully")
	if err := createTables(db); err != nil {
		utils.Log(utils.ERROR, "Table creation failed: %v", err)
	}
	return db
}

func createTables(db *sql.DB) error {
	statementes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id BLOB PRIMARY KEY,
			firstName TEXT NOT NULL,
			lastName TEXT NOT NULL,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			salt TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			id BLOB PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			user_id BLOB,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
		`CREATE TABLE IF NOT EXISTS comments (
			id BLOB PRIMARY KEY,
			content TEXT,
			post_id BLOB,
			user_id BLOB,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
	}

	for _, stmt := range statementes {
		_, err := db.Exec(stmt)
		if err != nil {
			utils.Log(utils.ERROR, "Failed to execute statement: %s, error: %v", stmt, err)
			return err
		}
	}

	utils.Log(utils.INFO, "All tables created successfully")
	return nil
}
