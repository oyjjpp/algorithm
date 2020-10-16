// 真实的并发bug
package concurrent

import (
	"log"
	"sync"
)

// 对共享内存保护的失误
func waitBug() {
	data := [10]int{}
	var group sync.WaitGroup
	group.Add(len(data))

	for _, p := range data {
		log.Println(p)
		go func(p int) {
			log.Println(p)
			defer group.Done()
		}(p)
		group.Wait()
	}
	// group.Wait()
}
