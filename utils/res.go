package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, payload interface{}, status int) {

	mapdata := map[string]interface{}{}
	var (
		ok = http.StatusOK
		created = http.StatusCreated
		internalservererror = http.StatusInternalServerError
		notfound = http.StatusNotFound
		badrequest = http.StatusBadRequest
		methodnotallowed = http.StatusMethodNotAllowed
	)

	switch status {

	case methodnotallowed:
		mapdata["status"] = methodnotallowed
		mapdata["message"] = "Method Error"
		mapdata["data"] = payload

	case badrequest:
		mapdata["status"] = badrequest
		mapdata["message"] = "request error"
		mapdata["data"] = payload

	case notfound:
		mapdata["status"] = notfound
		mapdata["message"] = "data not found"
		mapdata["data"] = payload

	case internalservererror:
		mapdata["status"] = internalservererror
		mapdata["message"] = "failed"
		mapdata["data"] = payload

	case created:
		mapdata["status"] = created
		mapdata["message"] = "success"
		mapdata["data"] = payload

	default:
		mapdata["status"] = ok
		mapdata["message"] = "success"
		mapdata["data"] = payload
	}


	response, err := json.Marshal(mapdata)


	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(response)

}
