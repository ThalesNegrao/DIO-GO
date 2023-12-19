/* Criar um código para imprimir todos os numeros de
0 a 100 divisivel por 3.*/

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d, ", i)
		}
	}
}
