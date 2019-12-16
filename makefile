run-redis:
	docker-compose up -d

shutdown-redis:
	docker-compose down

publisher:
	go run cmd/publisher/main.go ./config.json

medianfinder:
	go run cmd/subscriber/medianfinder/main.go ./config.json

sumfinder:
	go run cmd/subscriber/sumfinder/main.go ./config.json

test:
	go test ./...
