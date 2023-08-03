package sleepsort

import (
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Integer](values []T) []T {
	var wg sync.WaitGroup
	var ordered []T

	// Launch a goroutine for each value in the slice.
	for _, value := range values {
		wg.Add(1)
		go func(value T) {
			// Sleep for the duration of current value.
			time.Sleep(time.Duration(time.Duration(value) * time.Second))

			ordered = append(ordered, value)
			wg.Done()
		}(value)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	return ordered
}
