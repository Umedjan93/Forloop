package main

import "fmt"

/*func sayHello(){
	fmt.Println("Hello World!")
	}*/
func main() {
	/*a := 10
	b := 110
	c := b/a
	for i:=0; i<=c; i++ {
		fmt.Println("Krug", i)
	}*/

	/*i := 10
	b := 110
	c := b-i
	for; i<=110; {
	fmt.Println("Distance left:", c)
	}*/

	click := 0
	distanceCount := 0
	sprintZone := 110
	for distanceCount <= sprintZone {
		click ++
		fmt.Println("click:", click)
		distanceCount +=10
		fmt.Println("DistanceCount:", distanceCount)
	}
}
