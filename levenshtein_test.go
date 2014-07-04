// Copyright 2014 Sean Thorne.  All rights reserved.
// Use of this code is governed by the license found in LICENSE
package levenshtein

import "testing"

func TestCalc(t *testing.T) {

    One := "brndm"
    Two := "brand"

    d := Calc(One, Two)

    if d != 1 {
        t.Fail("Wrong Distance")
    }

    t.Log(One, Two, d)

    One = "brndms"
    Two = "brands"

    d = Calc(One, Two)

    if d != 1 {
        t.Fail("Wrong Distance")
    }

    t.Log(One, Two, d)

    One = "brndmst"
    Two = "brimstone"

    d = Calc(One, Two)

    if d != 4 {
        t.Fail("Wrong Distance")
    }

    t.Log(One, Two, d)
}