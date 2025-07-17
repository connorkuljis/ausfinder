build:
	go build --tags "fts5" ./cmd/server

clean:
	rm -f ./main

rsync:
	rsync -av --progress --exclude 'data/' . prod@${remote-ip}:~/app

load-business-names:
	./scripts/import_business_names_csv.sh
