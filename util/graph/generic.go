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

package graph

import (
	"fmt"
	"io"

	"github.com/fis/aoc/util/fn"
)

// AnyGraph represents the common subset of all graph types.
//
// It's useful for writing generic functions on both sparse and dense graphs.
// Note that there's generally some performance impact, so it's best left for non-sensitive utility functions.
type AnyGraph[It AnyIt] interface {
	Len() int
	V(string) (int, bool)
	Label(v int) string
	W(u, v int) int
	Succ(u int) It
	Next(It) It
	hasWeights() bool
}

// AnyIt represents the common interface for the graph edge iterators.
type AnyIt interface {
	Valid() bool
	At() (u, v int)
	Tail() (u int)
	Head() (v int)
}

// WriteDOT writes the graph out in GraphViz format. The `nodeAttr` and `edgeAttr` callback
// functions are optional, and can be used to add extra attributes to the node. If the callback
// returns a "label" attribute, it takes precedence over the usual node name / edge weight.
func WriteDOT[It AnyIt](g AnyGraph[It], w io.Writer, name string, directed bool, nodeAttr func(v int) map[string]string, edgeAttr func(u, v int) map[string]string) (err error) {
	fmt.Fprintf(w, "%s %s {\n", fn.If(directed, "digraph", "graph"), name)
	for v := 0; v < g.Len(); v++ {
		var attrs map[string]string
		if nodeAttr != nil {
			attrs = nodeAttr(v)
		}
		fmt.Fprintf(w, "  n%d [", v)
		writeAttrs(w, attrs, "label", fmt.Sprintf(`"%s"`, g.Label(v)))
		fmt.Fprintf(w, "];\n")
	}
	for u := 0; u < g.Len(); u++ {
		for it := g.Succ(u); it.Valid(); it = g.Next(it) {
			v := it.Head()
			if !directed && v < u {
				continue
			}
			var attrs map[string]string
			if edgeAttr != nil {
				attrs = edgeAttr(u, v)
			}
			edgeType := fn.If(directed, "->", "--")
			fmt.Fprintf(w, "  n%d %s n%d [", u, edgeType, v)
			if g.hasWeights() {
				writeAttrs(w, attrs, "label", fmt.Sprintf(`"%d"`, g.W(u, v)))
			} else {
				writeAttrs(w, attrs)
			}
			fmt.Fprintf(w, "];\n")
		}
	}
	_, err = fmt.Fprintln(w, "}")
	return err
}

func writeAttrs(w io.Writer, attr map[string]string, xattr ...string) error {
	i := 0
	for k, v := range attr {
		if err := writeAttr(w, k, v, i); err != nil {
			return err
		}
		i++
	}
	for x := 0; x+1 < len(xattr); x += 2 {
		if _, ok := attr[xattr[x]]; ok {
			continue
		}
		if err := writeAttr(w, xattr[x], xattr[x+1], i); err != nil {
			return err
		}
		i++
	}
	return nil
}

func writeAttr(w io.Writer, k, v string, i int) error {
	if i > 0 {
		_, err := fmt.Fprint(w, ",")
		if err != nil {
			return err
		}
	}
	// TODO: better marshalling
	_, err := fmt.Fprintf(w, "%s=%s", k, v)
	if err != nil {
		return err
	}
	return nil
}
