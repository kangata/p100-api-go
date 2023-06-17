# P100 API

A simple API for get device info and switch on/off device (Tapo P100)

Feature
- [x] Get device info
- [x] Swich on/off device
- [x] Auth basic

## How to run
```
curl -fsS -O https://raw.githubusercontent.com/kangata/p100-api-go/main/docker-compose.yml
```
Setup your environment and run
```
docker compose up -d
```

## Get device info
```
curl -i -X GET -H "Authorization: Basic {YOUR_TOKEN}" localhost:7300/device
```

## Switch on/off
```
curl -i -X POST -H "Authorization: Basic {YOUR_TOKEN}" -d "status=true" localhost:7300/switch
```
```
curl -i -X POST -H "Authorization: Basic {YOUR_TOKEN}" -d "status=false" localhost:7300/switch
```
