# Latihan modul 2
Penjelasan singkat pada tiap latihan soal modul 2

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 2A_1
```go
temp = satu //Nilai satu diubah ke temp
satu = dua //Nilai dua diubah ke satu
dua = tiga //Nilai tiga diubah ke dua
tiga = temp //Nilai temp (Awalnya satu) diubah ke tiga
```
Operasi diatas merupakan operasi pertukaran nilai dari variabel satu ke variabel lainnya dengan bantuan temp (Variabel sementara)

## Soal 2A_2
```go
func kabisat(tahun int) bool {
        return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0
```
Operasi diatas merupakan operasi menentukan apakah suatu inputan user (Tahun) adalah tahun kabisat atau tidak dengan melihat beberapa syarat apakah habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100


## Soal 2A_3
```go
pangkatTiga := jariJariFloat * jariJariFloat * jariJariFloat
pangkatDua := jariJariFloat * jariJariFloat

volume := (4.0 / 3.0) * pi * pangkatTiga
luasPermukaan := 4 * pi * pangkatDua
```
Operasi diatas merupakan operasi mencari volume dan luas permukaan dari bola catatan tambahan kenapa saya menggunakan variabel pangkatTiga dan pangkatDua karena jika dikali manual dalam rumus akan error dan dalam operasi matematika pangkat yang dikerjakan terlebih dahulu


## Soal 2A_4
```go
fahrenheit := (celsius * 9/5) + 32
reamur := celsius * 4/5
kelvin := (fahrenheit + 459.67) * 5/9

hasilFahrenheit := int(fahrenheit)
hasilReamur := int(reamur)
hasilKelvin := int(kelvin)
```
Operasi diatas merupakan operasi mengubah nilai dalam program ini mengubah nilai celcius ke dalam fahrenheit dan sebgainya alasan saya mengubah hasil setiap suhu menjadi integer karena di soal hasil output menggunakan bilangan bulat

## Soal 2A_5
```go
for i := 0; i < 5; i++ {
	fmt.Scan(&angka[i])
}

for i := 0; i < 4; i++ {
	fmt.Scanf("%c", &karakter[i])
}

for _, v := range angka {
	fmt.Printf("%c", v)
}
fmt.Print()

for _, v := range karakter {
	fmt.Printf("%c", v+1)
}
fmt.Print()
```

Operasi diatas merupakan program mengubah angka dan karakter ke angka dan angka setelahnya sesuai aturan ASCII

## Soal 2B_1
```go
func main() {
        var warna [5][4]string
        var warnaIdeal = []string{"merah", "kuning", "hijau", "ungu"}
        var berhasil bool

        for i := 0; i < 5; i++ {
                fmt.Printf("Percobaan %d: ", i+1)
                for j := 0; j < 4; j++ {
                        fmt.Scan(&warna[i][j])
                }
        }

        berhasil = true
        for i := 0; i < 5; i++ {
                for j := 0; j < 4; j++ {
                        if warna[i][j] != warnaIdeal[j] {
                                berhasil = false
                                break
                        }
                }
                if !berhasil {
                        break
                }
            }

        fmt.Println("BERHASIL:", berhasil)
```

Operasi diatas merupakan operasi pengecekan urutan warna yang diinput oleh user jika semuanya sesuai maka percobaan akan mengatakan berhasil : true

## Soal 2B_2
```go
fmt.Print("N : ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("N harus bilangan bulat positif")
        return
    }

    for i := 1; i <= n; i++ {
        fmt.Printf("Bunga %d : ", i)
        fmt.Scan(&bunga)
        pita += bunga + " - "
    }

    if len(pita) > 0 {
        pita = pita[:len(pita)-3]
    }

    fmt.Println("Pita : ", pita)
```

Operasi diatas merupakan operasi menginputkan kata-kata sebanyak N (Diinput user) dan hasil penginputan akan ditampilkan dengan tambahan tanda garis `-`

```go
 for {
                fmt.Printf("Bunga %d : ", jumlahBunga+1)
                fmt.Scan(&bunga)

                if bunga == "SELESAI" {
                        break
                }

                pita += bunga + " - "
                jumlahBunga++
        }

        if len(pita) > 0 {
                pita = pita[:len(pita)-3]
        }

        fmt.Println("Pita : ", pita)
        fmt.Println("Bunga : ", jumlahBunga)
```
Berbeda dengan operasi hasil modifikasi disini proses penginputan berhenti ketika user menginputkan kata `SELESAI` setelah itu program akan menampilkan kata-kata yang diinput user dan jumlah kata yang diinput

## Soal 2B_3
```go
func main() {
    var berat1, berat2 float64

    for {
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&berat1, &berat2)

        if berat1 >= 9 || berat2 >= 9 {
            fmt.Println("Proses selesai.")
            break
        }

    }
```
Operasi diatas merupakan operasi pengecekan suatu syarat dalam program diatas user akan terus menginput angka sampai salah satu angka mencapai syarat yakni lebih dari sama dengan 9 dan program akan berhenti meminta untuk menginput

