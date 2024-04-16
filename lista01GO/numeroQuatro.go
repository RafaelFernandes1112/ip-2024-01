package main

import "fmt"

func main() {
    var salarioMinimo, consumoKW float32

   
    fmt.Print("Digite o valor do salário mínimo: ")
    fmt.Scan(&salarioMinimo)

    fmt.Print("Digite a quantidade de kW gasta pela residência: ")
    fmt.Scan(&consumoKW)

    
    custoKW := (salarioMinimo * 0.7) / 100

    
    custoConsumo := custoKW * consumoKW

    custoDesconto := custoConsumo * 0.9

   
    fmt.Printf("Custo por kW: R$ %.2f\n", custoKW)
    fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
    fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}