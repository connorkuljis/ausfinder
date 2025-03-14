#!/bin/bash
# get_dataset.sh
# gets the latest business names dataset from data.gov.au

set -x

current_date=$(date +"%Y%m")

url="https://data.gov.au/data/dataset/bc515135-4bb6-4d50-957a-3713709a76d3/resource/55ad4b1c-5eeb-44ea-8b29-d410da431be3/download/business_names_${current_date}.csv"

curl -o "./data/business_names_${current_date}.csv" "$url"
