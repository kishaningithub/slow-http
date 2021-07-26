fmt: ## Run the code formatter
	gofmt -l -s -w .

build: fmt
	CGO_ENABLED=0 go build -o bin/slow-http main.go

docker-build:
	docker build . -t kishanb/slow-http:1.0.1

docker-push:
	docker push kishanb/slow-http:1.0.1

run: docker-build
	docker run -p 8080:8080 kishanb/slow-http:1.0.1