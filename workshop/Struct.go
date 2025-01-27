// Workshop: struct
// กำหนด: 1. ให้นิยามโครงสร้างข้อมูล movie เพื่อเก็บ ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool)
// 	  2. ให้ประกาศตัวแปรเพื่อเก็บหนัง Avengers: Endgame, ปี:2019, เรตติ้ง:8.4, ประเภท:["Action", "Drama"] และ isSuperHero:true
//	  3. ให้ประกาศตัวแปรเพื่อเก็บหนัง Infinity War, ปี:2018, เรตติ้ง:8.4, ประเภท:["Action", "Sci-Fi"] และ isSuperHero:true
//	  4. ให้เก็บหนังทั้งสองเรื่องไว้ในตัวแปร mvs
// 	  5. ทำการวนลูปเพื่อแสดงผลหนังทีละเรื่อง
//
// Output:
// main.movie{title:"Avengers: Endgame", year:2019, rating:8.4, genres:[]string{"Action", "Drama"}, isSuperHero:true}
// main.movie{title:"Avengers: Infinity War", year:2018, rating:8.4, genres:[]string{"Action", "Sci-Fi"}, isSuperHero:true}

package main

import "fmt"

type movie struct {
	Name        string
	Year        int
	rating      float64
	genre       []string
	isSuperHero bool
}

func main() {
	m1 := movie{
		Name:        "Avengers: Endgame",
		Year:        2019,
		rating:      8.4,
		genre:       []string{"Action", "Drama"},
		isSuperHero: true,
	}
	m2 := movie{
		Name:        "Infinity War",
		Year:        2018,
		rating:      8.4,
		genre:       []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}
	// fmt.Printf("%#v\n", m1)
	// fmt.Printf("%#v\n", m2)

	// mvs := []movie{m1, m2}
	var mvs []movie
	mvs = append(mvs, m1, m2)

	for _, mv := range mvs {
		fmt.Printf(" %#v\n", mv)
	}
}
