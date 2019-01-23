package main

import (
	"flag"
	"fmt"
)

type Flag interface {
	collect()
}

type StringArg struct {
	flag string
	unsupplied string
	description string
}

type BoolArg struct {
	flag string
	unsupplied bool
	description string
}

func (s StringArg) collect() {
	parsed := flag.String(s.flag, s.unsupplied, s.description)

	fmt.Println(parsed)
}

func (b BoolArg) collect() {
	parsed := flag.Bool(b.flag, b.unsupplied, b.description)

	fmt.Println(parsed)
}

func main() {
	// metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	// uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	// flag.Parse()

	// fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *projectName, *metricPtr, *uniquePtr)
}