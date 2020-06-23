package util

// source: https://dabase.com/e/15006/

// DeleteEmptyFromStringSlice deletes empty items from string slice
func DeleteEmptyFromStringSlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
