package main

import(
	"fmt"
	"os"
	"container/list"
	"bufio"
	"strconv"
)

type position struct {
	pos int
	cnt int
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	
	r.Scan()
	K, _ := strconv.Atoi(r.Text())
	
	q := list.New()
	
	var visit [100001] bool
	q.PushBack(position{N, 0})
	visit[N] = true
	
	for q.Len() > 0 {
		
		now := q.Front().Value.(position)
		q.Remove(q.Front())
		
		if now.pos == K {
			fmt.Fprintf(w, "%d\n", now.cnt)
			break
		}
		
		for i := 0; i < 3; i ++ {
			var next int
			if i == 0 {
				next = now.pos - 1
			} else if i == 1 {
				next = now.pos + 1
			} else if i == 2 {
				next = now.pos * 2
			}
			
			if next >= 0 && next <= 100000 && !visit[next] {
				q.PushBack(position{next, now.cnt + 1})
				visit[next] = true
			}
		}
	}
	w.Flush()
}