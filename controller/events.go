package controller

import (
	"ProtractorServer/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
)

type events struct{}

func (e events) registerRoutes() {
	http.HandleFunc("/api/events", e.GetAndAddEvent)
	http.HandleFunc("/api/events/", e.GetAndPutEvent)
}

func (e events) GetAndAddEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		result, _ := json.Marshal(model.MyEvents)
		w.Write(result)
	case http.MethodPost:
		addEvent(w, r)
	}

}

func (e events) GetAndPutEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idPattern, _ := regexp.Compile(`/api/events/(\d+)`)
	matches := idPattern.FindStringSubmatch(r.URL.Path)

	if len(matches) < 1 {
		http.NotFound(w, r)
		return
	}
	id, _ := strconv.Atoi(matches[1])

	if r.Method == http.MethodGet {

		for i := range model.MyEvents {
			if model.MyEvents[i].Id == id {
				result, _ := json.Marshal(&model.MyEvents[i])
				w.Write(result)
				return
			}
		}

	} else if r.Method == http.MethodPut {
		for i := range model.MyEvents {
			if model.MyEvents[i].Id == id {
				var eventBody model.Event
				decoder := json.NewDecoder(r.Body)
				decoder.Decode(&eventBody)
				model.MyEvents[i].Session = eventBody.Session
				return
			}
		}
	}
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) //10MB
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	base64String := base64.StdEncoding.EncodeToString(fileBytes)
	resultImage := fmt.Sprintf("data:image/png;base64,%s", base64String)
	price, _ := strconv.ParseFloat(r.Form.Get("price"), 32)
	resultEvent := model.Event{
		Id:    getNextEventId(),
		Name:  r.Form.Get("name"),
		Date:  r.Form.Get("date"),
		Time:  r.Form.Get("time"),
		Price: float32(price),
		Location: &model.Location{
			Address: r.Form.Get("location.address"),
			City:    r.Form.Get("location.city"),
			Country: r.Form.Get("location.country"),
		},
		Session:   &[]model.Session{},
		OnlineUrl: r.Form.Get("onlineUrl"),
		Image:     resultImage,
	}
	model.MyEvents = append(model.MyEvents, resultEvent)

}

func getNextEventId() int {
	sort.Slice(model.MyEvents, eventSort)
	maxId := model.MyEvents[len(model.MyEvents)-1].Id
	return maxId + 1
}

func eventSort(i int, j int) bool {
	return model.MyEvents[i].Id < model.MyEvents[j].Id
}
