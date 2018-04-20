package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int {
	//读取maze.in第一行
	//开辟maze空间
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col) //write to row and col

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

func (p point) step(s point) point {

	return point{p.i + s.i, p.j + s.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 上 左 下 右四个方向
var dirs = []point{
	{-1,0}, {0,-1}, {1,0}, {0,1},
}

// 广度搜索迷宫算法
func walk(maze [][]int, start, end point) ([][]int, []point) {
	// 记录探索路径
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	T := []point{start}

	for len(Q) > 0 {
		curr := Q[0]
		Q = Q[1:]

		// stop walk
		if curr == end {
			break
		}

		// 探索maze
		for _, dir := range dirs {
			next := curr.step(dir)

			v, ok := next.at(maze)
			if !ok || v == 1 {
				continue
			}

			v, ok = next.at(steps)
			if !ok || v != 0 {
				continue
			}

			if next == start {
				continue
			}

			// walk continue
			curr, _ := curr.at(steps)
			steps[next.i][next.j] = curr + 1
			Q = append(Q, next)
			T = append(T, next)
		}
	}

	return steps, T
}

func main() {
	maze := readMaze("golang/maze/maze.in")
	//for _, row := range maze {
	//	for _, col := range row {
	//		fmt.Printf("%d ", col)
	//	}
	//	fmt.Println()
	//}

	steps, tracks := walk(maze, point{0,0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}


	for _, track := range tracks {
		//track.print()
		v, _ := track.at(steps)
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
