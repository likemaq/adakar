package main
/*

*/
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"log"
	"os"
)

func main(){
	menu()
	exitprog();
}

func menu() {
	var url, searchword string
	fmt.Println("Enter URL:(http://***.** or https://***.**) ")
	fmt.Scanln(&url)
	fmt.Println("Enter the word for Repetition Count: ")
	fmt.Scanln(&searchword)
	if len(searchword) == 0{
		fmt.Println("No word entered. Enter Again.")
		menu()
	}
	count := find(url, searchword)
	print(count)
}

func exitprog(){
	var exit string
	fmt.Println("Exit? Y | N : ")
	fmt.Scanln(&exit)
	switch exit{
	case "N":
		menu()
		exitprog()
	case "Y":
		os.Exit(0)
	default:
		fmt.Println("Invalid Input Enter again")
		exitprog()
	}
}

func find(link, search string)(int){
	bodytxt,connected := connecturl(link)
	if connected != true{
		fmt.Println("Invalid URL")
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

func mergestring(one, two, three string)(string){
	return "\\s"+one+"|"+two+"|"+three+"\\s"
}

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

func print(number int){
	fmt.Println(number)
}
