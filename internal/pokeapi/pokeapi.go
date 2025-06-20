package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var offline_cache = NewCache(60 * time.Second)
var pokedex map[string]pokemon_response = make(map[string]pokemon_response)

func Get_map_data(url string) (map_response, error) {
	//dummy
	var body []byte

	//check offline cache
	body_offline, ok := offline_cache.Get(url)
	//if we have this in cache
	if ok {
		//fmt.Printf("Got data from cache: %v", url)
		//use the cahce
		body = body_offline
	} else {
		//get repsonse
		body_online, err := Get_request(url)
		//if error
		if err != nil {
			//print
			fmt.Printf("get_request err: '%v'\n", err)
			//stop
			return map_response{}, err
		}
		//add to cache
		offline_cache.Add(url, body_online)
		//set as data
		body = body_online
	}

	//empty struct to fill
	data := map_response{}
	//unmarshal the data
	err := Unmarshal(body, &data)
	//if there is an error
	if err != nil {
		//print
		fmt.Printf("unmarshal err: '%v'\n", err)
		//stop
		return data, err
	}
	//return the data
	return data, nil
}

func Get_request(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return body, nil
}

func Unmarshal(data []byte, v interface{}) error {
	//unmarshal
	err := json.Unmarshal(data, &v)
	//if error
	if err != nil {
		//print the error
		fmt.Println(err)
		//return the error
		return err
	}
	//return no error
	return nil
}

func Get_location_data(location string) (location_response, error) {
	//create the url
	url := "https://pokeapi.co/api/v2/location-area/" + location

	//dummy
	var body []byte

	//check offline cache
	body_offline, ok := offline_cache.Get(url)
	//if we have this in cache
	if ok {
		//fmt.Printf("Got data from cache: %v", url)
		//use the cahce
		body = body_offline
	} else {
		//get repsonse
		body_online, err := Get_request(url)
		//if error
		if err != nil {
			//print
			fmt.Printf("get_request err: '%v'\n", err)
			//stop
			return location_response{}, err
		}
		//add to cache
		offline_cache.Add(url, body_online)
		//set as data
		body = body_online
	}

	//empty struct to fill
	data := location_response{}
	//unmarshal the data
	err := Unmarshal(body, &data)
	//if there is an error
	if err != nil {
		//print
		fmt.Printf("unmarshal err: '%v'\n", err)
		//stop
		return data, err
	}

	//return the data
	return data, nil

}

func Catch_pokemon(pokemon string) error {
	//config
	max_catch_chance := 200
	//get the data
	data, err := Get_pokemon_data(pokemon)
	//if error
	if err != nil {
		//return the error
		return err
	}
	//print starting throw
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	//Give the user a chance to catch the Pokemon using the math/rand package.
	chance := rand.Intn(max_catch_chance)
	//debug
	//fmt.Printf("Chance: %v, Base experience: %v\n", chance, data.BaseExperience)

	//You can use the pokemon's "base experience" to determine the chance of catching it.
	// The higher the base experience, the harder it should be to catch.
	if chance > data.BaseExperience {
		//pokemon is caught
		//print message
		fmt.Printf("%v was caught!\n", pokemon)

		//check if already in pokedex
		_, ok := pokedex[pokemon]
		if ok {
			//print message
			fmt.Printf("Data of %v was already added to the pokedex!\n", pokemon)
		} else {
			//print message
			fmt.Printf("Data of %v has been added to the pokedex!\n", pokemon)
			//add to pokedex
			pokedex[pokemon] = data
		}

	} else {
		//print fail message
		fmt.Printf("%v escaped!\n", pokemon)
	}
	return nil

}

func Get_pokemon_data(name string) (pokemon_response, error) {
	//create url
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	//dummy
	var body []byte

	//check offline cache
	body_offline, ok := offline_cache.Get(url)
	//if we have this in cache
	if ok {
		//fmt.Printf("Got data from cache: %v", url)
		//use the cahce
		body = body_offline
	} else {
		//get repsonse
		body_online, err := Get_request(url)
		//if error
		if err != nil {
			//print
			fmt.Printf("get_request err: '%v'\n", err)
			//stop
			return pokemon_response{}, err
		}
		//add to cache
		offline_cache.Add(url, body_online)
		//set as data
		body = body_online
	}

	//empty struct to fill
	data := pokemon_response{}
	//unmarshal the data
	err := Unmarshal(body, &data)
	//if there is an error
	if err != nil {
		//print
		fmt.Printf("unmarshal err: '%v'\n", err)
		//stop
		return data, err
	}

	//return the data
	return data, nil
}
