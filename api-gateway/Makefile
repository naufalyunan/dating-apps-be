include .env

PROJECT_ID=date-apps-444103
SERVICE_NAME=api-gateway
IMAGE_NAME=gcr.io/$(PROJECT_ID)/$(SERVICE_NAME)
PAYMENT_SERVICE_URL=https://payment-service-611320088750.asia-southeast2.run.app:443
USER_SERVICE_URL=https://users-service-611320088750.asia-southeast2.run.app:443
DATE_SERVICE_URL=https://date-service-611320088750.asia-southeast2.run.app:443
PROFILE_SERVICE_URL=https://profiles-service-611320088750.asia-southeast2.run.app:443

protoc:
	protoc --proto_path=pb pb/*.proto --go_out=paths=source_relative:pb/generated --go-grpc_out=paths=source_relative:pb/generated
run:
	@go run main.go

build_push:
	docker build -t $(IMAGE_NAME) .
	docker push $(IMAGE_NAME)


cloud_run:
	gcloud run deploy $(SERVICE_NAME) --image $(IMAGE_NAME) --platform managed --region asia-southeast2 --allow-unauthenticated --port 8080 \
		--set-env-vars PAYMENT_SERVICE_URL=$(PAYMENT_SERVICE_URL) \
		--set-env-vars USER_SERVICE_URL=$(USER_SERVICE_URL) \
		--set-env-vars DATE_SERVICE_URL=$(DATE_SERVICE_URL) \
		--set-env-vars PROFILE_SERVICE_URL=$(PROFILE_SERVICE_URL) \

cloud:protoc build_push cloud_run