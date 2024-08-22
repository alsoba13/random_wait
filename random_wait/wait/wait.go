package wait

import "time"

func Wait(number int) {
	if number%5 == 0 {
		start := time.Now()
		for {
			now := time.Now()
			if now.Sub(start) > time.Second {
				break
			}
		}
	} else {
		time.Sleep(time.Second)
	}
}
