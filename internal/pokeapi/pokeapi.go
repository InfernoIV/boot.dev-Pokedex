package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// map data
type map_data struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// description of the JSON map data
type map_response struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []map_data `json:"results"`
}

// description of the JSON map data
type location_response struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var offline_cache = NewCache(60 * time.Second)

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

func Get_location_data(location string) (response location_response, err error) {
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
	err = Unmarshal(body, &data)
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