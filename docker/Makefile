TAG=latest
IMAGENAME=mumax3-dev
DOCKERPORT=35367
HOSTPORT=35367

HOST_PROJECT_DIR=$(shell readlink -f $(dir $(lastword $(MAKEFILE_LIST)))..)
MUMAX3_BIN=$(HOST_PROJECT_DIR)/bin/mumax3
CONTAINER_PROJECT_DIR=/go/src/github.com/mumax/3
HOST_JOB_DIR=/tmp/mumax3_jobs
CONTAINER_JOB_DIR=/tmp
CONTAINER_BUILD_NAME=mumax3-build
CONTAINER_RUN_NAME=mumax3
CONTAINER_SERVER_NAME=mumax3-server

.PHONY: build run server distclean

build: $(MUMAX3_BIN)

$(MUMAX3_BIN): Dockerfile
	docker build -t $(IMAGENAME):$(TAG) .
	docker run --runtime=nvidia -d --rm \
	--name $(CONTAINER_BUILD_NAME) \
	-v $(HOST_PROJECT_DIR):$(CONTAINER_PROJECT_DIR) \
	$(IMAGENAME):$(TAG) \
	/bin/bash -c \
	'cd $(CONTAINER_PROJECT_DIR) && ./make.bash && mkdir -p $(CONTAINER_PROJECT_DIR)/bin && cp /go/bin/* $(CONTAINER_PROJECT_DIR)/bin'

run: $(MUMAX3_BIN)
	docker run --runtime=nvidia -d --rm \
	--name $(CONTAINER_RUN_NAME) \
	-v $(HOST_PROJECT_DIR):$(CONTAINER_PROJECT_DIR) \
	-e PORT=$(DOCKERPORT) \
	-p $(HOSTPORT):$(DOCKERPORT) \
	$(IMAGENAME):$(TAG) \
	$(CONTAINER_PROJECT_DIR)/bin/mumax3 -http=:$(DOCKERPORT)

server: $(MUMAX3_BIN)
	docker run --runtime=nvidia -d --rm \
	--name $(CONTAINER_SERVER_NAME) \
	-v $(HOST_PROJECT_DIR):$(CONTAINER_PROJECT_DIR) \
	-v $(HOST_JOB_DIR):$(CONTAINER_JOB_DIR) \
	-e PORT=$(DOCKERPORT) \
	-p $(HOSTPORT):$(DOCKERPORT) \
	-w $(CONTAINER_JOB_DIR) \
	$(IMAGENAME):$(TAG) \
	$(CONTAINER_PROJECT_DIR)/bin/mumax3-server -exec=$(CONTAINER_PROJECT_DIR)/bin/mumax3 -l=":$(DOCKERPORT)" -ports="$(DOCKERPORT)" -scan="127.0.0.1"

distclean:
	rm $(HOST_PROJECT_DIR)/bin/* -rf
