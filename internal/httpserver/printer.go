// Package httpserver ...
package httpserver

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
