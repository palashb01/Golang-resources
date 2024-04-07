package main

import "fmt"

func main() {
	fmt.Println("loops in go")
	days := []string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}
	fmt.Println(days)
	// for i:=0; i<len(days);i++{
	// 	fmt.Println(days[i])
	// }
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }
	// for _,day:=range days{
	// 	fmt.Println(day)
	// }
	rogueValue:=1
	for rogueValue<10{
		if rogueValue==2{
			goto lco
		}
		if rogueValue==5{
			rogueValue++
			continue
		}
		fmt.Println("Value is: ",rogueValue)
		rogueValue++
	}
	lco:
		fmt.Println("Jumping to LCO")

}