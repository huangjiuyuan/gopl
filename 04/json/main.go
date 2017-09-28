package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/huangjiuyuan/gopl/04/json/github"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// Marshal produces a byte slice containing a very long string with no
	// extraneous white space.
	// Marshaling uses the Go struct field names as the field names for the
	// JSON objects. Only exported fields are marshaled.
	// A field tag is a string of metadata associated at compile time with the
	// field of a struct.
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	data, err = json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// The inverse operation to marshaling, decoding JSON and populating a Go
	// data structure, is called unmarshaling, it is done by json.Unmarshal.
	// We can select which parts of the JSON input to decode and which to
	// discard. When Unmarshal returns, it has filled in the slice with the
	// Title information; other names in the JSON are ignored.
	var titles []struct {
		Title string
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
