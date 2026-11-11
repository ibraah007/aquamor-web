package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB opens the SQLite database connection and creates required tables
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "file:./aquamor.db?_busy_timeout=5000&_journal_mode=WAL&_foreign_keys=on")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// SQLite allows one writer at a time. A single connection avoids
	// "database is locked" errors in this small service.
	DB.SetMaxOpenConns(1)

	// Test connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to SQLite database successfully!")
	createTables()
}

func createTables() {
	schemaQuery := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_name TEXT NOT NULL,
		customer_phone TEXT NOT NULL,
		delivery_estate TEXT NOT NULL,
		delivery_address TEXT NOT NULL,
		total_amount REAL NOT NULL,
		order_status TEXT DEFAULT 'pending',
		payment_status TEXT DEFAULT 'unpaid',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS order_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id INTEGER,
		product_id INTEGER,
		bottle_option TEXT,
		quantity INTEGER,
		unit_price REAL,
		subtotal REAL,
		FOREIGN KEY(order_id) REFERENCES orders(id)
	);
	`

	_, err := DB.Exec(schemaQuery)
	if err != nil {
		log.Fatalf("Error creating database tables: %v", err)
	}

	fmt.Println("Database tables verified and ready.")
}
