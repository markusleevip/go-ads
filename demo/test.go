package main

import (
	"fmt"
	"time"

	ads "github.com/markusleevip/go-ads"
)

func main() {
	// call HostInit once
	err := ads.HostInit()
	if err != nil {
		fmt.Println(err)
	}

	// create new ads with wanted busName and address.
	var ads1 *ads.ADS
	ads1, err = ads.NewADS("I2C1", 0x48, "")
	if err != nil {
		fmt.Println(err)
	}

	ads1.SetConfigGain(ads.ConfigGain1)

	for true {
		// read retry from ads chip
		//var result uint16
		var results = new([4]uint16)
		for i := 0; i < 4; i++ {
			if i == 0 {
				ads1.SetConfigInputMultiplexer(ads.ConfigInputMultiplexerSingle0)
			} else if i == 1 {
				ads1.SetConfigInputMultiplexer(ads.ConfigInputMultiplexerSingle1)
			} else if i == 2 {
				ads1.SetConfigInputMultiplexer(ads.ConfigInputMultiplexerSingle2)
			} else if i == 3 {
				ads1.SetConfigInputMultiplexer(ads.ConfigInputMultiplexerSingle3)
			}
			results[i], err = ads1.Read()
			if err != nil {
				ads1.Close()
				fmt.Println(err)
			}
		}
		fmt.Printf("v1:%v v2:%v v3:%v  v4:%v\n", results[0], results[1], results[2], results[3])

		time.Sleep(100 * time.Millisecond)
	}

	// close ads bus
	err = ads1.Close()
	if err != nil {
		fmt.Println(err)
	}

}
