#!/bin/bash

set -x 

# Set the database file and CSV file paths
csv="./data/business_names_202503.csv"

# Set the table name
table="business_names_csv"

# Import the CSV file into the SQLite database
sqlite3 "$GOOSE_DBSTRING" ".mode tabs" ".import $csv $table"
