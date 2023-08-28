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
		wg.Add(2)
		mtx.Lock()
		go cetak(dataSatu, i+1, &wg, &mtx)
		mtx.Lock()
		go cetak(dataDua, i+1, &wg, &mtx) 
	}

	wg.Wait()
}

func cetak(d interface{}, num int, wg &sync.WaitGroup, mtx *sync.Mutex) {
	fmt.Println(d, num)
	mtx.Unlock()
	wg.Done()
}
