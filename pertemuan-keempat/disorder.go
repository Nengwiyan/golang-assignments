package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	var dataSatu interface{}
// 	var dataDua interface{}

// 	dataSatu = []string{"coba1", "coba2", "coba3"}
// 	dataDua = []string{"bisa1", "bisa2", "bisa3"}

// 	for i := 1; i <= 4; i++ {
// 		wg.Add(2)
// 		go cetak(dataSatu, i)
// 		go cetak(dataDua, i)
// 	}

// 	wg.Wait()
// }

// func cetak(d interface{}, num int) {
// 	defer wg.Done()
// 	fmt.Println(d, num)
// }
