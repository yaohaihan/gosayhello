package gosayhello

import "fmt"

func SayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello %s", name)
}
