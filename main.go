package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const x = 4096
const y = 4096

func test(threadCount int) {
	var total float64 = 0
	var minimal float64 = 9999
	fmt.Printf("%d goroutines\n", threadCount)
	for i := 0; i < 10; i++ {
		var matrix [x][y]float64
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				matrix[i][j] = rand.Float64()
			}
		}
		ch := make([]chan float64, threadCount)
		pt := x / threadCount

		start := time.Now().UTC()
		for i := 0; i < threadCount; i++ {
			ch[i] = make(chan float64)
			go func(ch chan float64, num int) {
				end := pt * (num + 1)
				var sum float64 = 0
				for j := pt * num; j < end; j++ {
					for k := 0; k < y; k++ {
						sum += matrix[j][k]
					}
				}
				ch <- sum
				close(ch)
			}(ch[i], i)
		}
		var sum float64 = 0
		for i := 0; i < threadCount; i++ {
			sum += <-ch[i]
		}
		tt := float64(time.Now().Sub(start).Nanoseconds()) / 10e6
		total += tt
		if tt < minimal {
			minimal = tt
		}
		//fmt.Printf("%1.3f \n", tt)
	}
	fmt.Printf("Average time: %1.3fms\n", total/10)
	fmt.Printf("Minimal time: %1.3fms\n\n", minimal)
}
func main() {
	rand.Seed(time.Now().UTC().UnixMilli())
	for i := 0; i <= 5; i++ {
		test(int(math.Pow(2, float64(i))))
	}
}
