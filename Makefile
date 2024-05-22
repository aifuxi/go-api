# db container name
DB_CONTAINER_NAME = postgres16

# db name
DB_NAME = fuxiaochen_go_api

# db username
DB_USER = root

# db password
DB_PASSWORD = 123456

# db data source name
DB_DSN = postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable

# user input db migration name
DB_MIGRATION_NAME? =

# docker will automatic create a database use the given POSTGRES_USER name
# we will have a database called root when docker container is started
pg:
	docker run -d \
	-p 5432:5432 \
	--name $DB_CONTAINER_NAME \
	-e POSTGRES_USER=$DB_USER \
	-e POSTGRES_PASSWORD=$DB_PASSWORD \
	postgres:16-alpine

createdb:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it ${DB_CONTAINER_NAME} dropdb ${DB_NAME}

migrateup:
	migrate -path db/migration -database ${DB_DSN} -verbose up

migratedown:
	migrate -path db/migration -database ${DB_DSN} -verbose down

# automatic gen migration sql files, eg: DB_MIGRATION_NAME=test make migrategen
migrategen:
	migrate create -ext sql -dir db/migration -seq ${DB_MIGRATION_NAME}

sqlc:
	sqlc generate

.PHONY: pg createdb dropdb migrateup migratedown migrategen sqlc