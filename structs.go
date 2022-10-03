package main

import "log"

type pessoa struct {
	nome      string
	sobrenome string
	sabor []string
}
func main() {

	meumapa := map[string]pessoa{
		"carvalho": pessoa{
			nome:      "joão",
			sobrenome: "carvalho",
			sabor:     []string{"pistache,", "morango,", "chocolate"}},
		"eduarda": pessoa{
			nome:      "maria",
			sobrenome: "eduarda",
			sabor:     []string{"baunilha,", "flocos,", "açaí"}},
	}
	for _,valor := range meumapa{
		log.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabor{
			log.Println("-", valor)}
		log.Println("\n")
	}

}
