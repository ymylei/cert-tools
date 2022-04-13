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

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

test:
	go test -v ./...
