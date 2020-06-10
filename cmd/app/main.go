package main

//twitter scrap and download images from a user timeline.
// with threads

import (
	"log"
	"twitterscan/twitter"
	"twitterscan/util"
)



func main() {
	payload := util.CheckUp()

	log.Println("Profiles", payload)

	api := twitter.New()

	var results []chan twitter.Result
	for i, v := range payload {
		results = append(results, make(chan twitter.Result))
		go twitter.GetTweeter(api, v.Screenname, results[i]) // run goroutine
	}

	for i := range results {
		r := <-results[i]
		log.Printf("%v - %v", r.Screename, r.Total)
	}
	log.Println("Finished")
}
