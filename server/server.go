package server

import (
	"code/speller"
	"code/storage"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

// Запускает сервер
func RunServer() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/notes", notesHandler)
	log.Println("Server is running")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Println("Server is stopped")
	log.Fatal(err)
}

// Отправление и получение заметок
func notesHandler(writer http.ResponseWriter, request *http.Request) {
	// Отправление заметок
	if request.Method == "POST" {
		apiKey := request.Header.Get("api-key")
		if apiKey == "" {
			http.Error(writer, "api-key header missing", http.StatusUnauthorized)
			return
		}

		// Проверка авторизации
		userID, err := storage.Authorization(apiKey)
		if err != nil {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Failed to read body", http.StatusInternalServerError)
		}
		note := string(body)

		if note == "" {
			http.Error(writer, "No body provided", http.StatusBadRequest)
			return
		}
		// Исправление орфографии
		note, err = speller.Spell(note)
		if err != nil {
			http.Error(writer, "Failed to check spelling", http.StatusInternalServerError)
		}

		// Сохранение заметки
		err = storage.AddNote(userID, note)
		if err != nil {
			http.Error(writer, "Failed to save note", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(200)
		writer.Write([]byte("Note saved"))

		// Получение заметок
	} else if request.Method == "GET" {
		apiKey := request.Header.Get("api-key")
		if apiKey == "" {
			http.Error(writer, "api-key header missing", http.StatusUnauthorized)
			return
		}

		// Проверка авторизации
		userID, err := storage.Authorization(apiKey)
		if err != nil {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Получение заметок пользователя
		notes, err := storage.GetNotes(userID)
		if err != nil {
			http.Error(writer, "Failed to get notes", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(notes)

	} else {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Авторизация
func authHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		username, password, ok := request.BasicAuth()
		if !ok {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ok, err := storage.Authentication(username, password)
		if !ok || err != nil {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		apiKey := genApiKey()

		err = storage.UpdateApiKey(username, password, apiKey)
		if err != nil {
			http.Error(writer, "Failed to update apiKey", http.StatusInternalServerError)
		}

		type apiKeyResponse struct {
			APIKey string `json:"api-key"`
		}

		response := apiKeyResponse{APIKey: apiKey}
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	} else {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

// Создает случайный api-ключ
func genApiKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 20)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
