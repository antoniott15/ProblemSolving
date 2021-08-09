The technical assessment description/instructions:

 

Write a program that prints all the numbers from 1 to 100. However, for

multiples of 3, instead of the number, print “Linio”. For multiples of 5 print

“IT”. For numbers which are multiples of both 3 and 5, print “Linianos”.

 

But here’s the catch: you can use only one if. No multiple branches, ternary

operators or else.

Requirements

 

* 1 if

* You can’t use else, else if or ternary

* Unit tests

* Feel free to apply your SOLID knowledge

* Go

* Production ready code


Run main

```go
 go run linio.go
```

Run test

```go
 go test
```

Run benchmark
```go
 go test -benchmem -run=^$ -bench linio_test.go 
```