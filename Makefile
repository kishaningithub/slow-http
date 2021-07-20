build:
	go build -o bin/slow-http main.go


docker-build:
	docker build . -t kishanb/slow-http:1.0.0


docker-push:
	docker push kishanb/slow-http:1.0.0

run: docker-build
	docker run kishanb/slow-http:1.0.0