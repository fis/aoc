#! /usr/bin/env python3
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


import aocdata
import aocplot
import sys


def main():
    redraw = False
    if len(sys.argv) > 1 and sys.argv[1] == 'redraw':
        redraw = True

    changed = False
    if aocdata.leaderboard_update() or redraw:
        aocplot.plot_leaderboard()
        changed |= True
    if aocdata.stats_update() or redraw:
        aocplot.plot_stats()
        changed |= True
    if aocdata.gobench_update() or redraw:
        aocplot.plot_gobench()
        changed |= True

    if changed or redraw:
        aocplot.generate_index()


if __name__ == '__main__':
    main()
