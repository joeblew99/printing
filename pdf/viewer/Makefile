
## https://github.com/albo1337/flutter_full_pdf_viewer
PDFV_LIB_NAME=flutter_full_pdf_viewer
PDFV_LIB=github.com/albo1337/$(PDFV_LIB_NAME)
PDFV_LIB_PATH=$(GOPATH)/src/$(PDFV_LIB)
PDFV_LIB_DOC=https://$(PDFV_LIB)


help:  ## Display this help
	# from: https://suva.sh/posts/well-documented-makefiles/
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) 

print:
	@echo
	@echo /PDFV pdfveek
	@echo 'PDFV_LIB_NAME              $(PDFV_LIB_NAME)'
	@echo 'PDFV_LIB                   $(PDFV_LIB)'
	@echo 'PDFV_LIB_PATH              $(PDFV_LIB_PATH)'
	@echo 'PDFV_LIB_DOC               $(PDFV_LIB_DOC)'
	@echo

git-clone: pdfv-git-clone ### git-clone
	# all


git-clean: pdfv-git-clean ### git-clean
	# all

pdfv-git-clone: ### pdfv-git-clone
	mkdir -p $(PDFV_LIB_PATH)
	cd $(PDFV_LIB_PATH) && cd .. && rm -rf $(PDFV_LIB_NAME) && git clone ssh://git@$(PDFV_LIB).git
	
pdfv-git-clean: ### pdfv-git-clean
	rm -rf $(PDFV_LIB_PATH)

pdfv-doc-open: ### pdfv-doc-open
	open  $(PDFV_LIB_DOC)

pdfv-code-open: ### pdfv-code-open
	code $(PDFV_LIB_PATH)

pdfv-flu-mob-run: ### pdfv-flu-mob-run
	cd $(PDFV_LIB_PATH)/example && Flutter run -d all

	

