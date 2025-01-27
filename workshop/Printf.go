package main

import "fmt"

func main() {
	var title string = "Avengers: Endgame"
	var year int = 2019
	var rating float64 = 8.4
	category := "Sci-Fi"
	var isSuperHero = true
	var fav rune = '⭐'

	fmt.Printf("เรื่อง: %s\n", title)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ้ง: %.1f\n", rating)
	fmt.Printf("ประเภท: %#v\n", category)
	fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuperHero)
	fmt.Printf("เรื่องโปรด: %c\n", fav)
}
