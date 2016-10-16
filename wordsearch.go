package main
/*
This program is used to find the number of times the word has appeared on a word page.
It has a config file where you can edit the program properties. its named as adakaar.json.
The major functions used in this program are 
-Regexp (link - https://golang.org/pkg/regexp/)
-net/http (link - https://golang.org/pkg/net/http/ )
-json (link - https://golang.org/pkg/encoding/json/)
All functions except menu and exitprog have a return value.
To test individual functions a test file is given which can be executed using go test
To test the performance of individual function use go test -run=uper1-7 -bench=.
*/
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"log"
	"os"
	"os/exec"
	"encoding/json"
)//import packages

//For json property maping
type Configuration struct {
	Defaulturl	string `json:"defaulturl"`
	Inputurlmsg	string `json:"url_input_message"`
	Word		string `json:"word_for_repetition"`
	Eorword		string `json:"error_word"`
	Exit		string `json:"exit"`
	Invalidexit	string `json:"invalid_input_for_exit"`
	Invalidurl	string `json:"invalid_url"`
}

//initialization of properties
func init(){
	config := Configuration{}
	config.Defaulturl="https://blank.org"
}

//Globle variable
var config Configuration

func main(){	
	file,_ := ioutil.ReadFile("./adakaar.json") //converting json to go
	err := json.Unmarshal(file, &config)
	if err!=nil{panic(err)}
	fmt.Println("Runing default diagnostics")
	_,er := exec.Command("go","test").Output() //testing the functions before program runs
	if er!=nil{
		fmt.Println("diagnostic failed")
		panic(er)
	}
	fmt.Println("Passed")
	menu()
	exitprog()
}

//Menu asks the user for URL and WORD
func menu() {
	var url, searchword string
	fmt.Println(config.Inputurlmsg)
	fmt.Scanln(&url)
	fmt.Println(config.Word)
	fmt.Scanln(&searchword)
	if len(searchword) == 0{
		fmt.Println(config.Eorword)
		menu()
	}
	count := find(url, searchword)
	print(count)
}

//Asks if to exit the program or not
func exitprog(){
	var exit string
	fmt.Println(config.Exit)
	fmt.Scanln(&exit)
	switch exit{
	case "N":
		menu()
		exitprog()
	case "Y":
		os.Exit(0)
	default:
		fmt.Println(config.Invalidexit)
		exitprog()
	}
}

//Finds the word from the web page entered by the user
func find(link, search string)(int){
	bodytxt,connected := connecturl(link)
	if connected != true{
		fmt.Println(config.Invalidurl)
		exitprog()
	}
	uperfirst := upercasing(search)
	capitalword := strings.ToUpper(search)
	r := regexp.MustCompile(mergestring(search,capitalword,uperfirst))
	s := r.FindAllString(bodytxt, -1)
	count := 0
	for i := 0;i< len(s);i++{
		count++
	}
	return count
}

//converts the first letter of the word to upercase like the -> The
func upercasing(word string)(string){
	var mer string
	mer = strings.ToUpper(string(word[0]))
	n := 1
	for n<len(word){
		mer+=string(word[n])
		n++
	}
	return mer
}

//Tell the program how words are <space>word<space>
func mergestring(one, two, three string)(string){
	return "\\s"+one+"|"+two+"|"+three+"\\s"
}

//connecting to the web url for http connection
func connecturl(website string) (string, bool){
	resp, err := http.Get(website)
	connected := true
	if err != nil {
		connected = false
		return "Failed", connected
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return string(bytes), connected
}

//prints the total count
func print(num int){
	fmt.Println("number of times word repeted: ", num)
}
