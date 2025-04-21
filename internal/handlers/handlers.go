package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	//В заголовок записываем тип клиета
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//так как все успешно, то статус ОК
	w.WriteHeader(http.StatusOK)
	// выводим значение
	w.Write(data)
}

func SecondHendler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// чтение файла
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// конвертирование файла
	result, err := service.AutoConvert(string(content))
	if err != nil {
		http.Error(w, "Ошибка конвертирования", http.StatusInternalServerError)
		return
	}

	// создание имени файла
	ext := filepath.Ext(header.Filename)
	filename := strings.ReplaceAll(time.Now().UTC().String(), " ", "_")
	filename = strings.ReplaceAll(filename, ":", "-") + ext

	// создание самого файла
	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// запись результата
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Ошибка записи результата", http.StatusInternalServerError)
		return
	}

	// отправка результата пользователям
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
