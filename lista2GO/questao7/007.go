package main  

import "fmt"

func main() {

	var num, sumPar, countPar, sumImpar, countImpar float64

    for {
        fmt.Scan(&num)
        if num == 0 {
            break
        }
        if num == 0 {
            sumPar += num
            countPar++
        } else {
            sumImpar += num
            countImpar++
        }
    }

    mediaPar := sumPar / countPar
    mediaImpar := sumImpar / countImpar

    fmt.Printf("MEDIA PAR = %.2f\n", mediaPar)
    fmt.Printf("MEDIA IMPAR = %.2f\n", mediaImpar)
}


