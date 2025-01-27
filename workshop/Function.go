// สร้างฟังก์ชันชื่อ emote และรับพารามิเตอร์หนึ่งตัวชื่อ ratings เป็นตัวแปรชนิด float64 และคืนค่าเป็น string
// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการคืนค่าคำว่า “Disappointed 😞”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการคืนค่าคำว่า “Normal 😐”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการคืนค่าคำว่า “Good 🥰”
// กรณีอื่นๆ ให้ทำการคืนค่าคำว่า “🤷🤷🤷🤷”

package main

import "fmt"

func emote(ratings float64) string {
	if ratings < 5.0 {
		return "Disappointed 😞"
	} else if ratings >= 5.0 && ratings < 7.0 {
		return "Normal 😐"
	} else if ratings >= 7.0 && ratings < 10.0 {
		return "Good 🥰"
	} else {
		return "🤷🤷🤷🤷"
	}
}

func main() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}
