SHELL=/bin/bash
BUILDFILE=main.go
OUTFILE=main

all: run

compile: clean
	@go build -o $(OUTFILE) $(BUILDFILE)

run: compile
	@./main

clean:
ifneq (,$(wildcard ./$(OUTFILE)))
	@rm ./$(OUTFILE)
endif