.PHONY: generate
generate:
	mkdir -p pkg/command_runner
	protoc --proto_path vendor.protogen --proto_path api/command_runner \
	--go_out=pkg/command_runner --go-grpc_out=pkg/command_runner \
	--grpc-gateway_out=pkg/command_runner --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt generate_unbound_methods=true \
	--validate_out lang=go:pkg/command_runner \
	api/command_runner/command_runner.proto 
	make move

.PHONY: move
move:
	mv pkg/command_runner/github.com/Artenso/command_runner/pkg/command_runner/* pkg/command_runner &&\
	rm -rf pkg/command_runner/github.com/ 

.PHONY: vendor-proto
vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi

.PHONY: install-goose
install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

LOCAL_MIGRATION_DIR=migrations
PG_DB_NAME=commands_storage
USER=postgres
PWD=postgres
LOCAL_MIGRATION_DSN='host=localhost port=5432 dbname=${PG_DB_NAME} user=${USER} password=${PWD} sslmode=disable'
M_NAME=new_column

.PHONY: local-migration-create
local-migration-create:
	goose -dir ${LOCAL_MIGRATION_DIR} -v -s postgres ${LOCAL_MIGRATION_DSN} create ${M_NAME} sql

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down

# usage make local-migration-up-to V=5
.PHONY: local-migration-up-to
local-migration-up-to:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up-to ${V}

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status