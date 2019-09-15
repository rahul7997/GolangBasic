package main
import ("fmt")

func add(x int,y int) int{
	out := x+y
	return out
}

func multi(x,y int) int{
	var out = x*y;
	return out
}

func calc(x,y int) (int, int, int){
	var out1 = x+y
	var out2 = x-y;
	var out3 = x*y
	return out1, out2, out3
} 

func calc2(x,y int) (out1, out2 int){
	out1 = x+y
	out2 = x-y
	return
}

func main(){
    num1 := 5
	num2 := 4
	
    sum := add(num1, num2)
	fmt.Println(sum)

	var product = multi(num1, num2)
	fmt.Println(product)

	var result1, result2, result3 = calc(num1, num2)
	fmt.Println(result1, result2, result3)

	fmt.Println(calc2(num1,num2))
}