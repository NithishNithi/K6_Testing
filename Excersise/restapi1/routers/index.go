package routers

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/api", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
