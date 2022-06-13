# SriRejeki Backend Monorepo

## How to Deploy

copy .env.example to .env and edit it's content according to the configuration used.

```
sudo docker-compose up -d --build
```

## Repository Content

| Folder        | Description                                  |
| ------------- | -------------------------------------------- |
| `backend`     | Backend source code                          |
| `docker`      | Docker volume containing NGINX configuration |
| `mqtt-broker` | MQTT Broker source code                      |

| File                   | Description                                                  |
| ---------------------- | ------------------------------------------------------------ |
| .env                   | Environmental variable which contains the configuration used |
| docker-compose.yml     | Docker compose configuration                                 |
| Dockerfile.backend     | Dockerfile for building Backend container                    |
| Dockerfile.mqtt-broker | Dockerfile for building MQTT Broker container                |

## Tech Stack

- Docker
- PostgreSQL and Redis
- Aedes MQTT Broker
- Golang and Typescript (NodeJS)
