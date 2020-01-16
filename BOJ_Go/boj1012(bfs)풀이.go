package main

import "fmt"

func main() {

	var t int

	fmt.Scanln(&t)

	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}

	for cnt := 0; cnt < t; cnt++ {

		var m int
		var n int
		var k int

		fmt.Scanln(&m, &n, &k)

		field := make([][]int, n)
		chk := make([][]bool, n)

		for i := 0; i < n; i++ {

			field[i] = make([]int, m)
			chk[i] = make([]bool, m)
		}

		for i := 0; i < n; i++ {

			for j := 0; j < m; j++ {

				field[i][j] = 0
				chk[i][j] = false
			}
		}

		for i := 0; i < k; i++ {

			var x int
			var y int

			fmt.Scanln(&x, &y)

			field[y][x] = 1
		}

		result := 0

		for i := 0; i < n; i++ {

			for j := 0; j < m; j++ {

				if field[i][j] == 1 && !chk[i][j] {

					chk[i][j] = true
					result++

					var qx []int
					var qy []int

					qx = append(qx, j)
					qy = append(qy, i)

					for len(qx) > 0 && len(qy) > 0 {

						for c := 0; c < 4; c++ {

							if qx[0]+dx[c] >= 0 && qx[0]+dx[c] < m && qy[0]+dy[c] >= 0 && qy[0]+dy[c] < n {

								if field[qy[0]+dy[c]][qx[0]+dx[c]] == 1 && !chk[qy[0]+dy[c]][qx[0]+dx[c]] {

									chk[qy[0]+dy[c]][qx[0]+dx[c]] = true

									qx = append(qx, qx[0]+dx[c])
									qy = append(qy, qy[0]+dy[c])
								}
							}
						}

						qx = qx[1:]
						qy = qy[1:]
					}
				}
			}
		}

		fmt.Println(result)
	}
}
