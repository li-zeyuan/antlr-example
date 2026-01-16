PROJECT_ROOT := $(shell pwd)

.DEFAULT_GOAL = info
info:
	@echo "Usage:"
	@echo "      make          - show this message"
	@echo "      make gen-go-visitor   - generate golang visitor codes"
	@echo "      make gen-go-listener   - generate golang listener codes"

.PHONY: gen-go-visitor
gen-go-visitor:
	antlr -Dlanguage=Go -visitor -no-listener -o ${PROJECT_ROOT}/visitor-parser Calc.g4

.PHONY: gen-go-listener
gen-go-listener:
	antlr -Dlanguage=Go -listener -no-visitor -o ${PROJECT_ROOT}/listener-parser Calc.g4