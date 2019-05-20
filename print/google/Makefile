GO111MODULE=on

NAME=cloud-print-connector
LIB=github.com/google/$(NAME)
LIB_FSPATH=$(GOPATH)/src/$(LIB)

git-clone:
	go get github.com/google/cloud-print-connector/...

	#mkdir -p $(LIB_FSPATH)
	#cd $(LIB_FSPATH) && cd .. && rm -rf $(NAME) && git clone ssh://git@$(LIB).git

git-clean:
	rm -rf $(LIB_FSPATH)

code:
	code $(LIB_FSPATH)

build:
	cd $(LIB_FSPATH) && go build -o $(PWD)/bin/print .

run-util:
	# 1. init to make a config
	gcp-connector-util

run-conn:
	# 2 Run with the config
	gcp-cups-connector --config-filename $(PWD)/gcp-cups-connector.config.json --log-to-console 
	
tail-log:
	tail -f /tmp/cloud-print-connector

send-job:

try-listprinter:
	cd $(PWD)/try/listprinter && go run main.go

vagrant-win:	
	# - Freaking need windows 
	# https://app.vagrantup.com/gusztavvargadr/boxes/windows-10/
	vagrant init gusztavvargadr/windows-10
	vagrant up


	


