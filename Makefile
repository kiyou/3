# You don't need this makefile. It's only here because I'm used to typing ':make' in vim.
# Build with 'go install' if you have not edited any sources or with
# './make.bash' if you did.
all:
	./make.bash

TAG=latest
IMAGENAME=mumax3-dev
DOCKERPORT=35367
HOSTPORT=35367

.PHONY: docker_build mumax3_build mumax3_run

docker_build: docker/Dockerfile
	docker build -t $(IMAGENAME):$(TAG) docker/

bin/mumax3: docker_build
	docker run --runtime=nvidia -it --rm \
	-v $(PWD):/go/src/github.com/mumax/3 \
	-v $(PWD)/bin:/go/bin \
	$(IMAGENAME):$(TAG) \
	/bin/bash -c \
	'cd src/github.com/mumax/3 && make'

docker_run: bin/mumax3
	docker run --runtime=nvidia -it --rm \
	-v $(PWD):/go/src/github.com/mumax/3 \
	-v $(PWD)/bin:/go/bin \
	-e PORT=$(DOCKERPORT) \
	-p $(HOSTPORT):$(DOCKERPORT) \
	$(IMAGENAME):$(TAG) \
	/go/bin/mumax3 -http=:$(DOCKERPORT)
