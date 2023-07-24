package backjun

// var (
// 	sc = bufio.NewScanner(os.Stdin)
// 	w  = bufio.NewWriter(os.Stdout)
// )

func func1475(data string) int {
	// defer w.Flush()
	// sc.Scan()
	// data := sc.Text()
	table := make(map[rune]int)
	for _, v := range data {
		table[v]++
	}
	max := 0
	table[54] = (table[54] + table[57] + 1) / 2
	table[57] = 0
	for _, v := range table {
		if max < v {
			max = v
		}
	}
	return max
	// w.WriteString(strconv.Itoa(max))
}
