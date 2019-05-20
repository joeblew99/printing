# github.com/SebastiaanKlippert/go-wkhtmltopdf



GOCMD=go

NAME=go-wkhtmltopdf
LIB=github.com/SebastiaanKlippert/$(NAME)
LIB_FSPATH=$(GOPATH)/src/$(LIB)



print:
	@echo 
	@echo NAME :$(NAME)
	@echo LIB :$(LIB)
	@echo LIB_FSPATH  : $(LIB_FSPATH)
	@echo


dep-all: dep-os dep dep-modules
	# get all deps

dep-os:
	# get the binary (https://github.com/wkhtmltopdf/wkhtmltopdf/releases)
	# did it manually for now.
	chmod +x $(PWD)/bin/wkhtmltox-0.12.5-1.macos-cocoa.pkg

dep:
	## install on gopath
	mkdir -p $(LIB_FSPATH)
	cd $(LIB_FSPATH) && cd .. && rm -rf $(NAME) && git clone https://$(LIB).git
	
dep-modules:
	# theirs
	cd $(LIB_FSPATH) && make deps


code:
	code $(LIB_FSPATH)

clean:
	# gopath
	rm -rf $(LIB_FSPATH)

### Dev Tools

linters-install:
	$(GOCMD) get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint: linters-install
	cd $(LIB_FSPATH) && gometalinter --vendor --disable-all --enable=vet --enable=vetshadow --enable=golint --enable=maligned --enable=megacheck --enable=ineffassign --enable=misspell --enable=errcheck --enable=goconst ./...

test:
	cd $(LIB_FSPATH) && $(GOCMD) test -cover -race ./...

bench:
	cd $(LIB_FSPATH) && $(GOCMD) test -bench=. -benchmem ./...



### Metrics

qlog:
	# for debugging...
	tail -f $(TMPDIR)/q

#### Build

run:
	go run main.go

build:
	go build main.go


