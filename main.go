package main

import (
	"fmt"
	"math"
)
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func update(x string){
	fmt.Println(x)
	fmt.Println(&x)
}

type food struct{
	name string
	grid map[string]float64
	tip int
}


type Vertex struct {
	X, Y float64
}


func (f food) format() string{
	f.grid["new"] = 12
	f.tip = 12
	fs := "Item breakdown \n"
	total := 0.
	for i, v := range f.grid{
		fs += fmt.Sprintf(" %v ...$%v \n", i, v)
		total += v
	}
	fs += fmt.Sprintf("total: %v , tip: %v", total, f.tip)
	
	return fs
}

func compute(fn func(float64, float64) float64) float64{
	return fn(2.2,3.1)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}


func (v *Vertex) Abs() float64 {
	v.Y = 6
	fmt.Println(v.X, v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v *Vertex) float64{
	v.Y = 6
	fmt.Println(v.X, v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	WordFreqCounter("hello world, how yall don'n there buddy ?")

	// GradeCalculator()
}