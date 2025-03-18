package main

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func PrintParseTree(tree antlr.ParseTree, parser *LispyParser) {
	fmt.Println(tree.ToStringTree(nil, parser))
	printTree(tree, "", 0, parser)
}

func printTree(tree antlr.ParseTree, indent string, depth int, parser *LispyParser) {
	// 获取节点基本信息
	nodeText := antlr.TreesGetNodeText(tree, parser.RuleNames, parser)
	children := tree.GetChildren()

	// 生成节点前缀
	prefix := ""
	if depth > 0 {
		prefix += strings.Repeat("│   ", depth-1)
		prefix += "├── "
	}

	// 打印当前节点
	fmt.Printf("%s%s [%T]\n", prefix, nodeText, tree)

	// 递归打印子节点
	for i, child := range children {
		lastNode := i == len(children)-1
		newIndent := indent
		if depth > 0 {
			if lastNode {
				newIndent += "    "
			} else {
				newIndent += "│   "
			}
		}
		printTree(child.(antlr.ParseTree), newIndent, depth+1, parser)
	}
}
