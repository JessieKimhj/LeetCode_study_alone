package main

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// sort them first, the most different one will be in first and last
	sort.Strings(strs)

	// compare first and last
	l := len(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[l-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}

// func longestCommonPrefix(strs []string) string {
//     if len(strs) == 0 {
//         return ""
//     }
//     p := strs[0] // 기준 문자열
//     for i := range(len(p)){
//         for _, str := range strs {
//             if i >= len(str) || str[i] != p[i] {
//                 return p[:i] // 공통 접두사 여기까지
//             }
//         }
//     }
//     return p // 끝까지 공통 접두사면 첫 문자열 전체 반환
// }
