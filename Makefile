build:
	@echo "go build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .build/markman-server -i .

build-image:
	@docker build --file=Dockerfile . -t xem100744/markman:${tag}

push-image:
	@docker push xem100744/markman:${tag}