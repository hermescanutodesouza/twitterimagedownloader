package main

//twitter scrap and download images from a user timeline.
// with threads

import (
	"log"
	"time"
	"twitterscan/twitter"
	"twitterscan/util"
)

func main() {
	start := time.Now()
	payload := util.CheckUp()

	log.Println("Profiles", payload)

	api := twitter.New()

	var results = make([]chan twitter.Result, len(payload))
	for i, v := range payload {
		results[i] = make(chan twitter.Result)
		go twitter.GetTweeter(api, v.Screenname, results[i]) // run goroutine
	}

	for i := range results {
		r := <-results[i]
		log.Printf("%v - %v", r.Screename, r.Total)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Finished ", elapsed)
}
