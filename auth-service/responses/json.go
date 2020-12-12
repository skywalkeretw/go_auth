package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err = fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Println("Error encoding JSON: ", err)
		}
	}
}