```go
for {
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&berat1, &berat2)

        if berat1 < 0 || berat2 < 0 {
            fmt.Println("Proses selesai.")
            break
        }

        if berat1+berat2 > 150 {
            fmt.Println("Proses selesai.")
            break
        }

        if berat1-berat2 >= 9 || berat2-berat1 >= 9 {
            fmt.Println("Sepeda motor Pak Andi akan oleng: true")
        } else {
            fmt.Println("Sepeda motor Pak Andi akan oleng: false")
        }
    }
```
Berbeda dengan operasi hasil modifikasi dimana program akan berhenti menginput ketika total berat = 150 dan salah satu inputan merupakan bilangan negatif. Selain itu, program dapat mengecek apakah kedua inputan mempunyai selisih sebanyak 9 atau tidak

## Soal 2B_4
```go
func f(k int) float64 {
    pembilang := float64((4*k + 2) * (4*k + 2))
    penyebut := float64((4*k + 1) * (4*k + 3))
    return pembilang / penyebut
}

func main() {
    var k int
    fmt.Print("Nilai K = ")
    fmt.Scan(&k)

    hasil := f(k)
    fmt.Printf("Nilai f(K) = %.10f\n", hasil)
}
```
Operasi diatas merupakan operasi suatu fungsi/rumus dengan hasilnya akan menampilkan 10 bilangan desimal dibelakang koma. Berbeda dengan hasil modifikasi pada soal ini yang saya tidak paham cara menjelaskannya bagaimana karna jujur saya tidak tahu simbol apa yang disamping rumus

## Soal 2C_1
```go
fmt.Print("Berat parsel (gram) : ")
        fmt.Scan(&beratParsel)

        kg := beratParsel / 1000
        gram := beratParsel % 1000

        biayaDasar := kg * 10000
        if gram >= 500 {
                biayaTambahan = gram * 5
        } else {
                biayaTambahan = gram * 15
        }

        if kg > 10 {
                biayaTambahan = 0
        }

        totalBiaya := biayaDasar + biayaTambahan
```
Operasi diatas merupakan operasi perhitungan dengan beberapa syarat seperti biaya dasar(Per-kg) 10.000, untuk gram jika lebih dari sama dengan 500 gram adalah 5 per-gram dan 15 per-gram jika kurang dari 500 gram tetapi biaya gram ini akan tidak ada jika berat parsel lebih dari 10 kg

## Soal 2C_2
```go
var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah : ")
	fmt.Scanln(&nam)
	if nam > 80 {
		//nam = " A "
	}

	if nam > 72.5 {
		//nam = " AB "
	}

	if nam > 65 {
		//nam = " B "
	}

	if nam > 57.5 {
		//nam = " BC "
	}

	if nam > 50 {
		//nam = " C "
	}

	if nam > 40 {
		//nam = " D "
	} else if nam <= 40 {
		//nam = " E "
	}
	fmt.Println("Nilai mata kuliah : ", nmk)
```
Program diatas merupakan program yang tidak bisa dieksekusi karena nilai nam yang dideklarasikan float64 malah diisi string

```go
func main() {
        var nam float64
        var nmk string

        fmt.Print("Nilai akhir mata kuliah : ")
        fmt.Scanln(&nam)

        if nam > 80 {
                nmk = "A"
        } else if nam >= 72.5 {
                nmk = "AB"
        } else if nam >= 65 {
                nmk = "B"
        } else if nam >= 57.5 {
                nmk = "BC"
        } else if nam >= 50 {
                nmk = "C"
        } else if nam >= 40 {
                nmk = "D"
        } else {
                nmk = "E"
        }

        fmt.Println("Nilai mata kuliah : ", nmk)
```
Berbeda dengan operasi hasil modifikasi yang sudah diperbaiki dengan nilai string diberikan ke nmk dan perbaikan pada percabangan sehingga nilai yang diinput oleh user hasilnya sesuai dengan yang diharapkan

## Soal 2C_3
```go
func main() {
        var bilangan int
        fmt.Print("Bilangan : ")
        fmt.Scan(&bilangan)

		if bilangan > 1 {
			fmt.Print("Faktor : ")
        	for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
					fmt.Print(i, " ")
                }
            }
        fmt.Println()
		} else {
			fmt.Println("Bilangan harus lebih dari 1")
		}
        
}
```
Operasi diatas merupakan operasi pencari bilangan faktor dari bilangan yang diinput oleh user dengan ketentuan nilai yang diinput lebih dari `1`

```go
func main() {
        var bilangan, jumlahFaktor int
		
        fmt.Print("Bilangan : ")
        fmt.Scan(&bilangan)

		if bilangan > 0 {
			fmt.Print("Faktor : ")
        	jumlahFaktor = 0
        	for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
                    fmt.Print(i, " ")
                    jumlahFaktor++
                }
        }
        fmt.Println()

		if jumlahFaktor == 2 {
			fmt.Println("Prima : true")
	} else {
			fmt.Println("Prima : false")
	}
		} else {
			fmt.Println("Bilangan harus lebih dari 0")
		}
}
```
Berbeda dengan operasi hasil modifikasi yang ada tambahan fitur pengecekan bilangan prima jika iya maka akan bernilai `true` dan ada perubahan ketentuan dimana bilangan yang diinput oleh user lebih dari `0`
