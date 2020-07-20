.PHONY: docker
docker:
	docker build . -t store-svc:latest
