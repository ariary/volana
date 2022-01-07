before.build:
	go mod download && go mod vendor

build.volana:
	@echo "build in ${PWD}";go build cmd/volana/volana.go