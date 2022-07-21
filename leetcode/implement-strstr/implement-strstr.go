package implement_strstr

func StrStr(haystack, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)
	if lenHaystack == 0 || lenNeedle == 0 {
		return 0
	}
	if lenNeedle > lenHaystack {
		return -1
	}
	iFirstSymbol := -1
	iCurrSymbol := 0
	for i := 0; i < lenHaystack; i++ {
		if haystack[i] == needle[iCurrSymbol] {
			if iFirstSymbol == -1 {
				iFirstSymbol = i
			}
			if lenNeedle == iCurrSymbol+1 {
				return iFirstSymbol
			}
			iCurrSymbol++
		} else {
			if iFirstSymbol != -1 {
				i = iFirstSymbol
			}
			iFirstSymbol = -1
			iCurrSymbol = 0
		}
	}
	return -1
}
