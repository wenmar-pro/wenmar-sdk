package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync/atomic"
)

var (
	failCount  int32
	failStatus int
)

func main() {
	port := flag.Int("port", 18080, "port to listen on")
	failCountFlag := flag.Int("fail-count", 0, "number of times to fail before succeeding")
	flag.IntVar(&failStatus, "fail-status", 500, "status code to fail with")
	flag.Parse()
	atomic.StoreInt32(&failCount, int32(*failCountFlag))

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}

		switch r.Method {
		case http.MethodGet:
			if atomic.LoadInt32(&failCount) > 0 {
				atomic.AddInt32(&failCount, -1)
				writeError(w, "internal_error", "Simulated failure", failStatus)
				return
			}
			page := r.URL.Query().Get("page")
			if page == "" || page == "1" {
				w.Header().Set("Link", fmt.Sprintf("<http://localhost:%d/customers?page=2>; rel=\"next\"", *port))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode([]map[string]any{
				{"id": 1, "full_name": "Jane Doe"},
				{"id": 2, "full_name": "John Smith"},
			})
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(map[string]any{"id": 1, "full_name": "Jane Doe"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/customers/", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}

		id := r.URL.Path[len("/customers/"):]
		if id == "999999" {
			writeError(w, "not_found", "Customer not found", 404)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{"id": 1, "full_name": "Jane Doe"})
		case http.MethodPatch:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{"id": 1, "full_name": "Jane Doe"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/vehicles", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}

		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode([]map[string]any{
				{"id": 1, "make": "Toyota", "model": "Camry", "year": 2020, "vin": "ABC123"},
				{"id": 2, "make": "Honda", "model": "Civic", "year": 2018, "vin": "XYZ789"},
			})
			return
		}

		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(map[string]any{"id": 3, "make": "Honda", "model": "Civic", "year": 2020})
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/vehicles/vin_decode", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{"make": "Honda", "model": "Civic", "vin": "1HGCM82633A004352"})
	})

	http.HandleFunc("/vehicles/check_duplicate", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"matches": []map[string]any{
				{"id": 1, "display_name": "Toyota Camry", "url": "/vehicles/1", "reasons": []string{"vin"}},
			},
		})
	})

	http.HandleFunc("/vehicles/", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}

		id := r.URL.Path[len("/vehicles/"):]
		if id == "999999" {
			writeError(w, "not_found", "Vehicle not found", 404)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{
				"id":    1,
				"make":  "Toyota",
				"model": "Camry",
				"year":  2020,
				"vin":   "ABC123",
			})
		case http.MethodPatch:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{"id": 1, "make": "Toyota", "model": "Camry", "year": 2020})
		case http.MethodDelete:
			w.WriteHeader(204)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/work_orders", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}

		if r.Method == http.MethodGet {
			page := r.URL.Query().Get("page")
			if page == "" || page == "1" {
				w.Header().Set("Link", fmt.Sprintf("<http://localhost:%d/work_orders?page=2>; rel=\"next\"", *port))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode([]map[string]any{
				{"id": 1, "status": "open"},
			})
			return
		}

		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(map[string]any{"id": 3, "status": "open"})
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/work_orders/", func(w http.ResponseWriter, r *http.Request) {
		if !requireAuth(w, r) {
			return
		}

		id := r.URL.Path[len("/work_orders/"):]
		if id == "999999" {
			writeError(w, "not_found", "Work order not found", 404)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{"id": 1, "status": "open"})
		case http.MethodPatch:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{"id": 1, "status": "in_progress"})
		case http.MethodDelete:
			w.WriteHeader(204)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Mock server listening on :%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func requireAuth(w http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" || auth == "Bearer " {
		writeError(w, "unauthorized", "Invalid or missing API token", 401)
		return false
	}
	return true
}

func writeError(w http.ResponseWriter, code, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"error": map[string]any{
			"code":    code,
			"message": message,
			"details": map[string]any{},
		},
	})
}
