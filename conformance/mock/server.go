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

	http.HandleFunc("/api/customers", func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&failCount) > 0 {
			atomic.AddInt32(&failCount, -1)
			writeError(w, "internal_error", "Simulated failure", failStatus)
			return
		}

		auth := r.Header.Get("Authorization")
		if auth == "" || auth == "Bearer " {
			writeError(w, "unauthorized", "Invalid or missing API token", 401)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		page := r.URL.Query().Get("page")
		if page == "" || page == "1" {
			w.Header().Set("Link", fmt.Sprintf("<http://localhost:%d/api/customers?page=2>; rel=\"next\"", *port))
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"data": []map[string]any{
				{"id": 1, "full_name": "Jane Doe"},
				{"id": 2, "full_name": "John Smith"},
			},
		})
	})

	http.HandleFunc("/api/customers/", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || auth == "Bearer " {
			writeError(w, "unauthorized", "Invalid or missing API token", 401)
			return
		}

		id := r.URL.Path[len("/api/customers/"):]
		if id == "999999" {
			writeError(w, "not_found", "Customer not found", 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"id":        1,
				"full_name": "Jane Doe",
			},
		})
	})

	http.HandleFunc("/api/vehicles/", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || auth == "Bearer " {
			writeError(w, "unauthorized", "Invalid or missing API token", 401)
			return
		}

		id := r.URL.Path[len("/api/vehicles/"):]
		if id == "999999" {
			writeError(w, "not_found", "Vehicle not found", 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"id":    1,
				"make":  "Toyota",
				"model": "Camry",
				"year":  2020,
				"vin":   "ABC123",
			},
		})
	})

	http.HandleFunc("/api/work_orders", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || auth == "Bearer " {
			writeError(w, "unauthorized", "Invalid or missing API token", 401)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		page := r.URL.Query().Get("page")
		if page == "" || page == "1" {
			w.Header().Set("Link", fmt.Sprintf("<http://localhost:%d/api/work_orders?page=2>; rel=\"next\"", *port))
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"data": []map[string]any{
				{"id": 1, "status": "open", "customer": map[string]any{"id": 1, "full_name": "Jane Doe"}, "vehicle": map[string]any{"id": 1, "make": "Toyota", "model": "Camry"}},
			},
		})
	})

	http.HandleFunc("/api/work_orders/", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || auth == "Bearer " {
			writeError(w, "unauthorized", "Invalid or missing API token", 401)
			return
		}

		id := r.URL.Path[len("/api/work_orders/"):]
		if id == "999999" {
			writeError(w, "not_found", "Work order not found", 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"id":     1,
				"status": "open",
				"customer": map[string]any{"id": 1, "full_name": "Jane Doe"},
				"vehicle": map[string]any{"id": 1, "make": "Toyota", "model": "Camry", "year": 2020, "vin": "ABC123"},
			},
		})
	})

	fmt.Printf("Mock server listening on :%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
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
