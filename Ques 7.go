package main

/*
Write a function Max(nums ...int) (int, error) that:
accepts any number of arguments


returns an error if no arguments are provided

 */


func Max(nums ...int)int{
	max:=nums[0]
	for _,v:=range nums{
		if v>max{
			max=v
		}
	}
	return max
}

//with error handling
func MaxSafe(nums ...int)(int,error){
	if len(nums)==0{
		return 0,errors.New("No numbers provided")
	}
	max :=nums[0]
	for _,v:=range nums[1:]{
		if v>max {
			max =v
		}
	}
	return max,nil

}

func main(){

	result,err :=Max(3,4,5,6,8)
	if err!=nil {
		fmt.Println("ERROR",err)
	}else {
		fmt.Println("Max is",result)
	}


	//  with empty input
	result, err = Max()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Max is:", result)
	}
}