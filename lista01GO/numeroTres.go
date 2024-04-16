package main

import (
    "fmt"
    "strconv"
)

func main() {
    var n1, n2, n3 int
    fmt.Println("Digite três números inteiros separados por espaço:")
    _, err := fmt.Scan(&n1, &n2, &n3)
    if err != nil {
        fmt.Println("Erro ao ler os números:", err)
        return
    }

    
    if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
        fmt.Println("DIGITO INVALIDO")
        return
    }

    
    concatenated := strconv.Itoa(n1) + strconv.Itoa(n2) + strconv.Itoa(n3)


    square := n1*n1*10000 + n2*n2*100 + n3*n3

    fmt.Printf("%s, %d\n", concatenated, square)
}
