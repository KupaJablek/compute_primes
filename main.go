package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Container struct {
	mu          sync.Mutex
	prime_sum   uint64
	prime_count uint64
}

func compute_range(low, high int, c *Container) {
	flag, i := 0, 0

	var sum uint64
	var count uint64

	for low < high {
		flag = 0

		if low <= 1 {
			low++
			continue
		}

		for i = 2; i <= low/2; i++ {
			if low%i == 0 {
				flag = 1
				break
			}
		}

		if flag == 0 {
			count++
			sum += uint64(low)
		}

		low++
	}
	c.inc(sum, count)
}

func print_sum(sum, count uint64) {
	fmt.Printf("Sum of primes: [%d], Numbers of primes found: [%d]\n", sum, count)
}

func (c *Container) inc(sum, count uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prime_count += count
	c.prime_sum += sum
}

func main() {
	args := os.Args

	switch len(args) {
	case 3:
		break
	default:
		fmt.Println("!!ERROR!! Expected usage ./compute_primes [threads] [max prime]")
		return
	}
	//num_threads
	num_threads, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("ERROR: threads must be an integer")
		return
	}

	// max range
	max_range, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("ERROR: max prime must be an integer")
		return
	}

	if max_range < 0 {
		fmt.Println("ERROR: max prime must be positive integer")
		return
	}

	if num_threads < 0 {
		fmt.Println("ERROR: num threads must be positive integer")
		return
	}

	c := Container{
		prime_sum:   0,
		prime_count: 0,
	}

	if max_range < num_threads {
		for i := range max_range {
			compute_range(i, i+1, &c)
		}
		return
	}

	jobs_per_thread := max_range / num_threads
	spill_over := max_range % num_threads

	low, high := 0, jobs_per_thread

	for i := range num_threads {

		if i == num_threads-1 {
			high += spill_over
		}

		compute_range(low, high, &c)
		low = high
		high += jobs_per_thread
	}

	print_sum(c.prime_sum, c.prime_count)
}
