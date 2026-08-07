// go-wc 统计文本的行数、词数、字节数（类似 wc）。
// 支持 -l/-w/-c 单独显示，默认三项都显示。从文件或标准输入读。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// count 对一段文本做行/词/字节统计，返回三项计数。
func count(text string, data []byte) (lines, words, bytesN int) {
	lines = strings.Count(text, "\n")
	if len(text) > 0 && !strings.HasSuffix(text, "\n") {
		lines++
	}
	sc := bufio.NewScanner(strings.NewReader(text))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words++
	}
	bytesN = len(data)
	return
}

func main() {
	lF := flag.Bool("l", false, "显示行数")
	wF := flag.Bool("w", false, "显示词数")
	cF := flag.Bool("c", false, "显示字节数")
	flag.Parse()

	var r io.Reader = os.Stdin
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, "打开失败:", err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	}

	data, _ := io.ReadAll(r)
	text := string(data)
	lines, words, bytesN := count(text, data)

	showAll := !*lF && !*wF && !*cF
	var parts []string
	if *lF || showAll {
		parts = append(parts, fmt.Sprintf("%d", lines))
	}
	if *wF || showAll {
		parts = append(parts, fmt.Sprintf("%d", words))
	}
	if *cF || showAll {
		parts = append(parts, fmt.Sprintf("%d", bytesN))
	}
	fmt.Println(strings.Join(parts, " "))
}
