.PHONY : test build clean format build-lib

LIB_FOLDER='hello'

build-lib:
	@echo 'run cmake' \
	&& cd $(LIB_FOLDER) && rm -rf build \
	&& mkdir build \
	&& cd build \
	&& cmake .. \
	&& make \
	&& ls \
	&& cd ../..
	
build:
	go build github.com/telkomdev/satudua/cmd/satudua