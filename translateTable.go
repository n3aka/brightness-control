package main

func getAbsB(br int) []byte {

	i := br + cor
	i = i + 1000
	if i < 700 {
		return []byte("5")
	}
	if i > 9400 {
		return []byte("1500")
	}
	switch {
	case i < 700:
		return []byte("5")
	case i < 1500:
		return []byte("135")
	case i < 2300:
		return []byte("265")
	case i < 3100:
		return []byte("395")
	case i < 3900:
		return []byte("525")
	case i < 4700:
		return []byte("655")
	case i < 5500:
		return []byte("785")
	case i < 6300:
		return []byte("915")
	case i < 7100:
		return []byte("1045")
	case i < 7900:
		return []byte("1175")
	case i < 8700:
		return []byte("1305")
	case i < 9500:
		return []byte("1435")
	}
	return []byte("0")
}
