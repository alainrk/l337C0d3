GO=go

test:
	$(GO) test -v ./...

new:
	./newChallenge.sh