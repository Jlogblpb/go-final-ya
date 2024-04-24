package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// создаем новый роутер
	r := chi.NewRouter()

	// регистрируем в роутере эндпоинт `/` с методом getIndex
	r.Get("/", getIndex)

	// регистрируем в роутере эндпоинт `/artists` с методом POST,
	// для которого используется обработчик `postArtist`
	r.Post("/artists", postArtist)
	// регестрируем в роутере эндпоинт `/artist/{id}`  с методом GETб,
	// для которого используется обработчик `getArtist`
	r.Get("/artist/{id}", getArtist)

	// Регистрируем роутер Русский рок
	r.Get("/russian_rock/{id}", getRusArtists)

	// запускаем сервер
	if err := http.ListenAndServe(":7540", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
