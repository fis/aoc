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

package day10

// #cgo CFLAGS: -I/usr/include/lpsolve
// #cgo LDFLAGS: -llpsolve55 -lcolamd -lm
/*
#include <stdint.h>
#include <stdio.h> // XXX remove
#include "lp_lib.h"

int lpsolve_run(const uint32_t *buttons, const int (*bounds)[2], int nb, const int *jolts, int nj) {
  lprec *lp = make_lp(nj, nb);
  if (!lp)
    return -1;
  set_outputfile(lp, "");

  for (int bi = 0; bi < nb; bi++) {
    set_int(lp, 1+bi, TRUE);
    set_bounds(lp, 1+bi, bounds[bi][0], bounds[bi][1]);
  }

  REAL row[1+nb];
  for (int bi = 0; bi < nb; bi++)
    row[1+bi] = 1.0;
  set_obj_fn(lp, row);
  for (int ji = 0; ji < nj; ji++) {
    for (int bi = 0; bi < nb; bi++)
      row[1+bi] = (buttons[bi] >> ji) & 1;
    set_row(lp, 1+ji, row);
    set_constr_type(lp, 1+ji, EQ);
    set_rh(lp, 1+ji, jolts[ji]);
  }

  int result = -1;
  if (solve(lp) == OPTIMAL)
    result = (int) lrint(get_objective(lp));

  delete_lp(lp);
  return result;
}
*/
import "C"

func lpsolve(buttons []uint32, jolts []int, bounds [][2]int) (totalPresses int) {
	buttonData := make([]C.uint32_t, len(buttons))
	boundData := make([][2]C.int, len(buttons))
	for i, b := range buttons {
		buttonData[i] = C.uint32_t(b)
		boundData[i] = [2]C.int{C.int(bounds[i][0]), C.int(bounds[i][1])}
	}
	joltData := make([]C.int, len(jolts))
	for i, j := range jolts {
		joltData[i] = C.int(j)
	}
	obj := C.lpsolve_run(&buttonData[0], &boundData[0], C.int(len(buttonData)), &joltData[0], C.int(len(joltData)))
	return int(obj)
}
