package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times sequentially", func(t *testing.T) {
		counter:=NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t,counter,3)

	})

	t.Run("incrementing counter concurently",func(t *testing.T) {
		noOfIncrements:=1000
		counter:=NewCounter()

		var wg sync.WaitGroup
		wg.Add(noOfIncrements)

		for i:=0;i<noOfIncrements;i++{
			go func ()  {
				counter.Inc()
				wg.Done()			
			}()
		}
		wg.Wait()

		assertCounter(t,counter,noOfIncrements)
		
	})
}

func assertCounter(t testing.TB,got *Counter,want int){
	t.Helper()
	if(got.Value()!=want){
		t.Errorf("got %d want %d",got.Value(),want)
	}
}