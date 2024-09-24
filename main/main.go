package main

import (
	"exam4"
	"fmt"
)

func main() {
	fmt.Println(exam4.Ft_coin([]int{1, 2, 5}, 11), 3) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(exam4.Ft_coin([]int{1, 2}, 11), 6)    // resultat : 6
	fmt.Println(exam4.Ft_coin([]int{2}, 3), -1)       // resultat : -1
	fmt.Println(exam4.Ft_coin([]int{1}, 0), 0)        // resultat : 0

	fmt.Println(exam4.Ft_missing([]int{3, 1, 2}), 0)                   // resultat : 0
	fmt.Println(exam4.Ft_missing([]int{0, 1}), 2)                      // resultat : 2
	fmt.Println(exam4.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8) // resultat : 8

	fmt.Println(exam4.Ft_profit([]int{7, 1, 5, 3, 6, 4}), 5) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(exam4.Ft_profit([]int{7, 6, 4, 3, 1}), 0) // resultat : 0

	fmt.Println(exam4.Ft_max_substring("abcabcbb"), 3) // resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(exam4.Ft_max_substring("bbbbb"), 1) // resultat : 1
	// "b" est la plus grande sous chaine

	fmt.Println(exam4.Ft_min_window("ADOBECODEBANC", "ABC"), "BANC") // resultat : "BANC"
	fmt.Println(exam4.Ft_min_window("a", "aa"), "")                  // resultat : ""

	//fmt.Println(exam4.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}), 1) // resultat : 1
	//// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	//// il suffit de d'enlever qu'un seul intervalle, [1,3]
	//fmt.Println(exam4.Ft_non_overlap([][]int{{1, 2}, {2, 3}}), 0)         // resultat : 0
	//fmt.Println(exam4.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}), 2) // resultat : 2

}
