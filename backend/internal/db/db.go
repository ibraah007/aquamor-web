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
	DB, err = sql.Open("sqlite3", "./aquamor.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

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