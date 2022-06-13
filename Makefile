.PHONY: build clean deploy undeploy offline gomodgen

build: gomodgen
	export GO111MODULE=on
	export GODEBUG=netdns=go+1
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/db db/main.go
	env GOARCH=amd64 GOOS=linux  GODEBUG=netdns=go+1 go build  -o bin/language language/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/awslambdarpc awslambdarpc/awslambdarpc.go
	env GOARCH=amd64 GOOS=linux go build -o bin/experimental experimental/main.go

clean:
	rm -rf ./bin ./vendor go.sum go.mod

offline: build
	SLS_DEBUG=* sls offline --printOutput --dockerNetwork=db

deploy: clean build
	sls deploy --verbose

undeploy:
	sls remove

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
	go mod tidy
