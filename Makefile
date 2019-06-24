all: build down up

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

log-proxy:
	 docker logs -f trial-nginx_proxy_1

log-web:
	 docker logs -f trial-nginx_web_1

log-app:
	 docker logs -f trial-nginx_app_1
