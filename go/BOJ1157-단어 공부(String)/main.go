package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var cnt [26]int
	max := -1
	result := 0
	flag := false
	
	for {
		var ch uint8
		fmt.Fscanf(reader, "%1c", &ch)
		
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			idx := 0
			if ch >= 'A' && ch <= 'Z' {
				idx = int(ch - 'A')
			} else if ch >= 'a' && ch <= 'z' {
				idx = int(ch - 'a')
			}
			cnt[idx] ++
			if max < cnt[idx] {
				max = cnt[idx]
				result = idx
				flag = false
			} else if max == cnt[idx] {
				flag = true
			}
		} else if ch == '\n' {
			break;
		}
	}
	
	if flag {
		fmt.Fprintf(writer, "%s", "?")
		
	} else {
		fmt.Fprintf(writer, "%c\n", result + 'A')
	}
	writer.Flush()
}