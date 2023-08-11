package main

// type Student struct {
// 	id        string
// 	nama      string
// 	alamat    string
// 	pekerjaan string
// 	alasan    string
// }

// func main() {
// 	var students = []Student{
// 		{id: "1", nama: "Neneng", alamat: "jl.abc", pekerjaan: "backend", alasan: "alasan neneng"},
// 		{id: "2", nama: "Asrie", alamat: "jl. bca", pekerjaan: "backend", alasan: "alasan asrie"},
// 		{id: "3", nama: "Nana", alamat: "jl.xyz", pekerjaan: "backend", alasan: "alasan nana"},
// 		{id: "4", nama: "Tara", alamat: "jl.zyx", pekerjaan: "backend", alasan: "alasan tara"},
// 	}

// 	arg := os.Args
// 	lenArg := len(arg)

// 	// catatan penting
// 	//panjang arg akan bernilai 1 jika kita memasukan kode ini di cli "go run biodata.go"
// 	//panjang arg akan bernilai 2 jika kita me memasukan argumen di cli seperti ini "go run biodata.go 2" atau "go run biodata.go Neneng" >> neneng dan 2 adalah argumen
// 	if lenArg < 2 {
// 		fmt.Println("Silahkan masukkan nama atau nomor absen")
// 		fmt.Println("Contoh: 'go run biodata.go Neneng' atau 'go biodata.go 1'")
// 	} else {
// 		tampilData := cariData(students, arg[1])
// 		if tampilData != nil {
// 			fmt.Println("ID: ", *&tampilData.id)
// 			fmt.Println("Nama: ", *&tampilData.nama)
// 			fmt.Println("Alamat: ", *&tampilData.alamat)
// 			fmt.Println("Pekerjaan: ", *&tampilData.pekerjaan)
// 			fmt.Println("Alasan: ", *&tampilData.alasan)
// 		} else {
// 			fmt.Printf("Data tidak ditemukan!")
// 		}
// 	}
// }

// func cariData(students []Student, key string) *Student {

// 	for i := range students {
// 		if students[i].nama == key || students[i].id == key {
// 			return &students[i]
// 	}
// 	return nil

// }
