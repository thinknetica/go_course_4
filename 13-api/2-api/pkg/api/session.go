package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type authInfo struct {
	Usr string
	Pwd string
}

func (api *API) authSession(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var auth authInfo
	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if auth.Usr == "Usr" && auth.Pwd == "Pwd" {
		session, _ := api.store.Get(r, "session-cookie")
		session.Values["Usr"] = auth.Usr
		session.Values["Pwd"] = auth.Pwd
		session.Values["authenticated"] = true
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
