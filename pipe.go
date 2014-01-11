// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "reflect"
)

// A Pipe is a set of transforms being applied along the channel. We use this
// as a helper while constructing a chained pipe. It lets us use a nicer
// syntax.
type Pipe struct {
	Output chan interface{}
}

// Return a new Pipe object which echoes input to output
func NewPipe(input interface{}) *Pipe {
  // TODO: Check input is a chan here
	return &Pipe{Output: reflect.ValueOf(input).Interface().(chan interface{})}
}
