package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Priority represents the priority level of a task
type Priority string

const (
	PriorityLow    Priority = "L"
	PriorityMedium Priority = "M"
	PriorityHigh   Priority = "H"
)

// Task represents a todo item
type Task struct {
	ID       int
	UserID   int
	Text     string
	Priority Priority
}

// DB represents the database connection
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection
func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return nil, err
	}

	// Drop existing tables to ensure schema is up to date
	// _, err = db.Exec(`
	// 	DROP TABLE IF EXISTS todos;
	// 	DROP TABLE IF EXISTS users;
	// `)
	// if err != nil {
	// 	return nil, err
	// }

	// Create tables if they don't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT UNIQUE,
			password TEXT
		);
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY,
			user_id INTEGER,
			task TEXT,
			priority TEXT CHECK(priority IN ('L', 'M', 'H')) DEFAULT 'M',
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		return nil, err
	}

	// Insert default user if not exists
	_, err = db.Exec(`
		INSERT OR IGNORE INTO users (username, password) 
		VALUES ('ratnesh', 'ratnesh')
	`)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// AddTask adds a new task for a user
func (db *DB) AddTask(userID int, task string, priority Priority) error {
	_, err := db.Exec("INSERT INTO todos (user_id, task, priority) VALUES (?, ?, ?)",
		userID, task, string(priority))
	return err
}

// UpdateTask updates an existing task
func (db *DB) UpdateTask(taskID int, task string, priority Priority) error {
	_, err := db.Exec("UPDATE todos SET task = ?, priority = ? WHERE id = ?",
		task, string(priority), taskID)
	return err
}

// DeleteTask removes a task
func (db *DB) DeleteTask(taskID int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", taskID)
	return err
}

// GetTasks returns all tasks for a user
func (db *DB) GetTasks(userID int) ([]Task, error) {
	rows, err := db.Query("SELECT id, task, priority FROM todos WHERE user_id = ? ORDER BY id DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var priority string
		if err := rows.Scan(&task.ID, &task.Text, &priority); err != nil {
			return nil, err
		}
		task.UserID = userID
		task.Priority = Priority(priority)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// VerifyUser checks if the username and password match
func (db *DB) VerifyUser(username, password string) (int, bool) {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?",
		username, password).Scan(&id)
	if err != nil {
		log.Println("Login error:", err)
		return 0, false
	}
	return id, true
}
