package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/qjcg/kit"
)

func main() {
	conf, err := kit.NewConfig("http2stream", "config")
	if err != nil {
		log.Fatal("[INFO] Error reading config file:", err)
	}

	debug := conf.GetBool("debug")
	kit.NewLevelledLogger(debug)

	log.Println("[DEBUG] THIS IS A DEBUG MESSAGE!")

	cert := conf.GetString("cert")
	key := conf.GetString("key")
	if cert == "" || key == "" {
		log.Fatal("[INFO] Invalid TLS cert and/or key provided!")
	}

	http.HandleFunc("/", Handler)
	log.Println("[INFO] Listening on https://0.0.0.0:9999/")
	log.Fatal(http.ListenAndServeTLS(":9999", cert, key, nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing request form!", http.StatusInternalServerError)
		return
	}

	// Timeout defaults to 1 second. A timeout value can also be provided
	// via POST request in Milliseconds.
	timeout := time.Second
	if r.Method == "POST" {
		if v := r.Form.Get("timeout"); v != "" {
			t, err := strconv.Atoi(v)
			if err != nil {
				http.Error(w, "Error parsing POSTED timeout value!", http.StatusBadRequest)
				return
			}
			timeout = time.Duration(t) * time.Millisecond
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	var resp Response
	for {
		select {
		case <-ctx.Done():
			return

		default:
			resp.Time = time.Now()
			if err := enc.Encode(&resp); err != nil {
				http.Error(w, "Error encoding data!", http.StatusInternalServerError)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}
}

type Response struct {
	Time time.Time
}
