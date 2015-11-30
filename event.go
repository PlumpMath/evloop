// Copyright 2015 By Jash. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// mail: shijian0912@163.com

package evloop

import "time"

type Event struct {
	Function       func()
	DelayedRunTime time.Time
}

func NewEvent(function func(), delayed time.Duration) *Event {
	var event = new(Event)
	event.Function = function
	event.DelayedRunTime = time.Now().Add(delayed)
	return event
}

func (e *Event) Prior(a Interface) bool {
	var ea = a.(*Event)
	return e.DelayedRunTime.Before(ea.DelayedRunTime)
}
