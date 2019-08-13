# goimportstidy

This tool updates your Go import lines, grouping it into three groups: 
 - stdlib,
 - external libraries,
 - local libraries (optional).
 
Installation: 

     $ go get github.com/krzysztofdrys/goimportstidy
     
Usage:

    $ goimportstidy -w -local github.com/shipwallet main.go
