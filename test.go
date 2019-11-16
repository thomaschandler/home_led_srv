package main

func LedsOff(writeChannel chan []byte) {
  data, err := encode([]*LedState{})
	if err != nil {

	}
  writeChannel <- data
}

func LedsOn(writeChannel chan []byte) {
  leds :=	[]*LedState{ &LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_PURPLE },
			&LedState{ Color: Color_BLUE },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_RED },
			&LedState{ Color: Color_GREEN } }
  data, err := encode(leds)
	if err != nil {

	}
  writeChannel <- data
}
