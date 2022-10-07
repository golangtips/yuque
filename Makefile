build:
	docker buildx build --platform linux/amd64 -f Dockerfile -t yuque:latest .

run:
	go run cmd/main.go cmd/wire_gen.go