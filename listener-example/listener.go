package main

import (
	"fmt"
	"strconv"

	parser "github.com/li-zeyuan/antlr-example/listener-parser"
)

type Listener struct {
	parser.BaseCalcListener
	stack []int
}

func NewListener() *Listener {
	return &Listener{
		stack: make([]int, 0),
	}
}

func (l *Listener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *Listener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *Listener) EnterStart(ctx *parser.StartContext) {
	fmt.Println("enter start: ", ctx.GetText())
}

func (l *Listener) EnterNumber(ctx *parser.NumberContext) {
	fmt.Println("enter number:", ctx.GetText())
}

func (l *Listener) EnterMulDiv(ctx *parser.MulDivContext) {
	fmt.Println("enter mul div:", ctx.GetText())
}

func (l *Listener) EnterAddSub(ctx *parser.AddSubContext) {
	fmt.Println("enter add sub:", ctx.GetText())
}

func (l *Listener) EnterParenthesis(ctx *parser.ParenthesisContext) {
	fmt.Println("enter parenthesis:", ctx.GetText())
}

func (l *Listener) ExitStart(ctx *parser.StartContext) {
	fmt.Println("exit start")
}

func (l *Listener) ExitNumber(ctx *parser.NumberContext) {
	fmt.Println("exit number:", ctx.GetText())

	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

func (l *Listener) ExitMulDiv(c *parser.MulDivContext) {
	fmt.Println("exit mul div:", c.GetText())

	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *Listener) ExitAddSub(ctx *parser.AddSubContext) {
	fmt.Println("exit add sub:", ctx.GetText())

	right, left := l.pop(), l.pop()

	switch ctx.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (l *Listener) ExitParenthesis(ctx *parser.ParenthesisContext) {
	fmt.Println("exit parenthesis: ", ctx.GetText())
}
