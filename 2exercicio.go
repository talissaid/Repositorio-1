package main

import "log"

type veiculo struct {
		portas int
		cor string
	}
	type caminhonete struct {
		tracaonasquatro bool
		veiculo
	}
	type sedan struct {
		modeloluxo bool
		veiculo
	}

func main() {
	carrodotio := sedan{
		true, veiculo{4, "preto"}}
	fubicadovo := caminhonete{
		veiculo:veiculo{
			8,
			"ferrugem",
		},
	tracaonasquatro: false,
		}
		log.Println(carrodotio)
		log.Println(fubicadovo)
		log.Println(carrodotio.cor)
		log.Println(fubicadovo.tracaonasquatro)
	}

