# backside
Backend (microservices) for sk8 app. Written in Go.

## Environments
- https://api.sk8.town

## Local development (docker)
### Requirements
- [Docker Engine](https://docs.docker.com/engine/install/) - version 20.10.8 or later
- [Docker Compose](https://docs.docker.com/compose/install/) - version 1.29.2 or later

### Set environment variables
To have fully functional development setup, you must complete following steps:

1. Create `.env` file
```bash
cd backside # if you're not here already
cp .env.example .env
```

2. Update `.env` with your keys
- TomTom
  - register free account on [https://developer.tomtom.com/user/register](https://developer.tomtom.com/user/register)
  - replace `<your_tomtom_api_key>` with your key

### Run
```
make dev
```

Backside should be available at [http://localhost:8080](http://localhost:8080). The app will reload on every file change.

### Compile proto

1. Install `protoc`:
  [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/)
   

2. Install `protoc-gen-go` and `protoc-gen-validate`:
  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
  go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
  ```

3. Run:
  ```
  make proto
  ```

### Available scripts
```
+------------------------------------+
|              /                 |   |
|             /                  |   |
|            /                   |   |
|           |                    |   |
|           |          .'        |   |
|           |        .'  |       |   |
|           \      .'     \      |   |
|            |     |       \     |   |
|            |    /         ;    |   |
|            |   |          |   /    |
|            |   |          |   |    |
|            |   |          \   |    |
|            |   :           |  |    |
|            /    \          |  |    |
| _         /     .>         |   \   |
| "88b.__  |____.'           \   |   |
|   Y88888888888Oooooo._      \__| _ |
|       (_)    'Y888888888888888888P'|
|                         (_)  sk8   |
+------------------------------------+
Available scripts:
        make dev        start development server
        make dev:clean  start development server with clean db
        make test       run tests
        make proto      compile protobuf files
```
