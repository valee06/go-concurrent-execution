package main

import (
	"cmp"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"slices"
	"sync"
	"time"
)

type routine struct {
	routineNum int
	sleepTime  int
}

func input(n *int, m *int) {
	flag.IntVar(n, "n", 0, "Количество горутин")
	flag.IntVar(m, "m", 0, "Максимальное время сна в мс")
	flag.Parse()
}

func sortByTime(routines *[]routine) {
	slices.SortFunc(*routines, func(a, b routine) int {
		if a.sleepTime == b.sleepTime {
			return cmp.Compare(a.routineNum, b.routineNum)
		}
		return cmp.Compare(b.sleepTime, a.sleepTime)
	})
}

func output(routines []routine) {
	for _, r := range routines {
		fmt.Printf("Номер горутины: %d, время сна: %d\n", r.routineNum, r.sleepTime)
	}
}

func main() {
	var (
		n, m int
		wg   sync.WaitGroup
		mu   sync.Mutex
	)
	input(&n, &m)
	if n < 0 || m <= 0 {
		log.Fatalf("Invalid input")
	}
	routines := make([]routine, 0, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			random := rand.IntN(m)
			time.Sleep(time.Duration(random) * time.Millisecond)
			mu.Lock()
			routines = append(routines, routine{
				routineNum: i,
				sleepTime:  random,
			})
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	sortByTime(&routines)
	output(routines)
}
