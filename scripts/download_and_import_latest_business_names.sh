#!/bin/bash
# get_dataset.sh
# gets the latest business names dataset from data.gov.au

date=$(date "+%Y%m")
dataset_id="business_names_$date"
csv="./data/${dataset_id}.csv"
sqlite="./database/${dataset_id}.sqlite3"

echo "[info] Dataset ID: $dataset_id"
echo "[info] Importing $csv to $sqlite"

if [[ ! -f "$csv" ]]; then 
	url="https://data.gov.au/data/dataset/bc515135-4bb6-4d50-957a-3713709a76d3/resource/55ad4b1c-5eeb-44ea-8b29-d410da431be3/download/${dataset_id}.csv"
	echo "$url"
	curl -o "$csv" "$url"
else 
	echo "[csv] using existing dataset: $csv"
fi

if [[ -f "$sqlite" ]]; then 
	echo "[sqlite] database "$sqlite" already exists. aborting"
	exit 1
fi

echo "[sqlite] importing csv data"
sqlite3 "$sqlite" \
	".mode tabs" \
	".import "$csv" business_names_csv"
