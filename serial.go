package main

import (
  "github.com/tarm/serial"
  "log"
  "time"
)

// https://godoc.org/github.com/tarm/serial
func serial_write(data []byte) {
        c := &serial.Config{Name: "/dev/cu.usbmodem142101", Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
              log.Fatal(err)
        }

      time.Sleep(3 * time.Second)

      // Wait for startup to complete
      // TODO: Check returned data
			buf := make([]byte, 128)
      n, err := s.Read(buf)
      if err != nil {
              log.Fatal(err)
      }
      log.Print("%q", buf[:n])

       n, err = s.Write(data)
        if err != nil {
              log.Fatal(err)
        }
        if n == 0 {
              log.Fatal("No bytes written")
        }

      // Read to stop the arduino from resetting
      n, err = s.Read(buf)
      if err != nil {
              log.Fatal(err)
      }
      log.Print("%q", buf[:n])

}
