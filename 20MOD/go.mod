module github.com/anantacoder/mod

go 1.24.2

require github.com/gorilla/mux v1.8.1

// go env => go path , cd GOpath
// go mod tidy -- remove indirect comment and make them direct
// go  mod verify -- verify the dependencies
//go mod graph --- show all dependancies in graph format

//PS D:\GO\20MOD> go list -m all
// github.com/anantacoder/mod
// github.com/gorilla/mux v1.8.1

//PS D:\GO\20MOD> go list -m -versions github.com/gorilla/mux -> all gorilla mux versiions

// github.com/gorilla/mux v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.7.0 v1.7.1 v1.7.2 v1.7.3 v1.7.4 v1.8.0 v1.8.1


//mo mod why [module name] -- why this module is needed 

// go mod vender -- make the modules in the vender folder
// go mod -mod=vendor main.go --use modules in vender folder 