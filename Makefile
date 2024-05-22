pg:
	docker run -d \
	-p 5432:5432 \
	--name postgres16 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=123456 \
	postgres:16-alpine

.PHONY: pg