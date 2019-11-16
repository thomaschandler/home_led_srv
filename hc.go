// https://github.com/brutella/hc#example

package main

import (
  "log"
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
)

func hc_main() {
    // create an accessory
    info := accessory.Info{Name: "LED strip"}
    ac := accessory.NewSwitch(info)
    
    // configure the ip transport
    config := hc.Config{Pin: "00102003"}
    t, err := hc.NewIPTransport(config, ac.Accessory)
    if err != nil {
        log.Panic(err)
    }
    
    hc.OnTermination(func(){
        <-t.Stop()
    })

		ac.Switch.On.OnValueRemoteUpdate(func(on bool) {
		    if on == true {
		        log.Println("Switch is on")
						LedsOn();
		    } else {
		        log.Println("Switch is off")
						LedsOff();
		    }
		})
 
    t.Start()
}
