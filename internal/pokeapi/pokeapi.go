package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//map data
type map_data struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

//description of the JSON map data
type map_response struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []map_data `json:"results"`
}

func Get_map_data(url string) (map_response, error) {
	//empty struct to fill
	data := map_response{}

	//get repsonse
	body, err := Get_request(url)

	//if error
	if err != nil {
		//print
		fmt.Printf("get_request err: '%v'\n", err)
		//stop
		return data, err
	}
	
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

