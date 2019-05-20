
## https://github.com/signintech/pdft
LIB_NAME=pdft
LIB=github.com/signintech/$(LIB_NAME)
LIB_PATH=$(GOPATH)/src/$(LIB)
LIB_DOC=https://$(LIB)


help:  ## Display this help
	# from: https://suva.sh/posts/well-documented-makefiles/
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) 

print:
	@echo
	@echo /PDFV pdfveek
	@echo 'LIB_NAME              $(LIB_NAME)'
	@echo 'LIB                   $(LIB)'
	@echo 'LIB_PATH              $(LIB_PATH)'
	@echo 'LIB_DOC               $(LIB_DOC)'
	@echo

git-clone: git-clone ### git-clone
	# all


git-clean: git-clean ### git-clean
	# all

git-clone: ### git-clone
	mkdir -p $(LIB_PATH)
	cd $(LIB_PATH) && cd .. && rm -rf $(LIB_NAME) && git clone ssh://git@$(LIB).git
	
git-clean: ### git-clean
	rm -rf $(LIB_PATH)

doc-open: ### doc-open
	open  $(LIB_DOC)

code-open: ### code-open
	code $(LIB_PATH)

go-run: ### flu-mob-run
	cd $(LIB_PATH)/example && Flutter run -d all

	

