test: example/args.json.go
	go test ./...
	go run ./example

diff:
	cp ./example/args.json.go example/args.json.go.bak
	$(MAKE) example/args.json.go
	diff example/args.json.go example/args.json.go.bak

example/args.json.go: example/*
	go generate ./example/main.go
