package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID		uuid.UUID `json:"id"`
	Name	string	  `json:"name"`
}

type Server struct {
	*mux.Router

	shoppingList []Item // storing data right here because it's just an example
}

func NewServer() *Server {
	s := &Server {
		Router: mux.NewRouter(),
		shoppingList: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/list", s.listShoppingItems()).Methods("GET")
	s.HandleFunc("/list", s.createShoppingItem()).Methods("POST")
	s.HandleFunc("/list/{id}", s.removeShoppingItem()).Methods("DELETE")
}

func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i.ID = uuid.New()
		s.shoppingList = append(s.shoppingList, i)

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr   := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingList {
			if item.ID == id {
				s.shoppingList = append(s.shoppingList[:i], s.shoppingList[i + 1:]...)
				break
			}
		}
	}
}