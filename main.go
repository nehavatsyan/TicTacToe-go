package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/tictactoe/:turn", getValue)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func getValue(c *gin.Context) {
	turn := c.Param("turn")
	var a = make([]string, 9)
	var x string
	x = checkWin(a)
	if x != "" {
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
	} else {
		win = "Nobody"
	}
	return win
}
