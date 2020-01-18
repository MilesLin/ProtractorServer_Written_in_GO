package controller

import (
	"ProtractorServer/model"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type login struct {
}

func (lo login) registerRoutes() {

	http.HandleFunc("/api/login", lo.login)
	http.HandleFunc("/api/login/currentIdentity", lo.GetCurrentIdentity)
}

func (lo login) login(w http.ResponseWriter, r *http.Request) {

	var dbAct *model.Account

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		var body *model.Account

		// https://flaviocopes.com/golang-http-post-parameters/
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&body)

		if err != nil {
			log.Println(err)
		}

		for _, account := range model.Accounts {
			if strings.ToLower(account.Username) == strings.ToLower(body.Username) &&
				strings.ToLower(account.Password) == strings.ToLower(body.Password) {
				dbAct = &account
			}
		}

		if dbAct == nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			model.CurrentIdentity = dbAct
			w.Header().Add("Content-Type", "application/json")
			result, _ := json.Marshal(dbAct)
			w.Write(result)
		}

	} else {
		http.NotFound(w, r)
	}

}

func (lo login) GetCurrentIdentity(w http.ResponseWriter, r *http.Request) {
	if model.CurrentIdentity == nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.Header().Add("Content-Type", "application/json")
		result, _ := json.Marshal(model.CurrentIdentity)
		w.Write(result)
	}
}
