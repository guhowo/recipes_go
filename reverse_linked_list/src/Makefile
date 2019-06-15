export GOPATH=$(shell dirname `pwd`):$(GOINGPATH)

all:
	go build -o ../bin/go_delete_cos_svr ./main.go ./handle.go
	go build -o ../bin/trydelete ./trydelete.go

clean:
	-rm -f ../bin/*
