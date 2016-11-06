# go-lessons
Trying out Golang

## Notes
Book being referred: An Introduction to Programming in Go by Caleb Doxsey

### Introduction 
import is similar to #include

main() is like the all universal main()

Exported functions have the first letter Capitalized (Println)
fmt.Println() or fmt.Printf() can be used for prints for stdout
fmt.Scanf() (same/similar as c) for reading input from stdin
### Types

#### Integers
uint8,uint16,uint32,uint64,int8,int16,int32,int64
int means signed integer
byte is same as uint8 (byte)
rune is same as int32 (normal int?)

uint,int and uintptr are arch specific types.

#### Floating Point
are inexact so don't use them for comparisions (same in other langs)

32 and 64-bit varients
In addition to numbers, has "Nan" to represent 0/0 and infinity.

float32,float64 are the datatypes. float64 is like double in Java
complex64 and complex128 for representing complex numbers. I should use this sometime.


#### Character
a simple char in a string "Hello World" is represented as a byte i.e. ASCII

#### String
len() - length of string
zero indexed
concat  can be done using + sign.
== can be used to compare two strings(values)

#### Booleans
&& and
!! or
! not
"bool" is a type which can hold the value "false" or "true"

### Variables

var <name> <type> is the way to declare a variable

var <name> := <value> then type can be skipped; compiler will infer
 <name> := <value>  var is also options (: denotes initialization)

You can declare multiple variables using following syntax; more readable in every function; multiple datatypes are allowed
var (
	a = 42
	b = 3.14
)
var (
	a string //can also specify data-types
	b bool 
)
### Constants
const <name>  <type> = <value> 
cannot be changed once declared
Above multiple variable declaration can be applied to constants too

### Arithematic
Normal operations + - * / %
Shorthand operators available += and others
++ -- standard increment and decrement operators

### Scoping
Global definitions accessible to all functions in the file
Normal function closures exist
Lexically scoped using blocks.

### Control Structures

#### if else
if <condition> { //no () needed around the condition
} else if <condition> {
} else {
}


#### Loop
only one loop - for

eg- 
i := 1
for i <= 10 { //only validation here; this is the while loop in go
	i = i+1
}
OR
for i:= 1 ; i <=10 ; i++ { //normal standard for
}

#### switch
kinda cool because you don't need to write 'break';less debugging effort

switch i {
case 42:
case 43:
default:
}
