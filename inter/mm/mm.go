package mm

//match making

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/sirupsen/logrus"
)

var Games = make(map[string]*network.Client)

// func FindGIDs() string {
// 	var gameID string	
// 	for id := range Games {
// 		gameID = id
// 		jsonStr := []string{id}
// 		logrus.WithFields(logrus.Fields{
// 			" ": jsonStr,
// 		}).Info("===Player Joined Game=== " + id) // TODO +uID joined Server
// 	}
// 	return gameID

func FindGIDs(heroID string, ip string) []string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	logrus.Println("====MATCHMAKING REQUEST==============")

	client := &http.Client{Transport: tr}
	url := "http://127.0.0.1/api/mm/findgame/" + heroID + "/" + ip
	logrus.Println(url)
	logrus.Println("TESTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT")

	resp, err := client.Get(url)
	if err != nil {
		logrus.Println("Error making request to matchmaking api")
		return make([]string, 0)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Println("Error reading from response to matchmaking api")
		return make([]string, 0)
	}

	return strings.Split(string(body[:]), ",")

}
