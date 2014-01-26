// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
)

type testStringer int

func (t testStringer) String() string {
	return fmt.Sprintf("%d", t)
}
