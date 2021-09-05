# backside
Backend (microservices) for sk8 app. Written in Go.

## Environments
- https://api.sk8.town

## Local development (docker)
### Requirements
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

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
  - replace `<your_tomtom_api>` with your key in `.env` file

### Run
```
docker-compose -f docker-compose.dev.yml up -d
```

> *NOTE:* Hot reload is not available yet.