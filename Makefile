build:
	go build --tags "fts5" ./cmd/server

clean:
	rm -f ./main

deploy:
	rsync -av --progress --exclude 'data/' . prod@$(REMOTE_IP):~/app

load-business-names:
	./scripts/import_business_names_csv.sh
