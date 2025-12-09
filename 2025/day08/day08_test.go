// Copyright 2025 Google LLC
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

package day08

import "testing"

var ex = []p3{
	{162, 817, 812},
	{57, 618, 57},
	{906, 360, 560},
	{592, 479, 940},
	{352, 342, 300},
	{466, 668, 158},
	{542, 29, 236},
	{431, 825, 988},
	{739, 650, 466},
	{52, 470, 668},
	{216, 146, 977},
	{819, 987, 18},
	{117, 168, 530},
	{805, 96, 715},
	{346, 949, 466},
	{970, 615, 88},
	{941, 993, 340},
	{862, 61, 35},
	{984, 92, 344},
	{425, 690, 689},
}

func TestMakeConnections(t *testing.T) {
	want1, want2 := 40, 25272
	if got1, got2 := makeConnections(ex, 10); got1 != want1 || got2 != want2 {
		t.Errorf("makeConnections(ex, 10) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
