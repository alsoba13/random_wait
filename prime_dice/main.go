package main

import (
	"fmt"
	"github.com/fxtlabs/primes"
	"github.com/grafana/pyroscope-go"
	"math/rand"
	"os"
	"random-wait/dice"
	"runtime"
	"time"
)

func get_a_number() int {
	return rand.Intn(5)
}

func main() {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "prime.dice.app",

		// replace this with the address of pyroscope server
		ServerAddress:     "https://profiles-prod-003.grafana.net",
		BasicAuthUser:     "1012458",
		BasicAuthPassword: "",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,

		// you can provide static tags via a map:
		Tags: map[string]string{
			"hostname":           os.Getenv("HOSTNAME"),
			"service_git_ref":    "5a2ad1704c1ff5deda1eba5cc7c50f4297f5d788",
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
		sum := 1
		for sum == 1 || primes.IsPrime(sum) {
			result := dice.Roll()
			sum += result
			fmt.Printf("Rolled %d. Sum: %d\n", result, sum)
		}
		fmt.Printf("%d isn't prime. Starting over!\n", sum)
		time.Sleep(time.Second)
	}
}
