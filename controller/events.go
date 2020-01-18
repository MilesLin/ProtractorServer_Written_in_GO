package controller

import (
	"ProtractorServer/model"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

type events struct{}

func (e events) registerRoutes() {
	http.HandleFunc("/api/events", e.Get)
	http.HandleFunc("/api/events/", e.GetAnEvent)
}

func (e events) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")
		result, _ := json.Marshal(model.MyEvents)
		w.Write(result)
	} else {
		http.NotFound(w, r)
	}
}

func (e events) GetAnEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		idPattern, _ := regexp.Compile(`/api/events/(\d+)`)
		matches := idPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			id, _ := strconv.Atoi(matches[1])

			var event *model.Event

			for i := range model.MyEvents {
				if model.MyEvents[i].Id == id {
					event = &model.MyEvents[i]
				}
			}
			w.WriteHeader(http.StatusOK)
			if event != nil {
				result, _ := json.Marshal(event)
				w.Write(result)
			}

		} else {
			http.NotFound(w, r)
		}

	} else {
		http.NotFound(w, r)
	}
}
