usage::
	@echo "+------------------------------------+"
	@echo "|              /                 |   |"
	@echo "|             /                  |   |"
	@echo "|            /                   |   |"
	@echo "|           |                    |   |"
	@echo "|           |          .'        |   |"
	@echo "|           |        .'  |       |   |"
	@echo "|           \      .'     \      |   |"
	@echo "|            |     |       \     |   |"
	@echo "|            |    /         ;    |   |"
	@echo "|            |   |          |   /    |"
	@echo "|            |   |          |   |    |"
	@echo "|            |   |          \   |    |"
	@echo "|            |   :           |  |    |"
	@echo "|            /    \          |  |    |"
	@echo "| _         /     .>         |   \   |"
	@echo "| \"88b.__  |____.'           \   |   |"
	@echo "|   Y88888888888Oooooo._      \__| _ |"
	@echo "|       (_)    'Y888888888888888888P'|"
	@echo "|                         (_)  sk8   |"
	@echo "+------------------------------------+"
	@echo Available scripts:
	@echo "	make dev	start development server"
	@echo "	make dev:clean	start development server with clean db"
	@echo "	make test	run tests"
	@echo "	make proto	compile protobuf files"

dev::
	docker-compose -f docker-compose.dev.yml up --build

dev\:clean::
	SK8_DB_DROP_ALL_TABLES=true docker-compose -f docker-compose.dev.yml up --build

test::
	export $$(grep -v \'^#\' .env.test | xargs)
	go test ./...


PROTOC_GEN_VALIDATE=$$(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1
proto::
	protoc -I . -I $(PROTOC_GEN_VALIDATE) --go_out=./places --validate_out="lang=go:./places" ./places/proto/*.proto
	protoc -I . -I $(PROTOC_GEN_VALIDATE) --go_out=./spots --validate_out="lang=go:./spots" ./spots/proto/*.proto
	protoc -I . -I $(PROTOC_GEN_VALIDATE) --go_out=./auth ./auth/proto/*.proto
