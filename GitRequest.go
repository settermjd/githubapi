package githubapi

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"net/http"
)

type GitRequest struct {
	ServerUri   string
	Credentials RequestCredentials
}

func (g *GitRequest) MakeRequest(prId int) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(g.ServerUri, prId), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	req.SetBasicAuth(g.Credentials.Username, g.Credentials.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)

	return bodyText
}
