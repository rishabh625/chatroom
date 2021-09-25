APP		    := chatroom
GO_VARS ?=
GO ?= go

help:
	@echo "  clean        clean binary file created"	
	@echo "  build        to build the main binary for current platform"
	@echo "  test         to run unittests"
	@echo "  build-docker to create docker image"
	@echo "  run-chatroom   to run application"
	@echo "  stop-chatroom  to stop application with postgresql"


build-docker: 
	@echo "Building Docker image"
	docker build -t $(APP) .

build:
	go build -o=$(APP) $(GOPATH)/src/$(APP)/cmd/main/

test:
	$(GO_VARS) $(GO) test -v $(GOPATH)/src/$(APP)/tests

clean:
	rm -f $(APP)

run-chatroom:
	docker-compose up --build  -d 

stop-chatroom:
	docker-compose down