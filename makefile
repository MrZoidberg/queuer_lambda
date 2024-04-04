build-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -tags lambda.norpc -o .bin/bootstrap ./...
	cd .bin && chmod 755 bootstrap && zip function-arm64.zip bootstrap

build-x84:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"  -tags lambda.norpc -o .bin/bootstrap ./...
	cd .bin && zip function-amd64.zip bootstrap

build: build-arm64 build-x84