package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCoundown(t *testing.T) {
	t.Run("test for printing and sleeping",func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
	
		Countdown(buffer,spySleeper)
	
		got:= buffer.String()
		want:= `3
2
1
Go!`
	
		if got!=want {
			t.Errorf("got %q want %q",got,want)
		}
	
		if spySleeper.Calls!=3{
			t.Errorf("Not enouogh sleep calls, want 3 got %q",spySleeper.Calls)
		}
	})
	t.Run("test sleep print sequence", func(t *testing.T) {
		spySleepPrinter:= &SpyCountdownOperations{}
		Countdown(spySleepPrinter,spySleepPrinter)

		want:=[]string{write,sleep,write,sleep,write,sleep,write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls){
			t.Errorf("wanted calls %v got %v",want,spySleepPrinter.Calls)
		}

	})
	
}