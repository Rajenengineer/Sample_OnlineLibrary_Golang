package v1

import (
	"encoding/json"
	"log"
	"net/http"
	v1 "online-library/repo/v1"
)

//QueryHandler handles file related queries
func QueryHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		data, err := v1.GetFileInfo(req.FormValue("fileName"))
		if err != nil {
			log.Println("Failed to fetch the files , err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		byteData, err := json.Marshal(data)
		if err != nil {
			log.Println("Failed to marshal the data , err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(byteData)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
