package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// isMorse проверяет, является ли строка кодом Морзе
func isMorse(s string) bool {
	for _, r := range s {
		if !(r == '.' || r == '-' || r == ' ' || r == '/') {
			return false
		}
	}
	return true
}

// AutoConvert определяет тип входной строки, конвертируем и возвращаем результат
func AutoConvert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("строка input пустая")
	}

	if isMorse(input) {
		// конвертируем из морзе в текст
		return morse.ToText(input), nil
	}
	// конвертируем из текста в морзе
	return morse.ToMorse(input), nil
}
