
include ../Makefile



# rebase path
ROOT_APP_PATH=$(PWD)

LIBPATH=$(LIB_PATH)
APPPATH=$(CENF_EX_CALPROTO_FSPATH)

# names
NAME=epub
ORG=com.example.$(NAME)
DESC='Example of $(NAME)'


print:

	@echo 
	@echo 'ROOT_APP_PATH        $(ROOT_APP_PATH)'
	@echo 'LIBPATH              $(LIBPATH)'
	@echo 'APPPATH              $(APPPATH)'

	@echo 'NAME                 $(NAME)'
	@echo 'ORG                  $(ORG)'
	@echo 'DESC                 $(DESC)'

	@echo 
	

flu-create:
	# make new
	cd $(ROOT_APP_PATH) && flutter create --org $(ORG)-i --description $(DESC) $(NAME)


hover-init:
	cd $(APPPATH) && hover init $(LIBPATH)

hover-init-clean:
	rm -rf $(APPPATH)/desktop

inject-cenf:
	# stardard cenf stuff
	cp ../../injects/pubspec.yaml $(APPPATH)
	code --goto $(APPPATH)/pubspec.yaml:1
	
	cp ../../injects/main.dart $(APPPATH)/lib
	cp ../../injects/main_desktop.dart $(APPPATH)/lib
	cp ../../injects/app.dart $(APPPATH)/lib

	
inject-icon:
	# Mobile App Icon
	# https://dev.to/rkowase/how-to-generate-flutter-app-icons-for-ios-and-android-11gc
	#code --goto $(APPPATH)/pubspec.yaml:27
	#code --goto ../../injects/fluttericons.yaml

	mkdir -p $(APPPATH)/icon
	cp ../../injects/logo.png $(APPPATH)/icon/icon.png
	cd $(APPPATH) && flutter pub get
	cd $(APPPATH) && flutter pub pub run flutter_launcher_icons:main




inject-appbar:
	# App Bar Text 
	code --goto $(APPPATH)/lib/main.dart:18
	code --goto $(APPPATH)/lib/main_desktop.dart:28

	# App Bar Logo
	code --goto $(APPPATH)/lib/src/full_page.dart:14

	# Document Text
	# get from doc.dart
	code --goto $(APPPATH)/lib/src/full_page.dart:26
	code --goto $(APPPATH)/lib/src/view.dart:14




	
	

