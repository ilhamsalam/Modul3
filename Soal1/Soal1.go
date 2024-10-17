package main

import (
    "fmt"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func permutation(n, r int) int {
    return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
    return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
    var a, b, c, d int
    fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

    if a >= c && b >= d {
        fmt.Println(permutation(a, c), combination(a, c))
        fmt.Println(permutation(b, d), combination(b, d))
    } else {
        fmt.Println("Invalid input")
    }
}
