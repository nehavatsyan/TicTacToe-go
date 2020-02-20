package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/tictactoe/:turn/:arf/:brf/:crf/:ars/:brs/:crs/:art/:brt/:crt", getValue)
	r.Run() // listen and serve on 0.0.0.0:8080
}
func getValue(c *gin.Context) {
	turn := c.Param("turn")
	var a = make([]string, 9)
	a[0] = c.Param("arf")
	a[1] = c.Param("brf")
	a[2] = c.Param("crf")
	a[3] = c.Param("ars")
	a[4] = c.Param("brs")
	a[5] = c.Param("crs")
	a[6] = c.Param("art")
	a[7] = c.Param("brt")
	a[8] = c.Param("crt")
	var x string
	x = checkWin(a)
	//uri = "localhost:8080/" + turn + "/" + a[0] + "/" + a[1] + "/" + a[2] + "/" + a[3] + "/" + a[4] + "/" + a[5] + "/" + a[6] + "/" + a[7] + "/" + a[8]
	if x != "_" {
		fmt.Println(x)
		c.HTML(http.StatusOK, "win.html", gin.H{
			"winner": x,
		})
	} else {

		c.HTML(http.StatusOK, "main.html", gin.H{
			"val": turn,
			"a1":  a[0],
			"a2":  a[1],
			"a3":  a[2],
			"b1":  a[3],
			"b2":  a[4],
			"b3":  a[5],
			"c1":  a[6],
			"c2":  a[7],
			"c3":  a[8],
			//"uri": uri,
		})
	}
}
func checkWin(a []string) string {
	var win string
	if a[0] == a[1] && a[0] == a[2] {
		win = a[0]
	} else if a[3] == a[4] && a[3] == a[5] {
		win = a[3]
	} else if a[6] == a[7] && a[6] == a[8] {
		win = a[6]
	} else if a[0] == a[3] && a[0] == a[6] {
		win = a[0]
	} else if a[1] == a[4] && a[1] == a[7] {
		win = a[1]
	} else if a[2] == a[5] && a[2] == a[8] {
		win = a[2]
	} else if a[0] == a[4] && a[0] == a[8] {
		win = a[0]
	} else if a[2] == a[4] && a[2] == a[6] {
		win = a[2]
	} else if a[0] != "_" && a[1] != "_" && a[2] != "_" && a[3] != "_" && a[4] != "_" && a[5] != "_" && a[6] != "_" && a[7] != "_" && a[8] != "_" {
		win = "Nobody :-("
	} else {
		win = "_"
	}
	return win
}
