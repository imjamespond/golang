package jwtserver

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/generates"
)

const ServerPort = 8001

var (
	EcdsaPublicKey *ecdsa.PublicKey
	ecdsaPublicKey *ecdsa.PublicKey
	// ecdsaPrivateKey *ecdsa.PrivateKey
	privateKey []byte
	publicKey  []byte
)

func init() {
	var err error
	privateKey, err = ioutil.ReadFile("./keypair/ec256-private.pem")
	if err != nil {
		panic(err.Error())
	}
	// if ecdsaPrivateKey, err = jwt.ParseECPrivateKeyFromPEM(privateKey); err != nil {
	// 	panic(err.Error())
	// }
	publicKey, err = ioutil.ReadFile("./keypair/ec256-public.pem")
	if err != nil {
		panic(err.Error())
	}
	if ecdsaPublicKey, err = jwt.ParseECPublicKeyFromPEM(publicKey); err != nil {
		panic(err.Error())
	}
	EcdsaPublicKey = ecdsaPublicKey
}

func Start() {
	// store auth code...
	manager := manage.NewDefaultManager()
	// manager.MustTokenStorage(NewDummyTokenStore())
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("SignedKeyID", []byte("SignedKey"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("SignedKeyID", privateKey, jwt.SigningMethodES256))

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
		AccessTokenExp: 60 * time.Second,
		// IsGenerateRefresh: false,
	})
	manager.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp:    9000 * time.Second,
		IsGenerateRefresh: true,
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

	// In case requred by auth-code
	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		log.Println("UserAuthorizationHandle")
		return r.URL.Query().Get("userId"), nil
	})

	// In case oauth2.PasswordCredentials（只有password才能取得useId并放入jwt token中）
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		return "123456", nil
	})

	http.HandleFunc("/authorize-code", func(w http.ResponseWriter, r *http.Request) {
		// verify client domain by clientId, redirect_uri
		// srv.ValidationAuthorizeRequest(r) 读出 RedirectURI...of AuthorizeRequest
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		// Manager.getAndDelAuthorizationCode verify the redirect_uri of code previously stored
		srv.HandleTokenRequest(w, r)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		jwtToken, ok := srv.BearerAuth(r)
		if !ok {
			http.Error(w, "no token", http.StatusBadRequest)
			return
		}

		// Parse and verify jwt access token
		token, err := jwt.ParseWithClaims(jwtToken, &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
			// if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("parse error")
			}
			// return []byte("SignedKey"), nil
			return ecdsaPublicKey, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(*generates.JWTAccessClaims)
		if !ok || !token.Valid {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"expires_in": int64(time.Until(time.Unix(claims.ExpiresAt, 0)).Seconds()),
			"user_id":    claims.Subject,
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), nil))
}
