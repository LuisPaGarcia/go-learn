package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BaseUrl = "https://swapi.dev/api/"

type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func main() {

	resp, err := http.Get(BaseUrl + "people")
	if err != nil {
		log.Print("Failed get request to swapi")
	}

	bytes, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		log.Print("Failed reading resp.Body")
	}
	var people AllPeople

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsin json", err)
	}

	for _, peop := range people.People {
		peop.getPlanet()
		fmt.Println(peop)
	}

}

func (p *Person) getPlanet() {
	resp, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error on getting homeworld")
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Print("Error on getting homeworld")
	}

	json.Unmarshal(bytes, &p.Homeworld)
}
