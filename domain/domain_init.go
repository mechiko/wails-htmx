package domain

import (
	"os"
)

var Mode = "development"

// если нужно по умолчанию имя,
// используется в Config как имя файла конфиг и в repo//DbReport как имя БД
const Name = "alcogo"

func init() {
	// если установлена при компиляции то не берем из окружения
	if Mode == "" {
		// loads values from .env into the system
		// pwd, _ := os.Getwd()
		// fmt.Printf("%s\n", pwd)
		if value, exists := os.LookupEnv("MODE"); exists {
			Mode = value
		}
	}
}
