# Docker Get URL
Docker Test

## Build Docker Image


```bash
docker build -t aditirvan/get-ip:1.0 .
```

## Run Container

```bash
docker run -d -p 80:80 --name get-client-ip docker.io/aditirvan/get-ip:1.0
```

## Access
Open your browser and type the url

```bash
http://localhost
```