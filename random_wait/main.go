package main

import (
	"context"
	"fmt"
	"github.com/grafana/pyroscope-go"
	"math/rand"
	"os"
	"random-wait/wait"
	"runtime"
	"strconv"
)

func main() {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "service1.wait",

		// replace this with the address of pyroscope server
		ServerAddress:     "https://profiles-prod-003.grafana.net",
		BasicAuthUser:     "1012458",
		BasicAuthPassword: "",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,

		// you can provide static tags via a map:
		Tags: map[string]string{
			"hostname":           os.Getenv("HOSTNAME"),
			"service_git_ref":    "d374cfb406d9debbfc4769d5b2e92630732b3152",
			"service_repository": "https://github.com/alsoba13/random_wait",
			"service_path":       "/",
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
		number := rand.Intn(5)
		pyroscope.TagWrapper(context.Background(), pyroscope.Labels("number", strconv.Itoa(number)), func(c context.Context) {
			fmt.Println(number)
			wait.Wait(number)
		})
	}
}
