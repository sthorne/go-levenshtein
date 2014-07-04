// Copyright 2014 Sean Thorne.  All rights reserved.
// Use of this code is governed by the license found in LICENSE
package levenshtein

func Calc(One string, Two string) int {

    l1 := len(One)
    l2 := len(Two)

    if l1 == 0 {
        return l1
    }

    if l2 == 0 {
        return l2
    }

    c0, c1, c2 := 0, 0, 0

    p1 := make([]int, l2 + 1)
    p2 := make([]int, l2 + 1)

    for i := 0; i <= l2; i++ {
        p1[i] = i
    }

    for n := 0; n < l1; n++ {
        p2[0] = p1[0]

        for x := 0; x < l2; x++ {

            c0 = p1[x]
            c1 = p1[x + 1]
            c2 = p2[x] + 1

            if One[n] != Two[x] {
                c0 += 1
            }

            if c1 < c0 {
                c0 = c1
            }

            if c2 < c0 {
                c0 = c2
            }

            p2[x + 1] = c0
        }

        t := p1
        p1 = p2
        p2 = t
    }

    return p1[l2]
}