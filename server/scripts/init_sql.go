package scripts

func GetPGInitScript() string {
	return `CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username varchar(50) NOT NULL,
            email varchar(255) UNIQUE NOT NULL,
            password varchar(999) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            CONSTRAINT not_empty_username CHECK (username <> ''),
            CONSTRAINT not_empty_email CHECK (email <> '')
          )
  `
}

func GetSQLiteInitScript() string {
	return `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            CONSTRAINT not_empty_username CHECK (username <> ''),
            CONSTRAINT not_empty_email CHECK (email <> '')
          )
  `
}
