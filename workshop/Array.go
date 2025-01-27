// ให้ประกาศตัวแปรอาร์เรย์ประเภทหนัง(genres)ที่เก็บค่าเป็น "Action", “Adventure” และ “Fantasy” ตามลำดับ
// ให้อ่านค่าของอาร์เรย์แต่ละช่องเพื่อแสดงผล ให้แสดงผลทั้งหมดตามตัวอย่าง Output ด้านล่าง
// หลังจากนั้นเปลี่ยนแปลงค่าในอาร์เรย์ index ที่ 1 ให้เป็น Sci-Fi พร้อมทั้งแสดงผล เพื่อยืนยันว่าค่าเปลี่ยนจริง
// Output:
// “genres[0]: Action”
// “genres[1]: Adventure”
// “genres[2]: Fantasy”
// “genres[1]: Sci-Fi”

package main

import "fmt"

func main() {
	var genres [3]string = [3]string{"Action", "Adventure", "Fantasy"}
	l := len(genres)
	fmt.Println("Length of genres:", l)
	fmt.Println("genres[0]:", genres[0])
	fmt.Println("genres[1]:", genres[1])
	fmt.Println("genres[2]:", genres[2])

	genres[1] = "Sci-Fi"
	fmt.Println("genres[1]:", genres[1])
}
