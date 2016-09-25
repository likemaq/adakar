package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
)

func main(){
	resp, _ := http.Get("https://blog.golang.org/playground")
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	txt := string(bytes)
	count:=0;
	r := regexp.MustCompile("\\sthe|The\\s")
	s := r.FindAllString(txt, -1)
	//fmt.Println(s[5])
	for i := 0;i< len(s);i++{
		fmt.Println(s[i])
		count++
	}
	fmt.Println(count)

}