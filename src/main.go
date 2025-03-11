package main

import (
	"fmt"
)

func main() {

	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("1: Fatorial --> %d\n", fatorial(5))
	fmt.Printf("2: Sum --> %d\n", sum(5))
	fmt.Printf("3. Fibo --> %d\n", fibo(6))
	fmt.Printf("4. Somatório --> %d\n", somatorio(0, 100))
	fmt.Printf("5. Palíndromo --> %t\n", isPal("eduar"))
	fmt.Printf("6. ConvBase2 --> %s\n", convBase2(12))
	fmt.Printf("7. ArrayList Somatorio --> %d\n", arrayListSomatorio(lista))
	fmt.Printf("8. Biggest --> %d\n", findBiggest(lista))

}

//1. Modele e implemente um método recursivo que calcule o fatorial de um número n passado como parâmetro.

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

//2. Modele e implemente um método recursivo que calcule o somatório de um número n (passado como parâmetro) até 0.

func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}

//3. Modele e implemente um método recursivo que calcule o n-ésimo número da sequência de fibonacci.

func fibo(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

//4. Modele e implemente um método recursivo que calcule o somatório dos números inteiros entre os números k e j, passados como parâmetro.

func somatorio(k int64, j int64) int64 {
	if j == k {
		return j
	}
	return k + somatorio(k+1, j)
}

//5. Modele e implemente um método recursivo que recebe um String em retorna true se este String for um palíndrome, false caso contrário.

func isPal(s string) bool {
	return isPalAux(s) == s
}

func isPalAux(s string) string {
	if s == "" {
		return ""
	}
	return string(s[len(s)-1]) + isPalAux(string(s[:len(s)-1]))

}

//6. Modele e implemente um método recursivo que recebe um inteiro zero ou positivo e retorna um String com o número em binário.

func convBase2(n int) string {
	if n == 0 {
		return "0"
	}
	if n&1 == 1 {
		return convBase2(n>>1) + "1"
	}
	return convBase2(n>>1) + "0"
}

//7. Modele e implemente um método recursivo que calcule o somatório dos números contidos em um ArrayList de inteiros, passado como parâmetro.

func arrayListSomatorio(l []int) int64 {
	if len(l) == 1 {
		return int64(l[0])
	}

	return int64(l[len(l)-1]) + arrayListSomatorio(l[:len(l)-1])

}

func findBiggest(l []int) int {
	return findBiggestAux(l, l[len(l)-1])
}

func findBiggestAux(l []int, biggest int) int {
	if len(l) == 1 {
		if l[0] > biggest {
			return l[0]
		}
		return biggest
	}

	if l[len(l)-1] > biggest {
		return findBiggestAux(l[:len(l)-1], l[len(l)-1])
	}

	return findBiggestAux(l[:len(l)-1], biggest)
}
