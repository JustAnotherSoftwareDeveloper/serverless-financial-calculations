.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/auto-loan-calculator invoke/auto-loan-calcualtor/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/compound-interest-calcualtor invoke/compound-interest-calculator/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/mortgage-calcualtor invoke/mortgage-calculator/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/retirement-calcualtor invoke/retirement-calculator/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
