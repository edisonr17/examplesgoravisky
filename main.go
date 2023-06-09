package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateSecurityGroupNames(count int) []string {
	var names []string

	for i := 100; i < count; i++ {
		name := fmt.Sprintf("securitygroup%d", i)
		names = append(names, name)
	}

	return names
}

func generateRandomNumbers(count int) []int {
	numbers := make([]int, count)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(900) + 100
	}

	return numbers
}

func main() {
	reg := 10000900
	recorrerSgArray(reg)
	recorrerSgMap(reg)
}

func recorrerSgArray(numbers int) {
	start := time.Now()
	// Llamada a la función sum y almacenamiento del resultado en la variable result
	securityGroups := generateSecurityGroupNames(999)
	number := generateRandomNumbers(numbers)
	cont := 0
	for _, sg := range securityGroups {
		//fmt.Printf("Índice: %d, Valor: %s\n", indSg, sg)
		lastThree := sg[len(sg)-3:]
		for _, num := range number {
			nint, _ := strconv.Atoi(lastThree)
			if nint == num {

			}
			cont++
		}
	}

	fmt.Println("Lineas recorridas: ", cont)
	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución recorriendo 2 arrays: %s\n", elapsed)
}

func recorrerSgMap(numbers int) {
	start := time.Now()
	securityGroups := generateSecurityGroupNames(999)
	number := generateRandomNumbers(numbers)
	cont := 0

	var mapaUno map[string][]int
	mapaUno = make(map[string][]int)

	for _, sg := range securityGroups {
		cont++
		mapaUno[sg] = []int{}
	}
	prefix := "securitygroup"
	for _, num := range number {
		cont++
		sg := prefix + strconv.Itoa(num)
		mapaUno[sg] = append(mapaUno[sg], num)
	}
	fmt.Println("Lineas recorridas: ", cont)
	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución recorriendo maps: %s\n", elapsed)

}
