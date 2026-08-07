package main

import "testing"

func TestCount(t *testing.T) {
	lines, words, bytesN := count("a b c\nx y\n", []byte("a b c\nx y\n"))
	if lines != 2 {
		t.Fatalf("行数期望 2 实际 %d", lines)
	}
	if words != 5 {
		t.Fatalf("词数期望 5 实际 %d", words)
	}
	if bytesN != 10 {
		t.Fatalf("字节数期望 10 实际 %d", bytesN)
	}
}

func TestCountNoTrailingNewline(t *testing.T) {
	lines, _, _ := count("hello", []byte("hello"))
	if lines != 1 {
		t.Fatalf("无结尾换行应算 1 行，实际 %d", lines)
	}
}
