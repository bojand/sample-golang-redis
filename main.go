package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	defaultRedisURL = "localhost:6379"
	defaultPort     = "80"
)

// Declare a pool variable to hold the pool of Redis connections.
var pool *redis.Pool

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = defaultRedisURL
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize a connection pool and assign it to the pool global
	// variable.
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisURL)
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/counter", handleCount)

	log.Println("Listening on :4000...")
	http.ListenAndServe(":4000", mux)
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	c := pool.Get()
	defer c.Close()

	cmd := "GET"

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)

		return
	}
	if r.Method == http.MethodPost {
		cmd = "INCR"
	}

	n, err := redis.Int(c.Do(cmd, "k1"))
	if err != nil {
		http.Error(w, http.StatusText(500)+": "+err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "%+v\n", n)
}
