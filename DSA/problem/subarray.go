/*
	tags: subarray, brute force
/*



func subArrayBruteForce(list []int) int {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			fmt.Println(list[i : j+1])
		}
	}
}
