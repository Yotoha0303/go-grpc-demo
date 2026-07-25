.GLOAL_DEFAULT:=help

.PHONY:=help run

help:
	@echo Target: Target command
	@echo help		show Target
	@echo run		run application of local

run:
	go run main.go