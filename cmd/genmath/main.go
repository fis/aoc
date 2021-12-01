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
	"strings"

	"github.com/fis/aoc/util"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: genmath file.md")
		os.Exit(1)
	}

	if err := generateAll(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generateAll(path string) error {
	snippets, err := extract(path)
	if err != nil {
		return err
	}
	base := strings.TrimSuffix(filepath.Base(path), ".md")
	for _, snippet := range snippets {
		if err := generate(base, snippet.name, snippet.body); err != nil {
			return err
		}
	}
	return nil
}

type snippet struct {
	name string
	body string
}

func extract(path string) (snippets []snippet, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := util.ScanAll(f, bufio.ScanLines)
	if err != nil {
		return nil, err
	}

	for start := 0; start < len(lines); start++ {
		if !strings.HasPrefix(lines[start], "<!--math:") {
			continue
		}
		name := lines[start][9:]
		end := -1
		for e := start + 1; e < len(lines); e++ {
			if strings.HasPrefix(lines[e], "-->") {
				end = e
				break
			}
		}
		if end < 0 {
			return nil, fmt.Errorf("unterminated math block %q", name)
		}
		snippets = append(snippets, snippet{
			name: name,
			body: strings.Join(lines[start+1:end], "\n") + "\n",
		})
		start = end
	}

	return snippets, nil
}

const texHeader = `
\documentclass{article}
\usepackage[active,pdftext,tightpage]{preview}
\setlength{\PreviewBorder}{2ex}
\usepackage{amsmath}
\begin{document}
\begin{preview}
`
const texFooter = `
\end{preview}
\end{document}
`

func generate(base, name, body string) error {
	dir, err := os.MkdirTemp(".", name+".work")
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(dir)
	}()

	tex := filepath.Join(dir, "math.tex")
	body = strings.TrimPrefix(texHeader, "\n") + body + strings.TrimPrefix(texFooter, "\n")
	if err := os.WriteFile(tex, []byte(body), 0o666); err != nil {
		return err
	}

	pdfCmd := exec.Command("pdflatex", "math")
	pdfCmd.Dir, pdfCmd.Stdout, pdfCmd.Stderr = dir, os.Stdout, os.Stderr
	if err := pdfCmd.Run(); err != nil {
		return fmt.Errorf("pdflatex math: %w", err)
	}

	pngCmd := exec.Command(
		"gs",
		"-dSAFER", "-r160",
		"-sDEVICE=pnggray", "-dGraphicsAlphaBits=4", "-dTextAlphaBits=4",
		"-o", fmt.Sprintf("../%s-%s.png", base, name),
		"math.pdf",
	)
	pngCmd.Dir, pngCmd.Stdout, pngCmd.Stderr = dir, os.Stdout, os.Stderr
	if err := pngCmd.Run(); err != nil {
		return fmt.Errorf("gs ...: %w", err)
	}

	return nil
}
