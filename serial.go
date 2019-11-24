package main

import (
  "github.com/tarm/serial"
  "log"
  "time"
)

// TODO: Fix naming/spacing
func serial_init() (chan []byte, error) {
        c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
              log.Fatal(err)
        }

      time.Sleep(3 * time.Second)
      // https://gobyexample.com/channels
      writeChannel := make(chan []byte)

      go func () {
        // https://gobyexample.com/range-over-channels
        for data := range writeChannel {
          log.Print("Got data: %q", data)
           n, err := s.Write(data)
                if err != nil {
                      log.Fatal(err)
                }
                if n == 0 {
                      log.Fatal("No bytes written")
                }
        }
      }()
      return writeChannel, nil
}
