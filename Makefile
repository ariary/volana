# dependency on .PHONY prevents Make from 
# thinking there's `nothing to be done`
set_opts: .PHONY
  $(eval SECRET =$(shell tr -dc A-Za-z0-9 </dev/urandom | head -c 24 ; echo ''))

before.build:
	go mod download && go mod vendor

build.volana:
	@echo "build in ${PWD}";go build cmd/volana/volana.go

build.volana-with-encryption:
	@echo "build in ${PWD}";go build -ldflags "-X main.Secret=${SECRET}" cmd/volana-encr/volana-encr.go