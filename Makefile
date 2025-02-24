PROJECT_ROOT := $(shell pwd)

.DEFAULT_GOAL = info
info:
	@echo "Usage:"
	@echo "      make          - show this message"
	@echo "      make gen-go   - generate golang stub codes"

.PHONY: gen-go
gen-go:
	antlr -Dlanguage=Go -visitor -no-listener -o ${PROJECT_ROOT}/parser Calc.g4
