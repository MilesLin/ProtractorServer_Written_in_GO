package controller

import (
	"ProtractorServer/model"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type login struct {
}

func (lo login) registerRoutes() {

	http.HandleFunc("/api/login", lo.login)
	http.HandleFunc("/api/login/currentIdentity", lo.GetCurrentIdentity)
	http.HandleFunc("/api/login/logout", lo.Logout)
	http.HandleFunc("/api/login/users/", lo.PutUser)
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

func (lo login) Logout(w http.ResponseWriter, r *http.Request) {
	model.CurrentIdentity = nil
	w.WriteHeader(http.StatusOK)
}

func (lo login) PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		idPattern, _ := regexp.Compile(`/api/login/users/(\d+)`)
		matches := idPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			id, _ := strconv.Atoi(matches[1])

			var body *model.Account
			// https://flaviocopes.com/golang-http-post-parameters/
			decoder := json.NewDecoder(r.Body)
			_ = decoder.Decode(&body)

			for i := range model.Accounts {
				if model.Accounts[i].Id == id {
					model.Accounts[i].LastName = body.LastName
					model.Accounts[i].FirstName = body.FirstName
				}
			}
			w.WriteHeader(http.StatusOK)
		}else{
			http.NotFound(w, r)
		}

	} else {
		http.NotFound(w, r)
	}
}
