package model

const (
	UserSchema string = `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(64) NULL,
        email VARCHAR(64) NULL,
        city VARCHAR(64) NULL,
		created_at DATETIME NOT NULL
    );`
)
