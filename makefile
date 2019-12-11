run-redis:
	docker-compose up -d

shutdown-redis:
	docker-compose down

publisher:
	go run cmd/publisher/main.go

medianfinder:
	go run cmd/subscriber/medianfinder/main.go

sumfinder:
	go run cmd/subscriber/sumfinder/main.go
