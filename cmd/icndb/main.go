// Demo client for the ICNDB (icndb.com).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

const APIURLTmpl = "http://api.icndb.com/jokes/random/%d"

// APIResp represents an API response from the ICNDB.
type APIResp struct {
	Type  string
	Value []struct {
		Joke string
	}
}

func main() {
	nJokes := flag.Int("n", 1, "number of random jokes to retreive")
	flag.Parse()

	url := fmt.Sprintf(APIURLTmpl, *nJokes)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var ar APIResp
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ar)
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range ar.Value {
		fmt.Println(html.UnescapeString(result.Joke))
	}
}
