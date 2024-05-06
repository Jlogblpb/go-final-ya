package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Создаем БД
func createDB() {
	//создаем непосредственно БД
	fileDB, err := sql.Open("sqlite3", "./scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	//создаем таблицу
	new_table, err := fileDB.Prepare("CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date CHAR(8), title TEXT, comment TEXT, repeat VARCHAR(128))")
	if err != nil {
		log.Fatal(err)
	}
	new_table.Exec()
	//Сделать date индексируемой
	crateIndex, err := fileDB.Prepare("CREATE INDEX indexdate ON scheduler (date)")
	crateIndex.Exec()
	fileDB.Close()
}

func main() {

	// Работа с БД

	//Проверяем БД на наличие
	appPath, err := os.Getwd()
	//Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(appPath, "scheduler.db")
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX

	//Создаем БД
	if install == true {
		createDB()
	} else {
		db, err := sql.Open("sqlite3", "./scheduler.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}

	//Запуск WEB сервера
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.Handle("/js/scripts.min.js", http.FileServer(http.Dir(webDir)))
	http.Handle("/css/style.css", http.FileServer(http.Dir(webDir)))
	http.Handle("/favicon.ico", http.FileServer(http.Dir(webDir)))

	err = http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}

}
