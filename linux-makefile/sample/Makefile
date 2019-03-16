export BUILD_BASE = $(shell pwd)
MAKE= make -w

#按照次序执行
SRC_PATH += $(BUILD_BASE)/export/
SRC_PATH += $(BUILD_BASE)/lib1/
SRC_PATH += $(BUILD_BASE)/main1/

all:
	sed  -i -e "s#ROOT_PATH=.*#ROOT_PATH=${CURDIR}#" Makefile.pub
	@for SRC in $(SRC_PATH); do cd $${SRC} && $(MAKE); done

release: 
	sed  -i -e "s#ROOT_PATH=.*#ROOT_PATH=${CURDIR}#" Makefile.pub
	make clean
	@for SRC in $(SRC_PATH); do cd $${SRC} && $(MAKE); done
	tar -czvf out.tar.gz out

clean:
	sed  -i -e "s#ROOT_PATH=.*#ROOT_PATH=${CURDIR}#" Makefile.pub
	@for SRC in $(SRC_PATH); do cd $${SRC} && $(MAKE) clean; done
	rm -f out/bin/*
	rm -f out/include/*
	rm -f out/lib/*


