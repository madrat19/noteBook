package main

import (
	"code/server"
	"code/storage"
	"log"
	"os"
)

func main() {
	// Создаем файл для логов и настраиваем вывод
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
	log.SetFlags(0)
	log.Println("--------------------------------------------------------------------------------------------")
	log.SetFlags(log.LstdFlags)

	// Создаем таблицы в бд
	storage.InitTables()

	// Предустанавливаем пользователей
	storage.CreateUser("Admin", "12345")
	storage.CreateUser("John", "54321")
	storage.CreateUser("Ivan", "qwerty")

	// Запускаем сервер
	server.RunServer()
}
