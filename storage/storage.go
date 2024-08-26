package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Создает нужные таблицы в бд
func InitTables() error {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return err
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS users (
					id SERIAL PRIMARY KEY,
					username VARCHAR(255) UNIQUE NOT NULL,
					password VARCHAR(255) NOT NULL,
					api_key VARCHAR(255)
					);`

	_, err = db.Exec(query)
	if err != nil {
		log.Println("Database query error: ", err)
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		note TEXT
		);`

	_, err = db.Exec(query)
	if err != nil {
		log.Println("Database query error: ", err)
		return err
	}

	log.Println("Tables created successfully")
	return nil
}

// Создает пользователя
func CreateUser(username, password string) error {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%s', '%s')", username, password)
	_, err = db.Exec(query)
	if err != nil {
		log.Println("Database query error: ", err)
		return err
	}

	log.Printf("User <%s> created successfully\n", username)
	return nil
}

// Создает/обновляет api-ключ
func UpdateApiKey(username, password, apiKey string) error {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE users SET api_key = '%s' WHERE username = '%s'and password = '%s'", apiKey, username, password)
	_, err = db.Exec(query)
	if err != nil {
		log.Println("Database query error: ", err)
		return err
	}

	log.Printf("Api-key for <%s> updated successfully\n", username)
	return nil
}

// Аутентфкаця
func Authentication(username, password string) (bool, error) {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return false, err
	}
	defer db.Close()

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE username = $1 AND password = $2
		);
	`
	var ok bool
	err = db.QueryRow(query, username, password).Scan(&ok)
	if err != nil {
		log.Println("Database query error: ", err)
		return false, err
	}
	if ok {
		log.Printf("<%s> identified successfully\n", username)
	} else {
		log.Printf("Failed to identify <%s>", username)
	}
	return ok, nil
}

// Авторизация
func Authorization(apiKey string) (int, error) {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return 0, err
	}
	defer db.Close()

	query := `
		SELECT id, username
		FROM users
		WHERE api_key = $1;
	`

	var userID int
	var username string
	err = db.QueryRow(query, apiKey).Scan(&userID, &username)
	if err != nil {
		log.Println("Database query error: ", err)
		return 0, err
	}

	log.Printf("<%s> authorized successfully\n", username)
	return userID, nil
}

// Добавляет заметку
func AddNote(userID int, note string) error {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO notes (user_id, note) VALUES ('%d', '%s')", userID, note)
	_, err = db.Exec(query)
	if err != nil {
		log.Println("Database query error: ", err)
		return err
	}

	query = fmt.Sprintf("SELECT username FROM users where id = %d", userID)
	var username string
	err = db.QueryRow(query).Scan(&username)
	if err != nil {
		log.Println("Database query error: ", err)
		return err
	}

	log.Printf("Note for <%s> add successfully\n", username)
	return nil

}

// Получает список заметок
func GetNotes(userID int) ([]string, error) {
	db, err := sql.Open("postgres", auth)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT note FROM notes WHERE user_id = %d", userID)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Database query error: ", err)
		return nil, err
	}

	var notes []string
	for rows.Next() {
		var note string
		err := rows.Scan(&note)
		if err != nil {
			log.Println("Failed to scan notes: ", err)
			return nil, err
		}
		notes = append(notes, note)
	}

	query = fmt.Sprintf("SELECT username FROM users where id = %d", userID)
	var username string
	err = db.QueryRow(query).Scan(&username)
	if err != nil {
		log.Println("Database query error: ", err)
		return nil, err
	}

	log.Printf("Get notes for <%s> successfully\n", username)
	return notes, nil

}
