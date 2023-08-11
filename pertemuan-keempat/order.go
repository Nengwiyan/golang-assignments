package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex

func main() {
	var dataSatu interface{}
	var dataDua interface{}

	dataSatu = []string{"coba1", "coba2", "coba3"}
	dataDua = []string{"bisa1", "bisa2", "bisa3"}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go cetak(dataSatu, dataDua, i+1)
	}

	wg.Wait()
}

func cetak(d interface{}, e interface{}, num int) {
	mtx.Lock()
	defer mtx.Unlock()

	fmt.Println(d, num)
	fmt.Println(e, num)

	wg.Done()
}
