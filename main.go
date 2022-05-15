package main

import (
	"fmt"
    "strings"
)

func main() {
    tests := []string {
        "ABBA", // returns ""
        "AABBBBAA", // returns ""
        "AAABBDBBAADCD", // returns "ADD"
        "ABBAA", // returns "A"
        "ABCBA", // returns "C"
        "AABBCABBAACDCDBD", // returns ""
        "BBAAAABAAAAABBBCDBBBBBBBCDBBBBBBDCBBBBBAAAAAAAAABAAAAAAAAB", // returns ""
    }

    for _, test := range tests {
        // declare a storage array so we can return
        // the final filtered string
        var filtered []string

        // break the string into it's individual chars
        val := strings.Split(test, "")
    
        // send the string into the filter engine
        // if we have more than one char to compare
        if len(val) > 1 {
            filtered = filterString(val[0], val[1:], nil)
        } else {
            fmt.Println(test)
        }
    
        // finally join all of the individual chars together
        // these are the filtered chars that have passed the filter
        fmt.Println(test, "=>", strings.Join(filtered, ""))
    }
}

func filterString(compare string, values []string, filtered []string) []string {
    // initialize filtered array when nil
    if filtered == nil {
        filtered = []string{}
    }

    // if we have no values to filter return the filtered array
    if len(values) == 0 {
        return filtered
    }

    // compare the two elements, the current char vs the first
    // char in the array of values leftover - if we need to skip
    // these two values, filter them out
    if skipElements(compare, values[0]) {
        // if we have 1 or less values in the val array we are excluding that value
        if len(values) <= 1 {
            // check if the filtered array contains anything with the test sequence
            // if so we need to filter those values
            if (containsSequence(filtered)) {
                return filterString(filtered[0], filtered[1:], nil)
            }
            
            // if we've reached this point the values have been filtered
            return filtered
        } else {
            // set values to exclude the first char
            values = values[1:]
        }
    } else {
        // if we don't want to skip the current char we need
        // to append it to the filtered list of chars
        filtered = append(filtered, compare)
    }

    // continue filtering the rest of the string
    if len(values) > 1 {
        return filterString(values[0], values[1:], filtered)
    } else {
        // if any values are left in the values array
        // we should return those
        if len(filtered) == 0 {
            return values
        }

        // at this point there's only one value left in the vals array
        // determine if we need to skip the last char in the filtered array
        // compare vs the first value
        if skipElements(filtered[len(filtered) - 1], values[0]) {
            // remove the last value from the filtered array
            filtered = filtered[:len(filtered) - 1]
        } else {
            // append the first value to the filtered array
            filtered = append(filtered, values[0])
        }
    }

    // finally check if we need to filter any more values
    // from the filtered array - this will check the collisions
    // as in the case of "AABBBBAA" > "ABBA" > ""
	if (containsSequence(filtered)) {
		return filterString(filtered[0], filtered[1:], nil)
	}
    
    // no more conditions remain, return the filtered array
    return filtered
}

// function that determines whether or not we should
// skip the current elements as per the test cases
func skipElements(first string, second string) bool {
    skip := false

    if (first == "A" && second == "B") || (first == "B" && second == "A") {
        skip = true
    }

    if (first == "C" && second == "D") || (first == "D" && second == "C") {
        skip = true
    }

    return skip
}

// check if the compiled string has any
// sequence "AB" / "BA" / "CD" / "DC"
func containsSequence(compared []string) bool {
    final := strings.Join(compared, "")

    if strings.Contains(final, "AB") || strings.Contains(final, "BA") || strings.Contains(final, "CD") || strings.Contains(final, "DC") {
        return true
    }

    return false
}