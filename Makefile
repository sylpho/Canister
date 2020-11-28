build:
	go build -o canister *.go

init:
	go get ./...

clean:
	rm -rf canister