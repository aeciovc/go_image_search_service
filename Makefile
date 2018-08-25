ENTRY=cmd/imagesearch/main.go
NAME=image-search-service

ENTRY_CLIENT=cmd/client/main.go
NAME_CLIENT=image-search-client

ENTRY_AMQP=cmd/imagesearch_amqp/main.go
NAME_AMQP=image-search-service-amqp

ENTRY_AMQP_RPC=cmd/imagesearch_amqp_rpc/main.go
NAME_AMQP_RPC=image-search-service-amqp

# Go deps
dep_refresh:
	@printf "[+] Updating deps..\n "
	@dep ensure -update
	@printf "[+] Done!\n"

dep_add:
	@printf "[+] Adding dep..\n"
	@dep ensure -add $(src)
	@printf "[+] Done!\n"

# Protobuf complie
compile_proto:
	@printf "[+] Compiling protos.. "
	@protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. ./proto/*.proto
	@printf "[+] Done!\n"

#Build, Tests and Run
build:
	@printf "[+] Bulding go service.. "
	@mkdir -p bin
	@go build -o bin/$(NAME) $(ENTRY)
	@go build -o bin/$(NAME_CLIENT) $(ENTRY_CLIENT)
	@printf "Done!\n"

tests:
	@go test -v -short

run_server:
	@printf "[+] Running go service.. "
	@./bin/$(NAME)

run_client:
	@printf "[+] Running go client.. "
	@./bin/$(NAME_CLIENT)

run_server_amqp:
	go run $(ENTRY_AMQP) --broker=rabbitmq --broker_address=amqp://rabbitmq:rabbitmq@localhost

run_server_amqp_rpc:
	go run $(ENTRY_AMQP_RPC)