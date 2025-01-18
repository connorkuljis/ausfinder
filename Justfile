build:
    go build -v --tags "fts5" ./cmd/server/main.go

clean:
    rm -f ./main

rsync remote-ip: clean
    rsync -av --progress --exclude 'data/' . prod@{{ remote-ip }}:~/app

load-business-names:
    ./scripts/import_business_names_csv.sh
    
