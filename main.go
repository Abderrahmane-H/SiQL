package main

import sorm_parser "sorm/parser"

import "fmt"

func main() {
	action, tree, _ := sorm_parser.ParseToTree("select users[id, title, products[id, title]]")
	fmt.Println(action)
	fmt.Println(tree)
}
