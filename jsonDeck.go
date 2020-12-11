package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonDeck struct {
	Rivers [][]string `json:"rivers"`
	Cards [][]string `json:"cards"`
}

func readJSONDeck(filename string) deck {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jd jsonDeck
	json.Unmarshal(byteValue, &jd)

	return newDeck(jd.Cards, jd.Rivers)
}
