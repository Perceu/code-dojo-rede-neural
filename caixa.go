/*
Desenvolva um programa que simule a entrega de notas quando um cliente efetuar um saque em um caixa eletrônico. Os requisitos básicos são os seguintes:

    Entregar o menor número de notas;
    É possível sacar o valor solicitado com as notas disponíveis;
    Saldo do cliente infinito;
    Quantidade de notas infinito;
    Notas disponíveis de R$ 100,00; R$ 50,00; R$ 20,00 e R$ 10,00

Exemplos:

    Valor do Saque: R$ 30,00 – Resultado Esperado: Entregar 1 nota de R$20,00 e 1 nota de R$ 10,00.
    Valor do Saque: R$ 80,00 – Resultado Esperado: Entregar 1 nota de R$50,00 1 nota de R$ 20,00 e 1 nota de R$ 10,00. 

*/	

package main

func Saque(x int) [4]int {
    notas:= [4]int{0, 0, 0, 0} 
    for x >= 10 {
        var quantidade int = x / 100
        x = x % 100
        notas[3] = quantidade

        quantidade = x / 50
        x = x % 50
        notas[2] = quantidade

        quantidade = x / 20
        x = x % 20
        notas[1] = quantidade   

        quantidade = x / 10
        x = x % 10
        notas[0] = quantidade
	}
    if x != 0{
        notas = [4]int{-1, -1, -1, -1} 
    }
    return notas
}

func main() {
    Saque(10)
}
