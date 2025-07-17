#!/bin/bash
# download_latest_business_names.sh
# Downloads the latest business names dataset from data.gov.au into ./data/

set -euo pipefail

mkdir -p ./data

date=$(date "+%Y%m")
dataset_id="business_names_$date"
csv="./data/${dataset_id}.csv"

if [[ -f "$csv" ]]; then
  echo "[csv] Using existing dataset: $csv"
  exit 0
fi

url="https://data.gov.au/data/dataset/bc515135-4bb6-4d50-957a-3713709a76d3/resource/55ad4b1c-5eeb-44ea-8b29-d410da431be3/download/${dataset_id}.csv"
echo "[info] Downloading $url"
curl -fLo "$csv" "$url"
echo "[info] Downloaded to $csv"
