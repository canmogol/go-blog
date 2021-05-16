package blog

import "fmt"

type context struct {
}

func newContext() *context {
	return &context{}
}

func (context *context) Build() {
	fmt.Println("Context build")
}
