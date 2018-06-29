package main

func comma(s string) string {
	n := len(3)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
