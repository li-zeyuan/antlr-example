package main

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/li-zeyuan/antlr-example/visitor-parser"
)

/*
观察者模式：https://juejin.cn/post/7330052230472859663
*/

type Visitor struct {
	parser.BaseCalcVisitor
	stack []int
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) push(i int) {
	v.stack = append(v.stack, i)
}

func (v *Visitor) pop() int {
	if len(v.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := v.stack[len(v.stack)-1]

	// Remove the last element from the stack.
	v.stack = v.stack[:len(v.stack)-1]

	return result
}

func (v *Visitor) acceptNode(node antlr.RuleNode) interface{} {
	// 回调对应的 VisitXXX 方法
	node.Accept(v)
	return nil
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.acceptNode(ctx.Expression()) // start 只有一个节点
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	i, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		panic(err.Error())
	}

	v.push(i)
	return nil
}

func (v *Visitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	v.acceptNode(ctx.Expression(0)) // ctx.Expression(0)左节点
	v.acceptNode(ctx.Expression(1)) // ctx.Expression(1)右节点

	//push result to stack
	var t = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	switch t.GetTokenType() {
	case parser.CalcParserMUL:
		v.push(left * right)
	case parser.CalcParserDIV:
		v.push(left / right)
	default:
		panic("should not happen")
	}

	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	v.acceptNode(ctx.Expression(0))
	v.acceptNode(ctx.Expression(1))

	//push result to stack
	var t = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	switch t.GetTokenType() {
	case parser.CalcParserADD:
		v.push(left + right)
	case parser.CalcParserSUB:
		v.push(left - right)
	default:
		panic("should not happen")
	}

	return nil
}

func (v *Visitor) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
	v.acceptNode(ctx.Expression())
	return nil
}
