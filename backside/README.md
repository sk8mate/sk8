# backside
Backend (microservices) for sk8 app. Written in Go.

## Environments
- https://api.sk8.town

## Local development (docker)
### Requirements
- [Docker](https://docs.docker.com/engine/install/)

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
docker compose -f docker-compose.dev.yml up -d --build
```

Backside should be available at [http://localhost:8080](http://localhost:8080). The app will reload on every file change.
