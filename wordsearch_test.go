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
        count :=  find(c.url, c.search)
        var originalFile *os.File
        if count != c.expected {
                _, err := os.Stat("test.txt")
                if err != nil {
                        originalFile, err = os.Create("test.txt")
                        if err != nil {
                                log.Fatal(err)
                        }
                }
                originalFile, err = os.OpenFile("test.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
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

func TestUPERCASE(t *testing.T){
	cases := []struct{ word string; expected string }{
		{"the", "The"},
		{"t", "T"},
		{"THE", "THE"},
	}
	for _, c := range cases {
		checkword := upercasing(c.word)
		if checkword != c.expected {
			t.Errorf("Expected this %s but got this %s", c.expected, checkword)
			t.Fail()
		}
	}
}

func TestConnection(t *testing.T){
	cases := []struct{ url string; expected bool}{
		{"https://google.com",true},
		{"https://golang.org",true},
		{"https://hityy.com",false},	
	}
	for _, c := range cases {
		_, connected := connecturl(c.url)
		if connected != c.expected {
			t.Errorf("Expected this %b but got this %b", c.expected, connected)
			t.Fail()
		}
		
	}
}

func TestMerge(t *testing.T){
	cases := []struct{one string; two string; three string; expected string}{
		{"the","The","THE","\\sthe|The|THE\\s"},
	}
	for _, c := range cases{
		mergeword := mergestring(c.one,c.two,c.three)
		if mergeword != c.expected{
			t.Fail()
		}
	}
}
