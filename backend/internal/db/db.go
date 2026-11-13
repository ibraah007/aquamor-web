package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB opens the SQLite database at path, checks the connection and creates
// the tables if they do not exist. The caller owns the process lifetime; tests
// and the server can call Close when they are done.
func InitDB(path string) error {
	if path == "" {
		return errors.New("database path is empty")
	}

	var err error
	DB, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?_busy_timeout=5000&_journal_mode=WAL&_foreign_keys=on", path))
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	// SQLite allows one writer at a time. A single connection avoids
	// "database is locked" errors in this small service.
	DB.SetMaxOpenConns(1)

	if err = DB.Ping(); err != nil {
		DB.Close()
		DB = nil
		return fmt.Errorf("connect to database: %w", err)
	}

	if err = createTables(); err != nil {
		DB.Close()
		DB = nil
		return err
	}
	return nil
}

// Ping checks that the database is reachable. It is used by the health check.
func Ping(ctx context.Context) error {
	if DB == nil {
		return errors.New("database is not initialised")
	}
	return DB.PingContext(ctx)
}

// Close releases the database connection.
func Close() error {
	if DB == nil {
		return nil
	}
	err := DB.Close()
	DB = nil
	return err
}

func createTables() error {
	const schemaQuery = `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_name TEXT NOT NULL,
		customer_phone TEXT NOT NULL,
		customer_email TEXT,
		delivery_estate TEXT NOT NULL,
		delivery_address TEXT NOT NULL,
		total_amount REAL NOT NULL,
		order_status TEXT DEFAULT 'pending',
		payment_status TEXT DEFAULT 'unpaid',
		payment_provider TEXT,
		payment_reference TEXT,
		mpesa_receipt_number TEXT,
		paid_at DATETIME,
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

	if _, err := DB.Exec(schemaQuery); err != nil {
		return fmt.Errorf("create database tables: %w", err)
	}
	if err := migrate(); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}
	return nil
}

// migrate adds columns that were introduced after the first release. It is
// safe to run against both new and existing databases.
func migrate() error {
	columns := []struct {
		name string
		ddl  string
	}{
		{"customer_email", "ALTER TABLE orders ADD COLUMN customer_email TEXT"},
		{"payment_provider", "ALTER TABLE orders ADD COLUMN payment_provider TEXT"},
		{"payment_reference", "ALTER TABLE orders ADD COLUMN payment_reference TEXT"},
		{"mpesa_receipt_number", "ALTER TABLE orders ADD COLUMN mpesa_receipt_number TEXT"},
		{"paid_at", "ALTER TABLE orders ADD COLUMN paid_at DATETIME"},
	}
	for _, column := range columns {
		exists, err := columnExists("orders", column.name)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		if _, err := DB.Exec(column.ddl); err != nil {
			return fmt.Errorf("add column %s: %w", column.name, err)
		}
	}

	// Older rows may have an empty string instead of NULL; the unique index
	// only excludes NULL, so normalise those rows first.
	if _, err := DB.Exec(`UPDATE orders SET payment_reference = NULL WHERE payment_reference = ''`); err != nil {
		return fmt.Errorf("normalize payment references: %w", err)
	}
	if _, err := DB.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_orders_payment_reference ON orders(payment_reference) WHERE payment_reference IS NOT NULL`); err != nil {
		return fmt.Errorf("create payment reference index: %w", err)
	}
	return nil
}

func columnExists(table, column string) (bool, error) {
	rows, err := DB.Query("PRAGMA table_info(" + table + ")")
	if err != nil {
		return false, fmt.Errorf("read table info: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			cid       int
			name      string
			ctype     string
			notnull   int
			dfltValue interface{}
			pk        int
		)
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk); err != nil {
			return false, fmt.Errorf("scan table info: %w", err)
		}
		if name == column {
			return true, nil
		}
	}
	return false, rows.Err()
}
