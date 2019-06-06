TAG=latest
IMAGENAME=mumax3-dev
DOCKERPORT=35367
HOSTPORT=35367

HOST_PROJECT_DIR=$(shell readlink -f $(dir $(lastword $(MAKEFILE_LIST)))..)
MUMAX3_BIN=$(HOST_PROJECT_DIR)/bin/mumax3
CONTAINER_PROJECT_DIR=/go/src/github.com/mumax/3

.PHONY: build run distclean

build: $(MUMAX3_BIN)

$(MUMAX3_BIN): Dockerfile
	docker build -t $(IMAGENAME):$(TAG) .
	docker run --runtime=nvidia -it --rm \
	-v $(HOST_PROJECT_DIR):$(CONTAINER_PROJECT_DIR) \
	$(IMAGENAME):$(TAG) \
	/bin/bash -c \
	'cd $(CONTAINER_PROJECT_DIR) && ./make.bash && cp /go/bin/* $(CONTAINER_PROJECT_DIR)/bin'

run: $(MUMAX3_BIN)
	docker run --runtime=nvidia -it --rm \
	-v $(HOST_PROJECT_DIR):$(CONTAINER_PROJECT_DIR) \
	-e PORT=$(DOCKERPORT) \
	-p $(HOSTPORT):$(DOCKERPORT) \
	$(IMAGENAME):$(TAG) \
	$(CONTAINER_PROJECT_DIR)/bin/mumax3 -http=:$(DOCKERPORT)

distclean:
	rm $(HOST_PROJECT_DIR)/bin/* -rf
