package controller

import (
	"ProtractorServer/model"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type login struct{}

func (lo login) registerRoutes() {

	http.HandleFunc("/api/login", lo.login)
	http.HandleFunc("/api/login/currentIdentity", lo.GetCurrentIdentity)
	http.HandleFunc("/api/login/logout", lo.Logout)
	http.HandleFunc("/api/login/users/", lo.PutUser)
	http.HandleFunc("/api/login/new", lo.NewUser)
}

func (lo login) login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var dbAct *model.Account

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
			result, _ := json.Marshal(dbAct)
			w.Write(result)
		}

	}
}

func (lo login) GetCurrentIdentity(w http.ResponseWriter, r *http.Request) {
	if model.CurrentIdentity == nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
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
		} else {
			http.NotFound(w, r)
		}

	}
}

func (lo login) NewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sort.Slice(model.Accounts, accountSort)
		maxId := model.Accounts[len(model.Accounts)-1].Id

		var body *model.Account
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&body)

		model.Accounts = append(model.Accounts, model.Account{
			maxId + 1,
			body.FirstName,
			body.LastName,
			body.Username,
			body.Password,
		})
	}
}

func accountSort(i int, j int) bool {
	return model.Accounts[i].Id < model.Accounts[j].Id
}
