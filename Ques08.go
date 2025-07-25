package main

/*
Create and populate a map[string]int to store scores of 3 players, print keys and values
 */




func main() {

	var scores map[string]int  //declare
	scores = make(map[string]int) //intialise

	//populate
	scores["Nav"]=23
	scores["Kri"]=24
	scores["Kav"]=50

	fmt.Println(scores["Nav"])
	fmt.Println(scores["Kri"])
	fmt.Println(scores["Kav"])


	//second method by direct intialization
	scores :=map[string]int{
		"K":12,
		"M":4,
		"f":23,

	}

	fmt.Println(scores)

//idiomatic way by key and values
	scores := map[string]int{
		"A": 90,
		"B": 82,
		"C": 95,
	}

	//  keys and values
	for player, score := range scores {
		fmt.Printf("%s: %d\n", player, score)
	}


}

