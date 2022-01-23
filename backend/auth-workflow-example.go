package main

// This module can be used by admin for triggering various APIs

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	dao "./dao"
	jwt "github.com/golang-jwt/jwt/v4"
)

const privKeyPath = "auth/admin.rsa"
const baseURL = "http://localhost:8080"

var signKey *rsa.PrivateKey

func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %s", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("failed to parse private key: %s", err)
	}
}

func createToken(key *rsa.PrivateKey) (string, error) {
	// create a signer for rsa 512
	t := jwt.New(jwt.GetSigningMethod("RS512"))

	// set our claims
	t.Claims =
		&jwt.StandardClaims{
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "tamilaadal-admin",
			Audience:  "tamilaadal",
		}

	// Creat token string
	return t.SignedString(key)
}

func main() {
	// 1. Create an user
	userId := createUser(dao.User{Name: "sankar1", TwitterHandle: "psankar"})
	// 2. Mark the user active
	markUserActive(userId)
	// 3. Use admin to generate JWT token for the user
	token := getAuthToken(userId)
	// 4. Share the token and the user id with the user
	// 5. User has to use the token to generate & download private key
	privateKey := downloadPrivateKey(userId, token)
	// 6. For adding words, use the generated private key
	addWord("தமிழாடல்", "2022-01-19", userId, privateKey)
	addWord("தமிழாடல்", "2022-01-21", userId, privateKey)
	addWord("தமிழாடல்", "2022-01-22", userId, privateKey)
	addWord("தமிழாடல்", "2022-01-23", userId, privateKey)
}

func createUser(user dao.User) string {
	token, err := createToken(signKey)
	if err != nil {
		log.Fatalf("failed to create token: %s", err)
	}

	// send Authorization header
	// send user in body
	userJson, err := json.Marshal(user)
	req, err := http.NewRequest("POST", baseURL+"/admin/create-user", strings.NewReader(string(userJson)))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response: %s", err)
	}
	log.Printf("User: %s", body)
	return string(body)
}

func markUserActive(userId string) {
	token, err := createToken(signKey)
	if err != nil {
		log.Fatalf("failed to create token: %s", err)
	}

	// send Authorization header
	// send userid in body
	req, err := http.NewRequest("POST", baseURL+"/admin/mark-user-active", strings.NewReader(`"`+userId+`"`))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response: %s", err)
	}
	log.Printf("Response: %s", body)
}

func getAuthToken(userId string) string {
	token, err := createToken(signKey)
	if err != nil {
		log.Fatalf("failed to create token: %s", err)
	}

	// send Authorization header
	// send userid in body
	req, err := http.NewRequest("POST", baseURL+"/admin/generate-auth-token", strings.NewReader(`"`+userId+`"`))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response: %s", err)
	}
	log.Printf("Token: %s", body)
	return string(body)
}

func downloadPrivateKey(userId string, token string) string {
	// send Authorization header
	// send userid in body
	req, err := http.NewRequest("POST", baseURL+"/user/download-private-key", strings.NewReader(`"`+userId+`"`))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response: %s", err)
	}
	log.Printf("Private Key: %s", body)
	return string(body)
}

func addWord(word string, date string, userId string, privateKey string) {
	// get the private key
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		log.Fatalf("failed to parse private key: %s", err)
	}

	token, err := createToken(key)
	if err != nil {
		log.Fatalf("failed to create token: %s", err)
	}

	// send Authorization header
	// send userid in body
	req, err := http.NewRequest("POST", baseURL+"/user/add-word", strings.NewReader(`{"word":"`+word+`","date":"`+date+`", "userId":"`+userId+`"}`))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("x-user-id", userId)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response: %s", err)
	}
	log.Printf(string(body))
}
