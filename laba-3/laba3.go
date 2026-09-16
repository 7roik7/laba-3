package main

import (
	"fmt"

	"laba3/mathutils"
	"laba3/stringutils"
)

func main() {
	// ===== Задания 1–2: факториал =====
	var n int
	fmt.Print("Введите число для факториала: ")
	fmt.Scan(&n)
	fmt.Printf("%d! = %d\n\n", n, mathutils.Factorial(n))

	// ===== Задание 3: переворот строки =====
	fmt.Println("Переворот 'Привет':", stringutils.Reverse("Привет"))
	fmt.Println()

	// ===== Задание 4: массив из 5 чисел =====
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = (i + 1) * 10
	}
	fmt.Println("Массив:", arr)

	// ===== Задание 5: срез — добавление и удаление =====
	slice := arr[:]
	slice = append(slice, 60, 70)
	fmt.Println("После добавления:", slice)

	idx := 2
	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println("После удаления индекса 2:", slice)
	fmt.Println()

	// ===== Задание 6: самая длинная строка =====
	words := []string{"Go", "программирование", "код", "лабораторная", "срез"}
	longest := words[0]
	for _, w := range words {
		if len([]rune(w)) > len([]rune(longest)) {
			longest = w
		}
	}
	fmt.Println("Самая длинная строка:", longest)
}