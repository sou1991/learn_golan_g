package main

import {
	"fmt"
	"golang.org/x/tour/wc"
}

type Coordinate struct{
	//東経・北緯
	Longitude, Latitude float64
}

func main() {
									//key  //value
	coor := map[string]Coordinate{
		"Kobe city": Coordinate{ 
			34.69139, 135.18306, 
		},
		"Naha city": Coordinate{ 
			26.2125, 127.68111, 
		},
	}

	//Keyの存在チェック
	v, ok :=coor["Kobe city"]
	fmt.Println("The value:", v, "Present?", ok)
	fmt.Println(coor)
}
