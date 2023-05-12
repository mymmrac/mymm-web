local-build:
	go build -o bin/mymm .

build:
	GOOS=linux GOARCH=arm64 go build -o bin/mymm .

build-image:
	docker image build --tag mymmrac/mymm-web:latest-server .

push-image:
	docker image push mymmrac/mymm-web:latest-server

PHONY: local-build build
