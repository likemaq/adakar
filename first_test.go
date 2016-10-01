package main

import (
        "testing"
        "strconv"
        "log"
        "os"
)

func TestURLS(t *testing.T){
        cases := []struct{ url string; search string; expected int}{
                {"https://google.com", "the", 1},
                {"https://golang.org", "the", 7},
                {"http://blank.org/", "the", 0},
        }

        for _, c := range cases {
        count := find(c.url, c.search)
        var originalFile *os.File
        if count != c.expected {
                _, err := os.Stat("test.txt")
                if err != nil {
                        originalFile, err = os.Create("test.txt")
                        if err != nil {
                                log.Fatal(err)
                        }
                }
                originalFile, err = os.OpenFile("test.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666,)
                if err != nil {
                        log.Fatal(err)
                }
                byt := []byte("count should be " + strconv.Itoa(c.expected) + " but got " + strconv.Itoa(count) + "\n")
                _,err = originalFile.Write(byt)
                if err!=nil{
                        log.Fatal(err)
                }
                t.Fail()
        }
        defer originalFile.Close()
        }

}
