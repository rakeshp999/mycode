/* Alta3 Research | RZFeeser

Map several programming languages to their file extension, 
for example, the key Python should map to the value .py. 
The key Golang should map to .go, the key Java should map to .java, 
the key Ansible should map to .yml, and the key C++ should map to .cpp. 
Once the map is constructed, display it on the screen. Then remove the language C++, 
and add Julia and the key .jl. Display the new map after the change.

   CHALLENGE 01 - Mapping programming languages to extension names */


package main

import "fmt"

func main() {
    var fileExtensions = map[string]string{
        "Python": ".py",
        "C++":    ".cpp",
        "Java":   ".java",
        "Golang": ".go",
        "Ansible": ".yml",
    }

    fmt.Println(fileExtensions)

    // remove the key C++ and its associated value
    delete(fileExtensions, "C++")

    // because the map is initalized, we can add new keys to it
    fileExtensions["Julia"] = ".jl"

    // display the state of the map
    fmt.Println(fileExtensions)
}

