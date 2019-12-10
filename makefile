go-mod:
	go mod download
	go mod vendor

docker-publisher:
	go-mod
	docker build -t publisher cmd/publisher/

docker-subscriber:
	go-mod
	docker build -t consumer cmd/subscriber/sumfinder/