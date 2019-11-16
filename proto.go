// https://github.com/golang/protobuf

package main

import (
	"log"
	"github.com/golang/protobuf/proto"
)

func encode() {
		//led1 := &LedState{ Color: Color_PURPLE }
    led_string := &LedString{ Leds: []*LedState{ &LedState{ Color: Color_PURPLE } } }
    //led_string := &LedString{ }

		control := &ControlMessage{ LedString: led_string }
		data, err := proto.Marshal(control)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		log.Print("control msg: %q", data)
}
