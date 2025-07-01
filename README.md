<div align="center">
  <a href="https://github.com/gkits/kurz">
    <img src="public/logo.png" alt="Logo" width="120" height="120">
  </a>
  <h1 align="center">Kurz</h1>
  <p align="center">
    Make your URL's kurz.
  </p>
</div>

## Getting started

### Docker

1. Pull the docker image.

```bash
docker pull ghcr.io/gkits/kurz:latest
```

2. Run the docker container.

```bash
docker run -d -p 4000:4000 ghcr.io/gkits/kurz:latest
```

### Docker Compose

1. Download the `compose.yml`.

```bash
wget https://raw.githubusercontent.com/gkits/kurz/refs/heads/main/compose.yml -o compose.yml
```

2. Run the compose stack.

```bash
docker compose up -d
```

## TODO

- [ ] Implement OAuth
- [ ] Github Actions pipeline
  - [ ] Build and release docker image
