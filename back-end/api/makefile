.PHONY: clean build

build:
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap ./handler/main.go 
	chmod +x bin/bootstrap
	zip -j bin/api.zip bin/bootstrap

clean:
	rm -rf ./bin