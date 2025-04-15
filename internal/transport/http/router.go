package http

import (
	"net/http"

	"github.com/FrancoBarrera99/auth-service/internal/auth"
	"github.com/gorilla/mux"
)

type Router struct {
	serv    auth.AuthService
	handler Handler
}

func NewRouter(service auth.AuthService, handler Handler) *Router {
	return &Router{
		serv:    service,
		handler: handler,
	}
}

func (r *Router) Init() http.Handler {
	mux := mux.NewRouter()
	//mux.Use(loggingMiddleware)
	mux.HandleFunc("/register", r.handler.Register).Methods("POST")
	mux.HandleFunc("/login", r.handler.Login).Methods("POST")
	//mux.HandleFunc("/validate", r.serv.ValidateToken).Methods("POST")
	//mux.HandleFunc("/health", healthHandler).Methods("GET")
	return mux
}
