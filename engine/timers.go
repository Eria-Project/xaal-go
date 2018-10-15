// Package engine : Timers
package engine

import (
	"errors"
	"time"
)

var timers []timer // functions to call periodic

type timer struct {
	function func()
	period   int // seconds between 2 calls
	repeat   int // -1 if infinit
}

func (t *timer) start() {
	ticker := time.NewTicker(time.Duration(t.period) * time.Second)
	repeat := t.repeat
	t.function() // Run on ticker launch
	for range ticker.C {
		if repeat > 0 {
			repeat--
		}
		if repeat == 0 {
			ticker.Stop()
		}
		if repeat != 0 {
			t.function()
		}
	}
}

// AddTimer : Add a timer to the list
func AddTimer(f func(), period int, repeat int) int {
	t := timer{f, period, repeat}
	timers = append(timers, t)
	return len(timers) - 1 // Return the array index, for future deletion
}

// RemoveTimer : Remove a timer from the list
func RemoveTimer(index int) error {
	if index < 0 || index >= len(timers) {
		return errors.New("Index not found")
	}
	// Delete index (See https://github.com/golang/go/wiki/SliceTricks)
	timers = append(timers[:index], timers[index+1:]...)
	return nil
}

func processTimers() {

	//	expire_list = []

	for _, t := range timers {
		go t.start()
	}
}
