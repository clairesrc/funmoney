# FunMoney

Easy way to keep track of disposable income. Simple CRUD PWA using Go, MongoDB, and Docker, and vanilla JS frontend. This is a work in progress that does not support adding new transactions yet. I wrote this mainly to learn the Go MongoDB client, but also for use as a tenplate for small CRUD PWA apps. The data store can easily be swapped to something else, the build process is a simple shell script and everything is intended to be very easy to hack on.

## Run

Create `.env` and set these variables in it:

```
CAP=200
CURRENCY=USD
```

`CAP` is the monthly limit of disposable income you can spend per month.
`CURRENCY` is the currency you'd like to track.

Then run:

```
./build-and-run.sh
```

Navigate to `http://localhost:8002` in your browser.

## Dev

For Dockerized dev environment with live rebuild in a tmux shell, run:

```
./dev.sh
```
