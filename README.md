# Latihan modul 3
Penjelasan singkat pada tiap latihan soal modul 3

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `import math` untuk mengimpor modul math yang menyediakan fungsi dan konstanta matematika
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 1
```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```
Operasi diatas merupakan operasi menghitung faktorial dari sebuah bilangan bulat dengan cara fungsi memanggil dirinya sendiri untuk menghitung faktorial secara berulang

```go
func permutation(n, r int) int {
    return factorial(n) / factorial(n-r)
}
```
Operasi diatas merupakan operasi permutasi dengan rumus n! / (n-r)!. Untuk menghitung faktorial dipangillah fungsi factorial sebelumnya 

```go
func combination(n, r int) int {
    return factorial(n) / (factorial(r) * factorial(n-r))
}
```
Operasi diatas merupakan operasi kombinasi dengan rumus n! / (r! * (n-r)!). untuk menghitung faktorial dipangillah fungsi factorial sebelumnya

## Soal 2
```go
func f(x int) int {
    return x * x
}

func g(x int) int {
    return x - 2
}

func h(x int) int {
    return x + 1
}
```
Operasi diatas merupakan operasi Fungsi-fungsi Matematika yang terdiri dari : f(x): Fungsi kuadrat yang mengembalikan x^2, g(x): Fungsi linier yang mengembalikan x - 2, h(x): Fungsi linier yang mengembalikan x + 1

```go
func fogoh(x int) int {
    return f(g(h(x)))
}

func gohof(x int) int {
    return g(h(f(x)))
}

func hofog(x int) int {
    return h(f(g(x)))
}
```

Operasi diatas merupakan program Fungsi Komposisi yang terdiri dari : fogoh(x): Fungsi komposisi f(g(h(x))). Artinya, terlebih dahulu x diinput ke fungsi h, hasilnya diinput ke fungsi g, dan akhirnya hasilnya diinput ke fungsi f. gohof(x): Fungsi komposisi g(h(f(x))). hofog(x): Fungsi komposisi h(f(g(x))).

## Soal 3
```go
func jarak(a, b, c, d float64) float64 {
        return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}
```

Operasi diatas merupakan operasi  untuk menghitung jarak antara dua titik di bidang Cartesian. dengan hasil kembalian adalah jarak antara (a,b) dan (c,d)

```go
func didalam(cx, cy, r, x, y float64) bool {
        return jarak(cx, cy, x, y) <= r
}
```

Operasi diatas merupakan operasi menentukan apakah suatu titik berada di dalam lingkaran dengan hasil apabila titik berada dalam lingkaran maka hasilnya true dan jika tidak maka hasilnya false

`
