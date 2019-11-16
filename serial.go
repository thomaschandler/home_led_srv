package main

import (
  "github.com/tarm/serial"
  "log"
  "time"
)

// https://godoc.org/github.com/tarm/serial
func serial_main() {
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

       n, err = s.Write([]byte("\x12\x1e\x0a\x03\x08\xff\x01\x0a\x04\x08\x80\x80\x02\x0a\x05\x08\x80\x80\xfc\x07\x0a\x05\x08\x80\x81\x80\x04\x0a\x03\x08\xff\x01"))
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
