package util

// Substr - UTF-8 aware substring function
// stackoverflow.com/questions/12311033/extracting-substrings-in-go/ 
func Substr(input string, start int, length int) string {
    asRunes := []rune(input)
    
    if start >= len(asRunes) {
        return ""
    }
    
    if start+length > len(asRunes) {
        length = len(asRunes) - start
    }
    
    return string(asRunes[start : start+length])
}

// https://www.dotnetperls.com/duplicates-go
func RemoveDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func CountDuplicates(dupArr []int) int {
	dupsize := len(dupArr)
	dupcount := 0
	for i := 0; i < dupsize; i++ {
		for j := i + 1; j < dupsize; j++ {
			if dupArr[i] == dupArr[j] {
				dupcount++
				break
			}
		}
	}
	return dupcount
}

// subset returns true if the first array is completely
// contained in the second array.
func Subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

