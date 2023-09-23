package reverse

import "strconv"

// code taken from https://go.dev/doc/tutorial/workspaces
// return decimal reversal of int i
func Int(i int) (reversed_int int) {
	reversed_int, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return
}
