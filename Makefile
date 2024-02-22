REPO ?= example.com/delay
TAG ?= v0.0.1
IMG ?= $(REPO):$(TAG)
KIND_CLUSTER_NAME ?= ik8s
BINARY ?= delay-app

docker-build:
	GOOS=linux GOARCH=amd64 go build -o .output/$(BINARY) main.go && docker build --build-arg BINARY=$(BINARY) . -t $(IMG) -f Dockerfile

deploy: docker-build
	kind load docker-image $(IMG) --name $(KIND_CLUSTER_NAME)
	cd charts && helm upgrade --install delay delay --set image.repository=$(REPO) --set image.tag=$(TAG)

undeploy:
	cd charts && helm uninstall delay