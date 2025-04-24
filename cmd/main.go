package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// создаем логгер
	logger := log.New(os.Stdout, "MORSE:", log.LstdFlags|log.Lshortfile)

	// создаем сервер
	srv := server.NewServer(logger)

	// запускаем его
	logger.Println("Старт сервера на http://localhost:8080")
	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal("ошибка запуска сервер:", err)
	}

}
