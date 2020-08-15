# glow
__glow__ is a time tracking tool, written in Go.

## Build
```bash
docker build -t glow_server .
```

## Run
```bash
docker run --rm -p 9000:9000 --name glow glow_server
```
