package problems0412

func searchNext(haystack, needle string) int {
	retSlice := indexKmp(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[len(retSlice)-1]
	}

	return -1
}

func searchString(haystack, needle string) int {
	retSlice := indexKmp(haystack, needle)
	//fmt.Println(retSlice)
	if len(retSlice) > 0 {
		return retSlice[0]
	}

	return -1
}

func indexKmp(haystack, needle string) []int {
	next := prefixTable(needle)

	//fmt.Println(next)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		//fmt.Printf("i ==>> %v j ==>> %v n ==>> %v\n", i, j, n)
		for i > -1 && x[i] != y[j] {
			//fmt.Printf("i ==>> %v j ==>> %v n ==>> %v\n", i, j, n)
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

// prefixTable pattern prefix table
func prefixTable(x string) []int {
	var i, j int
	length := len(x) - 1
	next := make([]int, len(x)+1)
	i = 0
	j = -1

	next[0] = -1 // init first element

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = next[j]
		}

		i++
		j++

		if x[i] == x[j] {
			next[i] = next[j]
		} else {
			next[i] = j
		}
	}
	return next
}
