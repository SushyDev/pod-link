run:
	go run main.go

docker-build:
	docker build --network host -t pod-link:latest .
docker-run:
	docker run --rm --name pod-link --network host pod-link:latest
