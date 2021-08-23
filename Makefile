.PHONY: build
build: vendor-proto .generate .build

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-problem-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-problem-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-problem-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-problem-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-problem-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-problem-api/ocp-problem-api.proto
		mv pkg/ocp-problem-api/github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api/* pkg/ocp-problem-api/
		rm -rf pkg/ocp-problem-api/github.com
		mkdir -p cmd/ocp-problem-api

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-problem-api cmd/ocp-problem-api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-problem-api
		cp api/ocp-problem-api/ocp-problem-api.proto vendor.protogen/api/ocp-problem-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate
		go mod tidy

.PHONY: docker-start
docker-start:
		docker ps | grep ocp_problem_api && docker stop ocp_problem_api || echo 'clean'
		docker build -t alex19pov31/ocp-problem-api .
		docker run -ti --rm --name=ocp_problem_api -p 8082:8082 -p 8083:8083 alex19pov31/ocp-problem-api
