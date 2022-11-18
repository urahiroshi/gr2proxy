build-proxy:
	go build -o bin/proxy ./cmd/proxy

start-proxy:
	pkill -f bin/proxy | true
	bin/proxy &

build-test-proto:
	protoc --go_out=plugins=grpc,paths=source_relative:. test/helloworld/helloworld.proto

build-test-server:
	go build -o test/bin/server ./test/cmd/server

build-test-client:
	go build -o test/bin/client ./test/cmd/client

build-test-record:
	go build -o test/bin/record ./test/cmd/record

start-test-server:
	pkill -f test/bin/server | true
	test/bin/server &

start-test-client:
	test/bin/client

start-test-record:
	pkill -f test/bin/record | true
	test/bin/record &

get-fixtures:
	curl localhost:50151/fixtures
