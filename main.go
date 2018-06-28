package main

import (
	"fmt"
)

const LETTER = 0
const COLOR = 1
const NAME = 2
const SEX = 3
const CITY = 4
const COUNTRY = 5

var myArray = [][]string{
	{"a", "rojo", "ivan", "20", "H", "madrid", "españa"},
	{"b", "amarillo", "juan", "40", "M", "malaga", "españa"},
	{"c", "verde", "juan", "35", "H", "madrid", "francia"},
	{"c", "verde", "hugo", "35", "M", "barcelona", "francia"},
	{"b", "amarillo", "ivan", "35", "H", "madrid", "alemania"},
	{"f", "azul", "juan", "40", "M", "malaga", "alemania"},
	{"a", "verde", "hugo", "20", "M", "madrid", "francia"},
	{"b", "rojo", "hugo", "20", "M", "madrid", "españa"},
	{"b", "rojo", "hugo", "35", "H", "madrid", "francia"},
	{"b", "verde", "juan", "40", "M", "barcelona", "españa"},
	{"c", "azul", "ivan", "20", "H", "madrid", "españa"},
}

func main() {
	var myMap = make(map[string]map[string]map[string]map[string]string)
	fmt.Println(myArray)
}
