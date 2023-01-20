all: 
	make -C src
clean:
	make clean -C src
	make clean -C doc
.PHONY: doc
doc:
	make -C doc
