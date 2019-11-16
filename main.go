// https://golang.org/doc/code.html
package main

import "fmt"

func main() {
  fmt.Println("Hello, world.")
  writeChannel, err := serial_init()
  if err != nil {

  }
  hc_main(writeChannel)
}
