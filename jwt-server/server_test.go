package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
)

func TestJWT(t *testing.T) {
	// res, err := http.Get(fmt.Sprintf("http://localhost:%v/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read", ServerPort))
	res, err := http.PostForm(fmt.Sprintf("http://localhost:%v/token", ServerPort), url.Values{
		"username":      {"test"},
		"password":      {"known"},
		"grant_type":    {oauth2.PasswordCredentials.String()},
		"client_id":     {"000000"},
		"client_secret": {"999999"},
	})

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
		return
	}

	// Read the token out of the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	// tokenString := strings.TrimSpace(buf.String())

	var data map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		log.Fatal(err.Error())
		return
	}

	// Parse and verify jwt access token
	token, err := jwt.ParseWithClaims(data["access_token"].(string), &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte("SignedKey"), nil
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	claims, ok := token.Claims.(*generates.JWTAccessClaims)
	if !ok || !token.Valid {
		log.Fatal("invalid token")
		return
	}

	rs := map[string]interface{}{
		"expires_in": int64(time.Until(time.Unix(claims.ExpiresAt, 0)).Seconds()),
		"user_id":    claims.Subject,
	}
	fmt.Println(rs)
}
