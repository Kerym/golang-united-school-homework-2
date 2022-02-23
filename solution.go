package square

import(
	"math"
	// "fmt"
) 
type SidesN int

const (
	SidesTriangle = 3
	SidesSquare = 4
	SidesCircle = 0
)

func CalcSquare(sideLen float64, sidesNum SidesN) float64 {
	if sidesNum == SidesTriangle {
		return (3 * sideLen)
	} else if sidesNum == SidesSquare {
		return (2 * sideLen)
	} else if sidesNum == SidesCircle {
		return (math.Pi * 2 *sideLen)
	} else {
		return 0
	}
}

// func main(){
// 	fmt.Println(CalcSquare(2, SidesTriangle))
// 	fmt.Println(CalcSquare(2, SidesSquare))
// 	fmt.Println(CalcSquare(2, SidesCircle))
// 	fmt.Println(CalcSquare(2, 5))
// }

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)