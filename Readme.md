Adakar

A tool used to search for any word lower case or upper case like "the, The or THE" over http connection and tell how many times that number
has repeated.
by usman79979

Usage:

-Type "git clone https://github.com/sunny79979/adakar.git" in your terminal
-"cd" to the cloned git repository
-"go build" from inside the repository folder
-"go run wordsearch.go"

Outputs:
There are two prompts:
1)
"Enter Url"
- enter the url begining with "http:// or https://"

2)
"Search Repetition of a word"
-enter the word you wish to search

You should get the number of times the word has repeated.

Testing:
There are three cases that the test program runs with the code.
If you have cloned the folder. Then the test file will be in folder as wordsearch_test.go
To run it type:
go test

Log file:
test.txt will be generated if not present already to log all the errors for the expected test case result.
