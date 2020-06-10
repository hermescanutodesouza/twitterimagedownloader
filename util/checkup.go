package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func CheckUp() []Screename {

	var payload []Screename

	CheckFolder(ResolveFilePath("/files"))
	CheckFolder(ResolveFilePath("/img"))

	jsonfile := ResolveFilePath("/files/twitter.json")
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		var ok []Screename
		ok = append(ok, Screename{Screenname: "hermes"})
		e, err := json.MarshalIndent(ok, "", "\t")
		if err != nil {
			log.Println(err)
		}
		contend := string(e)
		log.Println(contend)
		CreateJsonFile(jsonfile, contend)
		log.Println("Create the File")
		payload = ok
	} else {
		err = json.Unmarshal(data, &payload)
		CatchGeneralError(&err)
	}
	return payload
}
