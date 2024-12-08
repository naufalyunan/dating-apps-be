include .env

PROJECT_ID=date-apps-444103
SERVICE_NAME=logs-service
IMAGE_NAME=gcr.io/$(PROJECT_ID)/$(SERVICE_NAME)
DB_HOST=junction.proxy.rlwy.net

protoc:
	protoc --proto_path=pb pb/*.proto --go_out=paths=source_relative:pb/generated --go-grpc_out=paths=source_relative:pb/generated
run:
	@go run main.go

build_push:
	docker build -t $(IMAGE_NAME) .
	docker push $(IMAGE_NAME)


cloud_run:
	gcloud run deploy $(SERVICE_NAME) --image $(IMAGE_NAME) --platform managed --region asia-southeast2 --allow-unauthenticated --port 8080 \
		--set-env-vars DB_HOST=$(DB_HOST) \
		--set-env-vars DB_PORT=28176 \
		--set-env-vars DB_NAME=railway \
		--set-env-vars DB_USER=postgres \
		--set-env-vars DB_PASS=$(DB_PASS) \

cloud:protoc build_push cloud_run