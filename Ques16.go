package main
/*
Write a function divide(a, b int) (int, error) that returns error if b == 0


Call it from main and handle the error gracefully

 */

//function with error return
func divide(a,b int)(int,error){
	if b==0{
		return 0,fmt.Errorf("divide by zero is not allowed")
	}
	return a/b,nil
}

func main() {
	res,err :=divide(10,2)
	if err!=nil{
		fmt.Println("Error",err)
		
		
		
		
	}else{
		fmt.Println("res",res)
	}

	//fallen case try
	res,err = divide(10,0)
	if err!=nil{
		fmt.Println("Error",err)
	}else {
		fmt.Println("res",res)
	}
}