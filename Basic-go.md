# Basic syntax for Go

## Pointer
1. memory address and change value through pointer

    ```Go
    x := 15
    a := &x // a point to x: a is the memory address of x;
    fmt.Println(a) // hex-value of memoray address
    *a // read through a memory address a to get the data.
    *a = 5 // change the value of x to be 5;
    ```
2. playing around with pointers
    ```go
    *a = *a**a
    fmt.Println(x) // output 25
    fmt.Println(*a) // output 25, Note that 
    ```
3. Class / Struct
    ```go
    
    ```