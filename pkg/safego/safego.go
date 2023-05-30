package safego

import "fmt"

func Go(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panice Received from safe goroutine: \n %#v \n", r)
			}
		}()

		f()
	}()
}
