#export VERSION = $(shell bash scripts/version.sh) # The single, trailing blank is essential
#export DATE    = $(shell bash scripts/date.sh) #    The single, trailing blank is essential

EXE = sus

all: $(EXE).go
	go build $(EXE).go
$(EXE).go: $(EXE).org
	bash ../scripts/org2nw $(EXE).org | notangle -R$(EXE).go | gofmt > $(EXE).go
	cp $(EXE).go ../
test: $(EXE).go $(EXE)_test.go
	go test -v
$(EXE)_test.go: $(EXE).org
	bash ../scripts/org2nw $(EXE).org | notangle -R$(EXE)_test.go | gofmt > $(EXE)_test.go
clean:
	rm -f $(EXE) *.go
doc:
	make -C doc
