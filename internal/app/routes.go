package app

import (
	"log"
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/handlers"
	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()

	r.HandleFunc("/people", h.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", h.GetPerson).Methods("GET")
	r.HandleFunc("/people", h.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", h.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", h.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// добавить инъекцию зависимостей ///////////////////////////
// расписать обработчики
// заполнить область с логикой сервисов
// заполнить часть с операциями с БД
// добавить миграции
// добавить .env
// покрыть дебагами и логами
// добавить в пкг чужие апи для обогащения
