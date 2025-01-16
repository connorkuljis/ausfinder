#!/bin/bash

set -x 

# Set the database file and CSV file paths
DB_FILE="./db/db.sqlite3"
CSV_FILE="./data/BUSINESS_NAMES_202501.csv"

# Set the table name
TABLE_NAME="business_names_csv"


# Import the CSV file into the SQLite database
sqlite3 "$DB_FILE" ".mode tabs" ".import $CSV_FILE $TABLE_NAME"
