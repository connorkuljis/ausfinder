build:
    go build -v --tags "fts5" ./cmd/server/main.go

load-business-names:
    ./scripts/import_business_names_csv.sh
    
