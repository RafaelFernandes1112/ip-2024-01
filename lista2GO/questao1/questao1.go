package main

import "fmt"

var P1, P2, P3, P4, P5, P6, P7, P8 float64
var L1, L2, L3, L4, L5 float64
var trabalhoFinal float64
var F int
var NF float64
var numeroMatricula int

func min() {

	fmt.Scan(P1, P2, P3, P4, P5, P6, P7, P8)

	fmt.Scan(L1, L2, L3, L4, L5)

	fmt.Scan(trabalhoFinal)

	MP := ((P1 + P2 + P3 + P4 + P5 + P6 + P7 + P8) / 8)

	ML := ((L1 + L2 + L3 + L4 + L5) / 5)

	NT := trabalhoFinal

	NF = 0.7*MP + 0.15*ML + 0.15*NT

	frequencia := (F / 128 * 100)

	matricula := numeroMatricula

	fmt.Scan(&MP)
	fmt.Scan(&ML)
	fmt.Scan(&NF)
	fmt.Scan(&frequencia)
	fmt.Scan(&matricula)

	fmt.Println("matricula:", numeroMatricula)
	fmt.Println("Nota Final:", NF)

	for
{
		if NF >= 6 {

			fmt.Println("Situação Final: REPROVADO POR FREQUENCIA")
		} else {
			fmt.Println("Situação Final: REPROVADO POR NOTA")}
		
		if frequencia >= 75 {

			fmt.Println("Situação Final: REPROVADO POR NOTA")

		} else {

			fmt.Println("Situação Final: REPROVADO POR NOTA E POR FREQUENCIA")
		}

	}
}
