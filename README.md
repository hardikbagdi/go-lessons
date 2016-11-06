# go-lessons
Trying out Golang

## Notes

### Introduction 
import is similar to #include

main() is like the all universal main()

Exported functions have the first letter Capitalized (Println)

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

#### Booleans
&& and
!! or
! not
"bool" is a type which can hold the value "false" or "true"


