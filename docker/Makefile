TAG=latest
IMAGENAME=mumax3-dev
DOCKERPORT=35367
HOSTPORT=35367

.PHONY: docker_build mumax3_build mumax3_run

docker_build: Dockerfile
	docker build -t $(IMAGENAME):$(TAG) .

mumax3_build: docker_build
	docker run --runtime=nvidia -it --rm \
	-v $(GOPATH):/go \
	$(IMAGENAME):$(TAG) \
	/bin/bash -c \
	'cd src/github.com/mumax/3 && make'

mumax3_run: mumax3_build
	docker run --runtime=nvidia -it --rm \
	-v $(GOPATH):/go \
	-e PORT=$(DOCKERPORT) \
	-p $(HOSTPORT):$(DOCKERPORT) \
	$(IMAGENAME):$(TAG) \
	mumax3 -http=:$(DOCKERPORT)
