#!/bin/bash
# import_business_names_csv.sh
# Imports a specified CSV file into a new SQLite database as business_names_csv table.

set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "Usage: $0 <csv_file> <sqlite_db>"
  exit 1
fi

csv="$1"
sqlite="$2"

if [[ ! -f "$csv" ]]; then
  echo "[error] CSV file not found: $csv"
  exit 1
fi

if [[ -f "$sqlite" ]]; then
  echo "[sqlite] Database $sqlite already exists. Aborting."
  exit 1
fi

mkdir -p "$(dirname "$sqlite")"
echo "[sqlite] Importing $csv to $sqlite as business_names_csv table"
sqlite3 "$sqlite" \
  ".mode tabs" \
  ".import '$csv' business_names_csv"
echo "[sqlite] Import complete."
