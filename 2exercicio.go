package main

import "log"

type veiculo struct {
		portas int
		cor string
	}
	type caminhonete struct {
		tracãonasquatro bool
		veiculo
	}
	type sedan struct {
		modeloluxo bool
		veiculo
	}

func main() {
	carrodotio := sedan{
		true, veiculo{4, "preto"}}
	fubicadovô := caminhonete{
		veiculo:veiculo{
			8,
			"ferrugem",
		},
	tracãonasquatro: false,
		}
		log.Println(carrodotio)
		log.Println(fubicadovô)
		log.Println(carrodotio.cor)
		log.Println(fubicadovô.tracãonasquatro)
	}

