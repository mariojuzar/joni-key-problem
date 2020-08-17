package main

import "fmt"

/**
Joni menititipkan kunci rumahnya pada seorang temannya bernama Andi.
Namun, Andi memberikan sebuah games kepada Joni untuk mencari kunci rumahnya.

Joni ditandai dengan titik X.
Untuk menemukan kunci rumahnya, Andi memberikan beberapa petunjuk kepada Joni yaitu:
1. Jalan ke utara sebanyak Y langkah.
2. Lalu jalan ke timur sebanyak Y langkah.
3. Terkahir ke selatan sebanyak Y langkah.
Variable ‘Y’ menandakan angka yang hilang.
Buktikan berapa banyak titik yang menjadi kemungkinan lokasi kunci rumah Joni.
Jika diketahui angka yang hilang merupakan bilangan bulat positif.

====
Bilangan bulat positif = 1,2,3,4,5, . . .
0 tidak termasuk bilangan bulat positif
====
 */

func main()  {
	// area from problem
	var area = [4][6]string {
		{".", ".", ".", ".", ".", "."},
		{".", "#", "#", "#", ".", "."},
		{".", ".", ".", "#", ".", "#"},
		{"X", "#", ".", ".", ".", "."},
	}

	count := 0

	for a:= 2; a >= 0; a-- { // jalan ke utara
		if area[a][0] == "." {
			b := 1
			for { // jalan ke timur
				if b > 5 {
					break
				}
				if area[a][b] == "#" {
					break
				}
				if area[a][b] == "." {
					c := 1
					for { // jalan ke selatan
						if a + c > 3 {
							break
						}
						if area[a + c][b] == "#" {
							break
						}
						if area[a + c][b] == "." {
							count++
						}
						c++
					}
				}
				b++
			}
		}
	}

	fmt.Println("Banyak titik yang menjadi kemungkinan lokasi yaitu: ", count)
}