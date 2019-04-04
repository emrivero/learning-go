# Learning Go
## The newbie way
[<img src="gopher.png">](gopher.png)

Gopher art-work by [egonelbre](https://github.com/egonelbre)

## Index
1. Variable declarations
   * `var` keyword
   * Short assigment `:=`
   
## Content

### 1. Variable declarations 
* `var`keyword
```go
// Typed declarations
var number int = 2 
var flag bool = true
var name string = "Jonh Doe"

// Untyped declarations
var number = 2 
var flag = true
var name = "Jonh Doe"

// You can also group by types
var a, b int = 2, 3

// Untyped version
var a, b = 2, 3

// This way is also correct
var (
      a int = 2
      b bool = false
)

// Untyped version
var (
      a = 2
      b = false
)

// You can also declarate the variable
// and initialize in separate lines
var b int
b = 10

var a string
a = "Hello, Go!"
```
