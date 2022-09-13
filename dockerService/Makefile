FRONT_END_BINARY=frontEnd
POSTGRES_BINARY=PostgresB
MONGO_BINARY=MongoB

up:buildFrontEnd buildPostgres buildMongo
	@echo "Starting  Services"
	docker compose up -d
	@echo "Services Started...."

down:	
	@echo "Services Stopping...."
	docker compose down
	@echo "Services Stopped!!!"

buildFrontEnd:
	cd ../frontEndService && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/api
buildPostgres:
	cd ../postgresService && env GOOS=linux CGO_ENABLED=0 go build -o ${POSTGRES_BINARY} ./cmd/api
runPostgres:
	cd ../postgresService && ./${POSTGRES_BINARY}
buildMongo:
	cd ../mongoService && env GOOS=linux CGO_ENABLED=0 go build -o ${MONGO_BINARY} ./cmd/api



start: buildFrontEnd
	@echo "Starting front end"
	cd ../frontEndService && ./${FRONT_END_BINARY}
	@echo "Front End Started..."

stop:
	@echo "Stopping front end"
	-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped Front End!!"	