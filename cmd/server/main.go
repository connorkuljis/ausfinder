package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/connorkuljis/backtrace/internal/abr"
	"github.com/connorkuljis/backtrace/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	dbstr     = "file:data/db.sqlite3"
	randomABN = "88573118334"
)

func main() {
	err := sqliteDemo()
	if err != nil {
		log.Fatal(fmt.Errorf("Error: %w", err))
	}

	fmt.Println("ok")
}

func sqliteDemo() error {
	db, err := connect()
	if err != nil {
		return err
	}

	var business model.Business
	err = db.Get(&business, `SELECT * FROM business_names WHERE trim(BN_ABN) = ?`, randomABN)
	if err != nil {
		return err
	}

	log.Println(business)

	return nil
}

func connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dbstr)
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to ", dbstr)

	return db, nil
}

func abrSearchExample() {
	client := abr.ABRXMLSearchClient{}

	res, err := client.SearchByABN(randomABN)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
