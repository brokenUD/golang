package main

import (
    "fmt"
	// "math"
)

type dataStruct struct{
    i int;
    j int;
}
func main1() {
    var a, b, c int
    // n, _ := fmt.Scan(&a, &b)
    // if n == 0 {
    //     return
    // }
	// var a int64
	// a = math.MaxInt64
	a = 3
	b = 5
    datamap := make(map[int][]dataStruct, 0)
    var data [100][100]int
	// var data2 = new([][]int, a, b)
	// data := make([][]int, a, b)
    for i:=0;i<a;i++{
		// s := make([]int, b)
        for j:=0;j<b;j++ {
            n, _ := fmt.Scan(&c)
            if n == 0 {
                return
            }
            if _,ok:=datamap[c];ok{
                datamap[c] = append(datamap[c], dataStruct{
                    i: i,
                    j: j,
                })
            } else {
                datamap[c] = make([]dataStruct, 0)
				datamap[c] = append(datamap[c], dataStruct{
                    i: i,
                    j: j,
                })
            }
			// s = append(s, 0)
        }
		// data = append(data, s)
    }
	// data1 := data[:a][:b]
	fmt.Println(data[:a][:b])
	// var tmp int
	// var xf, yf float64
	// for _, dataV := range datamap {
	// 	if len(dataV) == 1 {
	// 		data[dataV[0].i][dataV[0].j] = -1
	// 	} else {
	// 		for m:= 0; m<len(dataV); m++ {
	// 			for n:= m+1; n<len(dataV); n++ {
	// 				xf = math.Abs(float64(dataV[m].i) - float64(dataV[n].i))
	// 				yf = math.Abs(float64(dataV[m].j) - float64(dataV[n].j))
	// 				tmp = int(xf)+int(yf)
	// 				if tmp<data[dataV[m].i][dataV[m].j] || data[dataV[m].i][dataV[m].j] == 0{
	// 					data[dataV[m].i][dataV[m].j] = tmp
	// 				}
	// 				if tmp<data[dataV[n].i][dataV[n].j] || data[dataV[n].i][dataV[n].j] == 0{
	// 					data[dataV[n].i][dataV[n].j] = tmp
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
