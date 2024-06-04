package main

func main(){


	// sum := 0
	// sum2 := 0


	// for i := 0;i<=100;i++{
	// 	if i % 2 == 0{
	// 		sum = sum + i
	// 	}else{
	// 		sum2 = sum2 + i
	// 	}
	// }

	// println(sum - sum2)

	n := 17
	check := false

	for i := 2;i < n;i++{
		if n % i == 0{
			check=true
		}
	}

	if check==true{
		println(false)
	}else{
		println(true)
	}
	




}