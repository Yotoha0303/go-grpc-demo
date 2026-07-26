.GLOAL_DEFAULT:=help

.PHONY:=help run

help:
	@echo Target: Target command
	@echo help		show Target
	@echo run		run application of local
	@echo tidy 		go auto check and install mod

run:
	go run ./cmd

tidy:
	go mod tidy