package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/uts")
	err = db.Ping()
	if err != nil {
		panic("Gagal menghubungkan ke database . . . ")
	}
	defer db.Close()

	router := gin.Default() // routing / menjalankan framework gin

	// router.Run(":8080")

	type Progdi struct {
		Id       int    `json: "id"`
		Jenjang  string `json: "jenjang"`
		NmProgdi string `json: "nmprodi"`
	}

	// Menampilkan data berdasarkan id
	router.GET("/:id", func(c *gin.Context) {
		var (
			progdi Progdi
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, jenjang, nmprodi from progdi where id=?;", id)
		err = row.Scan(&progdi.Id, &progdi.Jenjang, &progdi.NmProgdi)
		if err != nil {
			// jika tidak ada result yand di kirim
			result = gin.H{
				"Hasil":  "Yah tidak ada yang di temukan",
				"Jumlah": 1,
			}
		} else {
			result = gin.H{
				"Hasil":  progdi,
				"Jumlah": 1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// Menampilkan semua data
	router.GET("/", func(c *gin.Context) {
		var (
			progdi  Progdi
			progdis []Progdi
		)

		rows, err := db.Query("select * from progdi;")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&progdi.Id, &progdi.Jenjang, &progdi.NmProgdi)
			progdis = append(progdis, progdi)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"Hasil":  progdis,
			"Jumlah": len(progdis),
		})

	})

	// Menambahkan Data
	router.POST("/", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.PostForm("id")
		jenjang := c.PostForm("jenjang")
		nmprodi := c.PostForm("nmprodi")
		stmt, err := db.Prepare("insert into progdi (id,jenjang,nmprodi) values(?,?,?);")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(id, jenjang, nmprodi)

		if err != nil {
			fmt.Println(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(jenjang) // menampilkan data yang berhasil ditambahkan
		buffer.WriteString(" ")
		buffer.WriteString(nmprodi) // menampilkan data yang berhasil ditambahkan
		defer stmt.Close()
		datane := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"Pesan": fmt.Sprintf("Data berhasil ditambahkan %s", datane),
		})
	})

	// PUT merubah data
	router.PUT("/", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.PostForm("id")
		jenjang := c.PostForm("jenjang")
		nmprodi := c.PostForm("nmprodi")
		stmt, err := db.Prepare("update progdi set jenjang=?, nmprodi=? where id=?;")
		if err != nil {
			fmt.Println(err.Error())
		}

		_, err = stmt.Exec(jenjang, nmprodi, id)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(jenjang)
		buffer.WriteString(" ")
		buffer.WriteString(nmprodi)
		defer stmt.Close()
		datane := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"Pesan": fmt.Sprintf("Berhasil merubah menjadi %s", datane),
		})

	})

	//Delete data
	router.DELETE("/", func(c *gin.Context) {
		id := c.PostForm("id")
		stmt, err := db.Prepare("delete from progdi where id=?;")
		if err != nil {
			fmt.Println(err.Error())
		}

		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"Pesan": fmt.Sprintf("Data berhasil di hapus %s", id),
		})
	})

	router.Run(":8080")
}
