package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)



func main(){    
slice1 := []string{"foo", "bar", "hello"}
slice2 := []string{"foo", "bar"}
fmt.Println(interSection(slice1, slice2))
// [foo bar]

ints1 := []int{1, 2, 3, 9, 8}
ints2 := []int{10, 4, 2, 4, 8, 9} // have duplicated values
ints3 := []int{2, 4, 8, 1}
fmt.Println(interSection(ints1, ints2, ints3))
//[2 8]
}

func interSection[T constraints.Ordered](in ...[]T) []T{
    hash := make(map[T]*int) // value, counter
    result := make([]T, 0)
    for _, slice := range in {
        duplicationHash := make(map[T]bool) // duplication checking for individual slice
        for _, value := range slice {
            if _, isDup := duplicationHash[value]; !isDup { // is not duplicated in slice
                if counter := hash[value]; counter != nil { // is found in hash counter map
                    if *counter++; *counter >= len(in) { // is found in every slice
                        result = append(result, value)
                    }
                } else { // not found in hash counter map
                    i := 1
                    hash[value] = &i
                }
                duplicationHash[value] = true
            }
        }
    }
    return result
}