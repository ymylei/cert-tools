build:
	go build -o ./out ./cmd/certtools 

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

install: build
	chmod +x ./out/certtools
	mv ./out/certtools ~/go/bin

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

test:
	go test -v ./...

.phony: build check clean clean-all format tools test