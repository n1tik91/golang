// Go program to illustrate the use of keywords 
package main 

import "fmt"

// Here, package, import, func & var are keywords 
func main() { 

// Here, a is a valid identifier 
var a = "GeeksforGeeks"

fmt.Println(a) 

// Here, the default is an illegal identifier and compiler will throw an error var default = "GFG" 


// Here, var keyword is used  to create variables Pname, Lname, and Cname are identifiers 
    var Sname = "GeeksforGeeks" 
    var Lname = "Go Language" 
    var Cname = "Keywords"
      
    fmt.Printf("Site name: %s", Sname) 
    fmt.Printf("\nLanguage name: %s", Lname) 
    fmt.Printf("\nChapter name: %s", Cname) 
} 
