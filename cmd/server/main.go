package main

import (
	"fmt"
	"log"

	"github.com/connorkuljis/backtrace/internal/abr"
)

const (
	business = "{{ Business Name Here }}"
)

func main() {
	fmt.Println(business, "- search for Australian companies\n")

	client := abr.ABRXMLSearchClient{}

	abn := "14159799550"
	res, err := client.SearchByABN(abn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
