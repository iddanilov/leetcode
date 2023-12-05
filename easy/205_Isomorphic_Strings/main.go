package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

//It is my solution
func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	checkerS := make(map[byte]byte, 0)
	checkerT := make(map[byte]byte, 0)

	for i := 0; i < len(s); i++ {
		switch checkerS[s[i]] {
		case 0:
			if checkerT[t[i]] != 0 {
				return false
			}
			checkerS[s[i]] = t[i]
			checkerT[t[i]] = s[i]
		case t[i]:
			continue
		default:
			return false
		}

	}

	return true
}

// It is https://leetcode.com/klakovskiy/ solution, whitch i am like

// func isIsomorphic(s string, t string) bool {
// 	m := make(map[uint8]uint8)
// 	m2 := make(map[uint8]struct{})
// 	if len(s) != len(t) {
// 	  return false
// 	}
// 	for i := 0; i < len(s); i++ {
// 	  x, ok := m[s[i]]
// 	  if !ok {
// 		if _, ok := m2[t[i]]; ok {
// 		  return false
// 		}
// 		m[s[i]] = t[i]
// 		m2[t[i]] = struct{}{}
// 	  } else {
// 		if x != t[i] {
// 		  return false
// 		}
// 	  }
// 	}
// 	return true
//   }
