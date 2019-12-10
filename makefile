go-mod:
	go mod download
	go mod vendor

docker-publisher:
	docker build -t publisher cmd/publisher/

docker-subscriber:
	docker build -t consumer cmd/subscriber/sumfinder/

docker-all:
	docker-compose build
	docker-compose up -d
