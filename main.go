package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	// http.HandleFunc("/find-path", findpath)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }

func main() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// api.GET("/", homepage)
	// api.POST("/findpath/", findpath)

	router.Run(":3000")
}

func homepage(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Homepage not implemented yet",
	})
}

func findpath(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Homepage not implemented yet",
	})
}

// func findpath() {
// 	var matrix [20][20]int
// 	var startX, startY, endX, endY = 0, 0, 4, 4

// 	if dfs(matrix, startX, startY, endX, endY) {
// 		fmt.Println("Path found from", startX, startY, "to", endX, endY)
// 	} else {
// 		fmt.Println("No path found from", startX, startY, "to", endX, endY)
// 	}
// }

// func dfs(matrix [][]int, startX, startY, endX, endY int) bool {
// 	m, n := len(matrix), len(matrix[0])
// 	visited := make([][]bool, m)
// 	for i := range visited {
// 		visited[i] = make([]bool, n)
// 	}

// 	var dfsFunc func(int, int)
// 	dfsFunc = func(x, y int) {
// 		if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
// 			return
// 		}
// 		visited[x][y] = true

// 		if x == endX && y == endY {
// 			return true
// 		}

// 		for _, dx := range []int{-1, 0, 1} {
// 			for _, dy := range []int{-1, 0, 1} {
// 				if dx == 0 && dy == 0 {
// 					continue
// 				}
// 				nx, ny := x+dx, y+dy
// 				if nx >= 0 && nx < m && ny >= 0 && ny < n {
// 					dfsFunc(nx, ny)
// 				}
// 			}
// 		}
// 		return false
// 	}

// 	dfsFunc(startX, startY)
// 	return false // or true, depending on whether you want to return a boolean indicating success
// }
