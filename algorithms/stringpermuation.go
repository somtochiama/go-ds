package algorithms

import (
	"fmt"
)

// func powerSet(str string) []string{
// 	ret := []string{""}

// 	for _, char := range str {
// 		copyRet := make([]string, len(ret))
// 		fmt.Println(len(copyRet), len(ret))
// 		copy(copyRet, ret)
// 		for i, str := range copyRet {
// 			copyRet[i] = str + string(char)
// 			fmt.Println(copyRet)
// 		}

// 		ret = append(ret, copyRet...)
// 	}

// 	return ret
// }

func StringPermuation(str string)  {
	permute("", str)
}

func permute(prefix, str string) {
	if len(str) == 0 {
		fmt.Println(prefix)
	}

	for i, _ := range str {
		permute(prefix + string(str[i]), str[:i] + str[i+1:])
	}
}