package v1

import (
	"net/http"
	v1 "online-library/src/v1"
)

func Routes() {
	http.HandleFunc("/api/v1/file", v1.FileHandler)   //Upload & downnload & delete
	http.HandleFunc("/api/v1/query", v1.QueryHandler) //get filemeta data from system
}
