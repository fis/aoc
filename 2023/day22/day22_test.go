// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package day22

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"1,0,1~1,2,1",
	"0,0,2~2,0,2",
	"0,2,3~2,2,3",
	"0,0,4~0,2,4",
	"2,0,5~2,2,5",
	"0,1,6~2,1,6",
	"1,1,8~1,1,9",
}

func TestDisintegrate(t *testing.T) {
	bricks, err := fn.MapE(ex, parseBrick)
	if err != nil {
		t.Fatal(err)
	}
	wantSafe, wantFallen := 5, 7
	if safe, fallen := disintegrate(bricks); safe != wantSafe || fallen != wantFallen {
		t.Errorf("disintegrate(ex) = (%d, %d), want (%d, %d)", safe, fallen, wantSafe, wantFallen)
	}
}
