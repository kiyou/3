all: *.go 
	make -C mx
	go tool vet *.go
	go install
