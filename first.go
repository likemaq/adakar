package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	//"bytes"
	"strings"
	"log"
)

func main(){

	var search, in string
	in,search = getin(in, search)
	find(in,search)
}

func getin(url, se string) (string,string){
	fmt.Println("Enter URL: ")
	fmt.Scanln(&url)
	fmt.Println("Search Repetition of a Word: ")
	fmt.Scanln(&se)
	return url,se
}

func find(inp, sear string ) {
	var mer string
	//fmt.Println(len(string(search)))
	up := strings.ToUpper(string(sear[0]))
	mer=up;
	n := 1
	for n<len(sear){
		mer+=string(sear[n])
	//	fmt.Println(string(sear[n]))
		n++
	}
	//fmt.Println(mer)
	resp, err := http.Get(inp)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	txt := string(bytes)
	cas := strings.ToUpper(sear)
	//fmt.Println(cas)
	count:=0;
	r := regexp.MustCompile("\\s"+sear+"|"+cas+"|"+mer+"\\s")
	s := r.FindAllString(txt, -1)
	//fmt.Println(s[5])
	for i := 0;i< len(s);i++{
		fmt.Println(s[i])
		count++
	}
	fmt.Println(count)

}