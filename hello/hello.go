package main

import (
  "fmt"
  "log"
  "rsc.io/quote"
  "example.com/greetings"
)

func main() {
  names := []string{"Gradys", "Pako", "Pete"}

  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  fmt.Println(quote.Go());
  fmt.Println("Hello, World!")

  message, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(message)
}
