package cap9exerciseslevel04

import "fmt"

func Exerc06() {

	brazilianStates := make([]string, 0, 27)

	brazilianStates = append(brazilianStates,
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará",
		"Distrito Federal", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins")

	fmt.Println(brazilianStates)
	fmt.Printf("Len: %v | Capability: %v", len(brazilianStates), cap(brazilianStates))

}
