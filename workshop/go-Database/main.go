package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver"
)

func main() {
	//เชื่อมต่อฐานข้อมูล
	db, err := sql.Open("ramsql", "testdb")
	if err != nil {
		log.Fatal("error:", err)
		return
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS testdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL,
	year INT NOT NULL,
	rating FLOAT NOT NULL,
	isSuperHero BOOLEAN NOT NULL,
	PRIMARY KEY (id)
	);
	`
	//สร้างฐานข้อมูล
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("create table error", err)
		return
	}

	fmt.Println("table created")

	//คำสั่ง SQL
	insert := `
	INSERT INTO testdb(imdbID,title,year,rating,isSuperHero)
	VALUES (?,?,?,?,?);
	`
	//เตรียมคำสั่ง SQL
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal("prepare statement error", err)
		return
	}

	//เพิ่มข้อมูล
	r, err := stmt.Exec("tt123456", "The Avengers", 2012, 8.1, true)
	if err != nil {
		log.Fatal("insert error", err)
		return
	}

	//แสดงค่า id ที่เพิ่มล่าสุด
	l, err := r.LastInsertId()
	if err != nil {
		log.Fatal("last insert id error", err)
		return
	}
	fmt.Println("lastInsertId:", l)

	//แสดงจำนวนแถวที่เพิ่ม
	ef, err := r.RowsAffected()
	if err != nil {
		log.Fatal("rows affected error", err)
		return
	}
	fmt.Println("rowsAffected:", ef)

	//ค้นหาข้อมูล (Query)
	rows, err := db.Query(`SELECT id, imdbID, title, year, rating, isSuperHero 
	FROM testdb
	`)
	if err != nil {
		log.Fatal("query error", err)
		return
	}

	//แสดงข้อมูล
	for rows.Next() {
		var id int
		var imdbID, title string
		var year int
		var rating float64
		var isSuperHero bool

		err = rows.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
		if err != nil {
			log.Fatal("for rows error", err)
			return
		}
		fmt.Println("row:", id, imdbID, title, year, rating, isSuperHero)
	}

	//Update ข้อมูล (จากการลองทำพบว่า ramsql driver มีปัญหาในการจับคู่ประเภทข้อมูล (type) ระหว่างคำสั่ง SQL และโครงสร้างฐานข้อมูลที่มันรองรับ)
	stmt2, err := db.Prepare(`UPDATE testdb SET rating = ? WHERE imdbID = ?`)
	if err != nil {
		log.Fatal("prepare statement2 error", err)
		return
	}

	//เปลี่ยนค่า rating จาก 8.1 เป็น 9.5
	_, err = stmt2.Exec(9.5, "tt123456")
	if err != nil {
		log.Fatal("update error", err)
		return
	}

	testrow, err := db.Query(`SELECT * FROM testdb`)
	if err != nil {
		log.Fatal("query error", err)
		return
	}

	for testrow.Next() {
		var id int
		var imdbID, title string
		var year int
		var rating float64
		var isSuperHero bool

		err = testrow.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
		if err != nil {
			log.Fatal("for rows error", err)
			return
		}
		fmt.Println("rowtest:", id, imdbID, title, year, rating, isSuperHero)
	}

	//ค้นหาข้อมูล (Query)
	rowx := db.QueryRow(`SELECT id, imdbID, title, year, rating, isSuperHero
	FROM testdb WHERE imdbID = ?`, "tt123456")

	//แสดงข้อมูล
	var id int
	var imdbID, title string
	var year int
	var rating float64
	var isSuperHero bool

	err = rowx.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
	if err != nil {
		log.Fatal("scan one rowx error", err)
		return
	}
	fmt.Println("one rowx:", id, imdbID, title, year, rating, isSuperHero)

}
