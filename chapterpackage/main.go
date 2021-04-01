package main
import "fmt"
import "learngo/chapterpackage/math"

func main(){
	xs:= []float64{1,2,3,4}
	avg:= math.Average(xs)
	fmt.Println(avg)
}