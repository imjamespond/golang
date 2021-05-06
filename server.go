package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func main() {
	manager := manage.NewDefaultManager()
	// token memory store, POSIBLE LEAD TO MEMORY LEAKING!
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	client := models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:8080",
	}
	fmt.Println(
		"Authorization: Basic " + base64.StdEncoding.EncodeToString([]byte(client.GetID()+":"+client.GetSecret())))
	clientStore := store.NewClientStore()
	clientStore.Set(client.GetID(), &client)
	manager.MapClientStorage(clientStore)
	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp: 30 * time.Second,
		// IsGenerateRefresh: false,
	})
	manager.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp: 9000 * time.Second,
	})
	// manager.SetAuthorizeCodeExp(10 * time.Second)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// FIXME fake userId for auth-code
	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return "123", nil
	})

	// case oauth2.PasswordCredentials
	// srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {

	http.HandleFunc("/authorize-code", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		token, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"expires_in": int64(time.Until(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn())).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	log.Fatal(http.ListenAndServe(":8001", nil))
}
