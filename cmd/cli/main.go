package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	mqttsnclient "github.com/ryodeushii/mqttsn-bombardier/mqttsn-client"
	u "github.com/ryodeushii/mqttsn-bombardier/utils"
)

// import "github.com/spf13/cobra"

const (
	U_DICT        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	P_DICT        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CONCURRENCY   = 100
	TEST_TIME_RAW = 30
	TEST_TIME     = TEST_TIME_RAW * time.Second
	HOST          = "localhost"
	PORT          = 1883
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log := u.NewLogger()
	stats := make(map[string]float64)
	statTimes := make([]float64, 0)
	avgTime := 0.0
	stats["test_time"] = TEST_TIME_RAW
	stats["concurrency"] = float64(CONCURRENCY)
	stats["errored"] = 0
	stats["total_connections"] = 0
	var wg sync.WaitGroup
	keep_alive := TEST_TIME_RAW
	for j := 0; j < TEST_TIME_RAW; j++ {
		log.Info("Test tick", j)
		for i := 0; i < CONCURRENCY; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				username := RandomString(16, U_DICT)
				password := RandomString(16, P_DICT)
				_start := time.Now()
				err := mqttsnclient.Connect(log, username, password, HOST, PORT, &keep_alive)
				if err != nil {
					stats["errored"]++
				}
				statTimes = append(statTimes, float64(time.Since(_start).Milliseconds()))
				stats["total_connections"]++
			}()
		}
		time.Sleep(5 * time.Second)

	}

	for _, time := range statTimes {
		avgTime += time
	}
	avgTime = avgTime / float64(len(statTimes))
	wg.Wait()
	log.Info(fmt.Sprintf("Stats: %+v, average exec time %fms", stats, avgTime))
	log.Info("Test finished")

}

func RandomString(len int, dict string) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = dict[rand.Intn(len)]
	}
	return string(b)
}
