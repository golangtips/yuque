build:
	docker buildx build --platform linux/amd64 -f Dockerfile -t yuque:latest .