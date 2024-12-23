run: 
	@go run .

test:
	@go test ./...

race:
	@go test ./... --race

clean:
	go clean

.PHONY: version
version:
	sed -i.bck "s|SERVICE_VERSION=\"[0-9]*.[0-9]*.[0-9]*.*\"|SERVICE_VERSION=\"${VERSION}\"|" "Dockerfile"
	rm -fr "Dockerfile.bck"
	git add "Dockerfile"
	git commit -m "bump v${VERSION}"
	git tag v${VERSION}