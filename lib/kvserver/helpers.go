package kvserver

import "net/http"

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getID(r *http.Request) string {
	if id := r.Header.Get("X-Forwarded-For"); id != "" {
		return id
	}

	return r.Host
}
