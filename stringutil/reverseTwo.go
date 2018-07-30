package stringutil

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
//func reverseTwo(s string) string {
//	r:=[]rune(s)
//	j:=len(r)-1
//	for i:= 0 ; i < j ; i+1j-1 {
//		temp = r[i]
//		r[i] = r[j]
//		r[j] = temp
//	}
//	return string(r)
//}

// this demonstrates how an unexported function
// can be used by an exported function in the same package
