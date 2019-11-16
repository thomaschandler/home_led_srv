// https://github.com/golang/protobuf

package main

import (
	"log"
	"github.com/golang/protobuf/proto"
)

// TODO: Named return val
func encode(leds []*LedState) ([]byte, error) {
		//led1 := &LedState{ Color: Color_PURPLE }
    led_string := &LedString{ Leds: leds }
    //led_string := &LedString{ }

		control := &ControlMessage{ LedString: led_string }
		data, err := proto.Marshal(control)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		log.Print("control msg: %q", data)
	return data, nil
}

// TODO: Proper casing on func names
func encodeGlobalState(gs *GlobalState) ([]byte, error) {
		control := &ControlMessage{ GlobalState: gs }
		data, err := proto.Marshal(control)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		log.Print("gs control msg: %q", data)
	return data, nil
}
