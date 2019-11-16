package main

func LedsOff(writeChannel chan []byte) {
  data, err := encode([]*LedState{})
	if err != nil {

	}
  writeChannel <- data
}

func LedsOn(writeChannel chan []byte) {
  // 45 LEDs. Go Blue
  leds :=	[]*LedState{ }
  for i := 0; i < 45; i++ {
    leds = append(leds, &LedState{ Color: Color_BLUE })
  }
  data, err := encode(leds)
	if err != nil {

	}
  writeChannel <- data
}
