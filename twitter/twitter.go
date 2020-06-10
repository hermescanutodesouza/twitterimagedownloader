package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"twitterscan/config"
	"twitterscan/util"
)

type Result struct {
	Screename string
	Total     int
}

func New() *anaconda.TwitterApi {
	config:=config.AppConfig()
	api := anaconda.NewTwitterApiWithCredentials(
		config.AcessToken,
		config.AcessTokenSecret,
		config.ConsumerKey,
		config.ConsumerSecret)
	return api
}

func GetTweeter(api *anaconda.TwitterApi, screenname string, done chan Result) {
	folder := util.ResolveFilePath("/img/" + screenname)
	util.CheckFolder(folder)
	v := url.Values{}
	v.Set("screen_name", screenname)
	v.Set("count", "6000")
	searchResult, _ := api.GetUserTimeline(v)

	a := 0

	for _, tweet := range searchResult {
		media := tweet.Entities.Media
		for _, v := range media {
			file := filepath.Base(v.Media_url)
			if _, err := os.Stat(folder + "/" + file); os.IsNotExist(err) {
				err := util.DownloadFile(folder+"/"+file, v.Media_url)
				if err != nil {
					log.Println(err)
					return
				}
				log.Println(screenname, file)
				a += 1
			}
		}
	}
	var r Result
	r.Screename = screenname
	r.Total = a
	done <- r
}
