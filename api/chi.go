package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	Service service.Service
}

func New(srv service.Service) http.Handler {
	return Handler(API{Service: srv})
}

func Handler(a API) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/users", func(r chi.Router) {
		r.Post("/", a.Insert)
		r.Get("/{id}", a.ById)
		r.Put("/{id}", a.Update)
		r.Delete("/{id}", a.Delete)
	})
	return r
}

// ById gets a user info addressing to the database
// and responses with the requested data AND row id
func (a API) ById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := a.Service.Db.User(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResp(w, http.StatusOK, respBytes)
}

// Insert parses a request data then sends it to the database.
// and responses with the stored data with row's id
func (a API) Insert(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.Service.Db.Insert(context.Background(), &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uBytes, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResp(w, http.StatusCreated, uBytes)

}

// Update replaces the existing data to the requested one
// and responses with the new data
func (a API) Update(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.ID = int64(id)

	resp, err := a.Service.Db.Update(context.Background(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResp(w, http.StatusOK, respBytes)
}

// Delete removes a row from a database
func (a API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.Service.Db.Delete(context.Background(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResp(w, http.StatusOK, nil)
}

func writeResp(w http.ResponseWriter, status int, data []byte) {
	if data == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
