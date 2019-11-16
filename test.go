package main

func LedsOff() {

}

func LedsOn() {
  data, err := encode()
	if err != nil {

	}
  serial_write(data)
}
