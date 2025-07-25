package main

/*
write a function that takes a slice of ints and returns the average
 */

func Average(nums[]int)float64{
	sum:=0
	for i := 0; i <len(nums) ; i++ {

		sum+=nums[i]

	}
	return float64(sum)/float64(len(nums))
}
/*
with error handling and idiomatic go
*/

for AverageSafe(nums[]int)(float64,error){
if(len(nums)==0){
return 0,errors.New("cannot computer average of empty slice")
}
sum:=0
for _,v :=range nums {
sum+=v
}
return float64(sum)/float64(len(nums)),nil
}
func main() {

	nums:=[]int{2,3,4,5}
	avg:=Average(nums)
	fmt.Println(avg)


	nums2 := []int{}
	avg, err := AverageSafe(nums2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average is:", avg)
	}


}


