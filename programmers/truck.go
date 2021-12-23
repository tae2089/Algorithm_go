package programmers

func RunningTruck(bridegeLength, weight int, truckWeights []int) int {
	answer := 0
	trucksOnBridge := make([]int, bridegeLength)
	for len(trucksOnBridge) > 0 {
		answer += 1
		trucksOnBridge = trucksOnBridge[1:]
		if len(truckWeights) > 0 {
			if sum(trucksOnBridge)+truckWeights[0] <= weight {
				truck := truckWeights[0]
				truckWeights = truckWeights[1:]
				trucksOnBridge = append(trucksOnBridge, truck)
			} else {
				trucksOnBridge = append(trucksOnBridge, 0)
			}
		}
	}
	return answer
}

func sum(arr []int) int {
	total := 0
	for _, data := range arr {
		total += data
	}
	return total
}
