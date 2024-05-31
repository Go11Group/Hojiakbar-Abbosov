package main

import "fmt"

func generateMatrix(n int) [][]int {
 matrix := make([][]int, n)
 for i := range matrix {
  matrix[i] = make([]int, n)
 }

 directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
 directionIndex := 0 
 row, col := 0, 0    

 for num := 1; num <= n*n; num++ {
  matrix[row][col] = num
  nextRow, nextCol := row+directions[directionIndex][0], col+directions[directionIndex][1]

  if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || matrix[nextRow][nextCol] != 0 {
   directionIndex = (directionIndex + 1) % 4
   nextRow, nextCol = row+directions[directionIndex][0], col+directions[directionIndex][1]
  }

  row, col = nextRow, nextCol 
 }

 return matrix
}

func main() {
	
 fmt.Println("Input: n = 3")
 fmt.Println("Output:", generateMatrix(3))

 fmt.Println("\nInput: n = 1")
 fmt.Println("Output:", generateMatrix(1))
}