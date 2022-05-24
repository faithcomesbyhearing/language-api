.PHONY: build clean deploy undeploy offline gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/language language/main.go

clean:
	rm -rf ./bin ./vendor go.sum

offline: build
	sls offline

deploy: clean build
	sls deploy --verbose

undeploy:
	sls remove

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
	go mod tidy
