GO_MODULE := github.com/espitman/gws-chat/pkg/protos

.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen	
else
	rm -fR ./protogen
endif




.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
  --go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
  ./proto/user-service/*.proto \

.PHONY: protoc-tag
protoc-tag:
	protoc-go-inject-tag -input="./protogen/*/*.pb.go"

.PHONY: build
build: clean protoc protoc-tag tidy


