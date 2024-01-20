package handlers

import "github.com/Rizabekus/effective-mobile-rest-api/internal/services"

type Handlers struct {
	Service *services.Services
}

func HandlersInstance(services *services.Services) *Handlers {
	return &Handlers{Service: services}
}
