# go-service-monitor
Service monitor and status page.

Monitors for HTTP/HTTPS (200 OK), TCP ports and DNS IP queries.

# Requirements

- Go 1.18+. Tested on 1.20.5.
- npm. Tested on 8.1.2.

# How to use

## Setup

```
git clone https://github.com/haroflow/go-service-monitor.git

cd go-service-monitor
go run .
cp sample.config.json config.json

cd frontend
npm install
```

## Development

1. Start frontend:
	```
	cd go-service-monitor/frontend
	npm run dev
	```

2. Start backend and monitoring on port 8080:
	```
	cd go-service-monitor
	go run .
	```

3. Open frontend server in browser: `http://127.0.0.1:5173`

	P.S.: Frontend proxies requests from `/api` to `localhost:8080/api`.

## Production

1. Build frontend to `go-service-monitor/frontend/dist`:
	```
	cd go-service-monitor/frontend
	npm run build
	```

2. Build backend and embed frontend:
	```
	cd go-service-monitor
	go build .
	```

3. Run `go-service-monitor`
4. Open server in browser: `http://localhost:8080`