func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var results []int32
	var result int32 = 0

	sort.SliceStable(candles, func(i, j int) bool {
		return candles[i] < candles[j]
	})

	for i := len(candles) - 1; i >= 0; i-- {
		if candles[i] >= int32(result) {
			result = candles[i]
			results = append(results, candles[i])
		}
	}

	akhir := len(results)
	// fmt.Println(akhir)
	return int32(akhir)

}
