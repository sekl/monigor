package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

type Config struct {
	WebhookURL string `json:"webhook_url"`
	Channel    string `json:"channel"`
	BotName    string `json:"bot_name"`
}

type Site struct {
	URL     string `json:"url"`
	Element string `json:"element"`
	Text    string `json:"text"`
}

type Payload struct {
	Username string `json:"username"`
	Channel  string `json:"channel"`
	Text     string `json:"text"`
}

var waitGroup sync.WaitGroup
var config Config

func main() {
	raw := read("./config.json")
	err := json.Unmarshal(raw, &config)
	if err != nil {
		log.Print(err)
	}

	raw = read("./urls.json")
	var sites []Site
	err = json.Unmarshal(raw, &sites)
	if err != nil {
		log.Print(err)
	}

	waitGroup.Add(len(sites))

	for _, site := range sites {
		go scanSite(site)
	}

	waitGroup.Wait()

}

func scanSite(site Site) {
	doc, err := goquery.NewDocument(site.URL)
	if err != nil {
		log.Print(err)
		waitGroup.Done()
		return
	}

	doc.Find(site.Element).Each(func(i int, s *goquery.Selection) {
		result := s.Text()

		if len(result) == 0 {
			fmt.Printf("No result for %s at %s", site.Element, site.URL)
			return
		}
		fmt.Printf("Found: %s\n", result)

		webhookURL := config.WebhookURL
		text := fmt.Sprintf("Found: %s - <%s>", result, site.URL)
		payload := Payload{Username: config.BotName, Channel: config.Channel, Text: text}
		request := gorequest.New()
		resp, _, errs := request.Post(webhookURL).
			Send(payload).
			End()
		fmt.Println(resp.Status)
		if errs != nil {
			log.Print(errs)
			waitGroup.Done()
			return
		}
	})

	defer waitGroup.Done()
}

func read(path string) []byte {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	return raw
}
