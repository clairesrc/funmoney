# FunMoney

Easy way to keep track of disposable income. Simple CRUD PWA using Go, MongoDB, and Docker, and vanilla JS frontend.

## Setup

Check values in `docker-compose.yml` and change value of `HOSTNAME` in `frontend/app.js` to the hostname and port the backend is running on (usually `"localhost:8082"`).

## Run

```
./build-and-run.sh
```

Navigate to `http://localhost:8002` in your browser.

## Dev

For Dockerized dev environment with live-rebuilds (live reload is a WIP) and tmux shell, run:

```
./dev.sh
```
