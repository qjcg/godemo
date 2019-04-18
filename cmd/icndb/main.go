// Command-line client for the ICNDB JSON API (http://www.icndb.com/api).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

// Ref: http://www.icndb.com/api/
const URLTemplate = "https://api.icndb.com/jokes/random/%d"

// APIResp represents an API response from the ICNDB.
type APIResp struct {
	Value []struct {
		Joke string
	}
}

func main() {

	// Accept a CLI flag specifying the number of jokes to print.
	nJokes := flag.Int("n", 1, "number of random jokes to retreive")
	flag.Parse()

	// Make an API request.
	url := fmt.Sprintf(URLTemplate, *nJokes)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var ar APIResp
	decode(resp.Body, &ar)
	print(os.Stdout, &ar)
}

// Decode the JSON API response into an APIResp struct.
func decode(r io.Reader, v *APIResp) error {
	return json.NewDecoder(r).Decode(v)
}

// Print the jokes from our API response.
// We also unescape any HTML character entity references.
// Ref: https://en.wikipedia.org/wiki/List_of_XML_and_HTML_character_entity_references#Character_entity_references_in_HTML
func print(w io.Writer, v *APIResp) {
	for _, item := range v.Value {
		item.Joke = html.UnescapeString(item.Joke)
		fmt.Fprintln(w, item.Joke)
	}
}
