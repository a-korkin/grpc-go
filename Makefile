include .env

proto:
	protoc --go_out=./common --go_opt=paths=source_relative \
		--go-grpc_out=./common --go-grpc_opt=paths=source_relative \
		notes.proto
init_db:
	PGPASSWORD=${DB_ADM_PWD} psql -h localhost -U ${DB_ADM_USR} -d ${DB_NAME} \
		-v DB_NAME=${DB_NAME} -v DB_USR=${DB_USR} -v DB_PWD=${DB_PWD} \
		< ./scripts/init.sql
run:
	go run main.go
