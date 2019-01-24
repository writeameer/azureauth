package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/writeameer/azureauth/models"
)

func main() {

	postURL := "https://login.microsoftonline.com/cloudoman.com/oauth2/token"

	resp, err := http.PostForm(postURL, url.Values{
		"resource":   {"https://graph.windows.net"},
		"client_id":  {os.Getenv("AZURE_CLIENT_ID")}, // This is the application ID of a "native application" you've create on Azure AD
		"grant_type": {"password"},
		"username":   {os.Getenv("AZURE_USERNAME")},
		"password":   {os.Getenv("AZURE_PASSWORD")},
		"scope":      {"openid"},
	})

	// Get response details
	responseBytes, err := ioutil.ReadAll(resp.Body)

	// Quit if Login failed
	if resp.StatusCode != 200 {
		log.Printf("Login failed: %s ", resp.StatusCode)
		fmt.Println(string(responseBytes))
		os.Exit(1)
	} else {
		log.Println("======== Login Success =============")
	}

	// Output Azure token
	token, err := models.UnmarshalToken(responseBytes)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(token.ExpiresOn)
	log.Println(token.AccessToken)

}
