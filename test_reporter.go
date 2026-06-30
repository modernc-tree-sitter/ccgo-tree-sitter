package main

import (
    "fmt"
    "github.com/lucasew/ccgo-tree-sitter/internal/reporter"
)

func main() {
    reporter.HandleError(fmt.Errorf("test error"))
}
