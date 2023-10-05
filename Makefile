TELEGRAM_API_KEY ?= ""

run: build
	docker run --name reviewbot -d -p 127.0.0.1:50051:50051 -e "TELEGRAM_API_KEY=$(TELEGRAM_API_KEY)" reviewbot:latest
	docker run --name reviewbot-mgr -d -p 127.0.0.1:5000:5000 reviewbot-mgr:latest

build: build-reviewbot_mgr build-reviewbot

build-reviewbot_mgr:
	pip install -r reviewbot_mgr/requirements.txt
	cd reviewbot_mgr && docker build -t reviewbot-mgr:latest .

build-reviewbot: test
	mkdir -p reviewbot/build
	cd reviewbot && go mod tidy
	cd reviewbot && GOOS=linux GOARCH=amd64 go build -o build/reviewbot ./cmd/app/main.go && chmod +x build/reviewbot
	cd reviewbot && docker build -t reviewbot:latest .

generate: gen-go gen-python

gen-go:
	protoc --go_out=. --go-grpc_out=. proto/reviewbot.proto

gen-python:
	python3 -m grpc_tools.protoc -I=proto --python_out=reviewbot_mgr/api/ --pyi_out=reviewbot_mgr/api/ --grpc_python_out=reviewbot_mgr/api/ reviewbot.proto

test:
	cd reviewbot && go test -cover ./...

clean:
	rm -rf reviewbot/build
	cd reviewbot && go clean -testcache -cache -modcache
	cd reviewbot && go mod tidy

stop:
	docker kill reviewbot reviewbot-mgr
	docker rm reviewbot reviewbot-mgr
