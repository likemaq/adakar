package main

import (
        "testing"
        "strconv"
)

func TestURLS(t *testing.T){
        cases := []struct{ url string; search string; expected int}{
                {"https://google.com", "the", 1},
                {"https://golang.org", "the", 7},
                {"http://blank.org/", "the", 0},
        }

        for _, c := range cases {
        count := find(c.url, c.search)
        if count != c.expected {
                t.Log("count should be " + strconv.Itoa(c.expected) + " but got$
                t.Fail()
        }
        }

}
