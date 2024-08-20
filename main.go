package main

import (
	"fmt"
	"github.com/grafana/pyroscope-go"
	"math/rand"
	"time"
	"os"
	"runtime"
	"context"
	"strconv"
)

func wait(number int) {
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

func get_a_number() int {
	return rand.Intn(5)
}

func main() {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "random.wait.golang.app",
		
		// replace this with the address of pyroscope server
		ServerAddress: "https://profiles-prod-003.grafana.net",
		BasicAuthUser:     "",
		BasicAuthPassword: "",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,

		// you can provide static tags via a map:
		Tags: map[string]string{
			"hostname": os.Getenv("HOSTNAME"),
			"service_git_ref":    "",
      		"service_repository": "",
		},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	for {
		number := get_a_number()
		pyroscope.TagWrapper(context.Background(), pyroscope.Labels("number", strconv.Itoa(number)), func(c context.Context) {
			fmt.Println(number)
			wait(number)
		})
	}
}
