package main

import "log"

func main() {
	amigos := map[int]string {
	23: "muito legal",
	98: "bacaninha",
	73: "melhor número",
	18: "mia idade",
	}
log.Println(amigos)
	for key, v := range amigos {
		log.Println(key, v)
	}

}