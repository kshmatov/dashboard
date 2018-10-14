package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kshmatov/dashboard/log"
	"github.com/kshmatov/dashboard/model"
)

func GetLastN() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feeds, err := model.GetLastN(100)
		if err != nil {
			http.Error(w, "This is an error", http.StatusBadRequest)
			log.Printf("On GetLast100: %v", err)
			return
		}

		js, err := json.Marshal(people)
		if err != nil {
			http.Error(w, "This is an error", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, string(js))
	})
}