.PHONY: build-docker run-docker

build-docker:
	@docker build -f ./build/package/treebabel/Dockerfile . -t treebabel:latest

run-docker:
	@docker run -it treebabel:latest bash