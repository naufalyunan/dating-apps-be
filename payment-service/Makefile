include .env

PROJECT_ID=date-apps-444103
SERVICE_NAME=payment-service
IMAGE_NAME=gcr.io/$(PROJECT_ID)/$(SERVICE_NAME)
DB_HOST=junction.proxy.rlwy.net
USER_SERVICE_ADDR=users-service-611320088750.asia-southeast2.run.app:443
LOG_SERVICE_ADDR=logs-service-611320088750.asia-southeast2.run.app:443

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
		--set-env-vars USER_SERVICE_ADDR=$(USER_SERVICE_ADDR) \
		--set-env-vars XENDIT_API_KEY=$(XENDIT_API_KEY) \
		--set-env-vars XENDIT_WEBHOOK_TOKEN=$(XENDIT_WEBHOOK_TOKEN) \
		--set-env-vars XENDIT_INVOICE_CALLBACK=$(XENDIT_INVOICE_CALLBACK) \
		--set-env-vars LOG_SERVICE_ADDR=$(LOG_SERVICE_ADDR) \

cloud:protoc build_push cloud_run