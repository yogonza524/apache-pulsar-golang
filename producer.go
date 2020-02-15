package main

import (
	"os"
	"yogonza524/pulsar-client/src/model"

	twitterscraper "github.com/n0madic/twitter-scraper"
	log "github.com/sirupsen/logrus"
)

func main() {
	p := model.Pulsar{}

	p.Connect()
	if p.Status == 0 {
		panic("Error: Connection not opened")
	}
	for twitt := range twitterscraper.GetTweets(os.Getenv("TWITTER_ACCOUNT"), 25) {
		if twitt.Error != nil {
			panic(twitt.Error)
		}
		p.Produce(twitt.Text)
		log.WithField("tweet", twitt.Text).Info("Tweeter Scraper")
	}
}