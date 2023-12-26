// Copyright 2021 Google LLC
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

// Binary genmath extracts TeX code from Markdown files and generates corresponding PNGs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/fis/aoc/util"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "usage: genmath file.md [figure]")
		os.Exit(1)
	}
	file, filter := os.Args[1], ""
	if len(os.Args) == 3 {
		filter = os.Args[2]
	}

	if err := generateAll(file, filter); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generateAll(path, filter string) error {
	snippets, err := extract(path, filter)
	if err != nil {
		return err
	}
	base := strings.TrimSuffix(filepath.Base(path), ".md")
	for _, snippet := range snippets {
		if err := generate(base, snippet.name, snippet.tags, snippet.body); err != nil {
			return err
		}
	}
	return nil
}

type snippet struct {
	name string
	tags []string
	body string
}

func extract(path, filter string) (snippets []snippet, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := util.ScanAll(f, bufio.ScanLines)
	if err != nil {
		return nil, err
	}

	for len(lines) > 0 && lines[0] != "<!--math" {
		lines = lines[1:]
	}
	if len(lines) == 0 {
		return nil, nil
	}

	for start := 0; start < len(lines); start++ {
		if lines[start] == "-->" {
			break
		}
		if !strings.HasPrefix(lines[start], "%: ") {
			continue
		}
		name := lines[start][3:]
		var tags []string
		if sep := strings.IndexByte(name, ' '); sep > 0 {
			tags = strings.Split(name[sep+1:], " ")
			name = name[:sep]
		}
		if filter != "" && name != filter {
			continue
		}
		start++
		end := -1
		for e := start; e < len(lines); e++ {
			if lines[e] == "-->" || strings.HasPrefix(lines[e], "%: ") {
				end = e
				break
			}
		}
		if end < 0 {
			return nil, fmt.Errorf("unterminated math block %q", name)
		}
		for start < end && lines[start] == "" {
			start++
		}
		for end > start+1 && lines[end-1] == "" {
			end--
		}
		snippets = append(snippets, snippet{
			name: name,
			tags: tags,
			body: strings.Join(lines[start:end], "\n") + "\n",
		})
		start = end - 1
	}

	return snippets, nil
}

const texHeader = `
\documentclass{article}
\usepackage[active,pdftext,tightpage]{preview}
\setlength{\PreviewBorder}{2ex}
\usepackage{amsmath}
%EXTRAPACKAGES
\begin{document}
\begin{preview}
`

const texFooter = `
\end{preview}
\end{document}
`

func generate(base, name string, tags []string, body string) error {
	dir, err := os.MkdirTemp(".", name+".work")
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(dir)
	}()

	var extraPackages []string
	for _, pkg := range []string{"tikz", "textcomp"} {
		if slices.Contains(tags, pkg) {
			extraPackages = append(extraPackages, fmt.Sprintf("\\usepackage{%s}\n", pkg))
		}
	}
	header := strings.Replace(texHeader, "%EXTRAPACKAGES\n", strings.Join(extraPackages, "\n"), 1)

	tex := filepath.Join(dir, "math.tex")
	body = strings.TrimPrefix(header, "\n") + body + strings.TrimPrefix(texFooter, "\n")
	if err := os.WriteFile(tex, []byte(body), 0o666); err != nil {
		return err
	}

	pdfCmd := exec.Command("pdflatex", "math")
	pdfCmd.Dir, pdfCmd.Stdout, pdfCmd.Stderr = dir, os.Stdout, os.Stderr
	if err := pdfCmd.Run(); err != nil {
		return fmt.Errorf("pdflatex math: %w", err)
	}

	gsDev := "-sDEVICE=pnggray"
	if slices.Contains(tags, "color") {
		gsDev = "-sDEVICE=png16m"
	}
	pngCmd := exec.Command(
		"gs",
		"-dSAFER", "-r160",
		gsDev, "-dGraphicsAlphaBits=4", "-dTextAlphaBits=4",
		"-o", fmt.Sprintf("../%s-%s.png", base, name),
		"math.pdf",
	)
	pngCmd.Dir, pngCmd.Stdout, pngCmd.Stderr = dir, os.Stdout, os.Stderr
	if err := pngCmd.Run(); err != nil {
		return fmt.Errorf("gs ...: %w", err)
	}

	return nil
}
