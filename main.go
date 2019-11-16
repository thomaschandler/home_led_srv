// https://golang.org/doc/code.html
package main

import "fmt"

func main() {
  fmt.Println("Hello, world.")
  data, err := encode()
	if err != nil {

	}
  serial_write(data)
}
