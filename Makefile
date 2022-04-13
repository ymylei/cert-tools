build:
	go build ./cmd/cert-tools

check:
	goimports -l -d .
	go vet ./...
	staticcheck ./...

clean:
	go clean -cache -testcache ./...

clean-all:
	go clean -cache -testcache -modcache ./...

format:
	goimports -l -w .

test:
	go test -v ./...
