package main

import (
  "math"
)

func LedsOff(writeChannel chan []byte) {
  data, err := encode([]*LedState{})
	if err != nil {

	}
  writeChannel <- data
}

func LedsSetBrightness(writeChannel chan []byte, brightness uint32) {
  // TODO: Check brightness bounds here. Not essential as arduino impls
  gs :=	&GlobalState{ Brightness: brightness }
  data, err := encodeGlobalState(gs)
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

// TODO: Dedupe with LedsOn
func LedsOnWithHue(writeChannel chan []byte, hue uint32) {
  // 45 LEDs with custom Hue

  // Scale from [0,359] to [0,255]
  fastLedHue := uint32(math.Floor((float64(hue)/360.0) * 255))

  leds :=	[]*LedState{ }
  for i := 0; i < 45; i++ {
    leds = append(leds, &LedState{ Hue: fastLedHue })
  }
  data, err := encode(leds)
	if err != nil {

	}
  writeChannel <- data
}
