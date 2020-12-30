package server

import(
	"lib"
	"os"
	"fmt"
	"encoding/json"
	"objects"
	"io/ioutil"
)

func LoadConfig(){

	defer lib.Handlepanic()

	// Open our jsonFile
	jsonFile, err := os.Open("config.json")

	// if we os.Open returns an error then handle it
	if err != nil {
	    fmt.Println(err)
	    return
	}

	fmt.Println(string(objects.ColorGreen), "Config json file loaded to inmemory...")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
	    fmt.Println(err)
	    return
	}

	objects.Conf = &objects.Config{}

	json.Unmarshal(byteValue, objects.Conf)

	fmt.Println(string(objects.ColorGreen), "Config parsed successfully and loaded...")
}