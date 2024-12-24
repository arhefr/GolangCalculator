package application

import (
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc(urlCalc, calcHandler)
	mux.HandleFunc(urlPong, pongHandler)

	http.ListenAndServe(":"+port, mux)
}
