postgres:
	docker run -it -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres postgres