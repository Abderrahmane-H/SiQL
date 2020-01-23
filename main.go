package main

import sorm_parser "sorm/parser"

import "fmt"

func main() {
	action, tree, _ := sorm_parser.ParseToTree("select users[id,email,,password,firstname,lastname]")
	fmt.Println(action)
	fmt.Println(tree)
}
