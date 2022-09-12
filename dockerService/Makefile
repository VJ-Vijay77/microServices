FRONT_END_BINARY=frontEnd

up:
	@echo "Starting MongoDb Server"
	docker compose down
	docker compose up -d
	@echo "Mongo Server Started...."

down:	
	@echo "Mongo Server Stopping...."
	docker compose down
	@echo "Mongo Server Stopped!!!"

buildFrontEnd:
	cd ../frontEndService && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/api

start: buildFrontEnd
	@echo "Starting front end"
	cd ../frontEndService && ./${FRONT_END_BINARY}
	@echo "Front End Started..."

stop:
	@echo "Stopping front end"
	-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped Front End!!"	