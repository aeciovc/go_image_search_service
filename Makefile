
# go-micro framework
ENTRY_MICRO_SERVER=cmd/imagesearch_micro/main.go
NAME=image-search-service

ENTRY_MICRO_CLIENT=cmd/imagesearch_micro_client/main.go
NAME_CLIENT=image-search-client

# rabbitmq framework
ENTRY_AMQP_RPC=cmd/imagesearch_amqp_rpc/main.go

# Go deps
dep_refresh:
	@printf "[+] Updating deps..\n "
	@dep ensure -update
	@printf "[+] Done!\n"

dep_add:
	@printf "[+] Adding dep..\n"
	@dep ensure -add $(src)
	@printf "[+] Done!\n"

# Protobuf compile
compile_proto:
	@printf "[+] Compiling protos.. "
	@protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. ./proto/*.proto
	@printf "[+] Done!\n"

#Build
build:
	@printf "[+] Bulding go service.. "
	@mkdir -p bin
	@go build -o bin/$(NAME) $(ENTRY_MICRO_SERVER)
	@go build -o bin/$(NAME_CLIENT) $(ENTRY_MICRO_CLIENT)
	@printf "Done!\n"

#Tests
tests:
	@go test ./... -v -short

#Run commands
run_client_micro:
	@printf "[+] Running go client.. "
	@./bin/$(NAME_CLIENT)

run_server_micro:
	@printf "[+] Running go service.. "
	@./bin/$(NAME)

run_server_amqp_rpc:
	go run $(ENTRY_AMQP_RPC)