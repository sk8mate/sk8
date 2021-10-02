usage:
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

dev:
	docker-compose -f docker-compose.dev.yml up --build

dev\:clean:
	SK8_DB_DROP_ALL_TABLES=true docker-compose -f docker-compose.dev.yml up --build

test:
	export $$(grep -v \'^#\' .env.test | xargs)
	go test ./...
