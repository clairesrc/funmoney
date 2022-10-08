# FunMoney

[![Go](https://github.com/clairesrc/funmoney/actions/workflows/go.yml/badge.svg)](https://github.com/clairesrc/funmoney/actions/workflows/go.yml) [![Deploy static content to Pages](https://github.com/clairesrc/funmoney/actions/workflows/build-frontend.yml/badge.svg)](https://github.com/clairesrc/funmoney/actions/workflows/build-frontend.yml)

Easy way to keep track of disposable income. Simple CRUD PWA using Go, MongoDB, and Docker, and vanilla JS frontend. This is a work in progress that is not yet feature-complete. I wrote this mainly to learn the Go MongoDB client, but also for use as a tenplate for small CRUD PWA apps. The data store can easily be swapped to something else, the build process is a simple shell script and everything is intended to be very easy to hack on.

## Run

Create `.env` and set these variables in it:

```
CAP=200
CURRENCY=USD
ENV=dev
MONGODB_CONNECTION_URI="mongodb://root:example@mongo:27017/?maxPoolSize=20&w=majority"
```

`CAP` is the monthly limit of disposable income you can spend per month.

`CURRENCY` is the currency you'd like to track.

`ENV` controls log levels and live-reload. Set to `prod` to disable live-reload polling and frontend logging.

`MONGODB_CONNECTION_URI` points to the MongoDB instance. Can be left default if you're just going with the instance included in `docker-compose.yml`.

Then run:

```
./build-and-run.sh
```

This will create and run a docker-compose stack on your machine.

Navigate to `http://localhost:8002` in your browser.

## Dev

For a dev shell with live reload enabled and log tail pane in a tmux envionment, run:

```
./dev.sh
```

You can then navigate to `http://localhost:8002` in your browser. If you're working from a remote machine, append `?hostname=host.name.ip.here:8082` to the URL so the frontend knows where to look.

## Github Pages action

This project leverages Github Actions to automatically build & host the frontend code which is then hooked up to the backend API in an arbitrary location. You can access http://claire.zone/funmoney/ for a quick look at the frontend, though it is non-functional as I'm not pointing it to a live public API instance yet. You can override the default API hostname by appending `?hostname=...` to the URL i.e. http://claire.zone/funmoney/?hostname=192.168.193.208:8082
