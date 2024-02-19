package main

import (
	"fmt"
	"math"
)

func devisivelPorTres(n int) bool {
	const numero = 0

	for i := 1; i <= n; i++ {
		if math.Mod(float64(i), 3) == numero {
			fmt.Println(i)
		}
	}

	return true
}

func gamePinPan(n int) bool {
	const numero = 0

	for i := 1; i <= n; i++ {
		if math.Mod(float64(i), 3) == numero {
			fmt.Println("Pin")
		} else if math.Mod(float64(i), 5) == numero {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}

	return true

}

func main() {
	devisivelPorTres(100)
	gamePinPan(100)
}

//criar um programa na linguagem go que trabalhe com operador % e for(resto)
// voce e seus colegs querem criar um codigo que exiba todos os numeros entre 1 e 100 que sao divisiveis por 3

// você e seus colegas querem criar em formato de codigo uma antiga brincadeira, ao falar os numeros sempre que aparecer um multiplo de 3, o participante deve falar"Pin",
// e nos multiplos de 5 o partcipante deve falar "Pan".
// Entao, seu programa deve imprimir numeros de 1 a 100 e nos citados, não deve aparecer os numeros, mas sim, o que o participante deve falar.
