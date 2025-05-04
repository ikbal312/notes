package main

// inline comment

/*
	multiline comment
*/

/*
 go variable types
*/

/*
 var x int    // declare
 x = 5		// assigned

 y := 5     declared and assigned
*



/*
GO Types

int / int8 / int16 / int32 / int64
uint / uint8 / uint16 / uint32 / uint64

bool		// boolean either false or true
string		// a string with UTF-8 character
float 		// float32
complex12 	// complex number with float64 real and imaginary part
complex64	// complex number with float32 real and imaginary part

bytes		// alias for uint8
rune		// alias for int32


array		// a non-growable list of items
slice
map			// key-value pair similar to python dictionaries
struct		// collection of named attribute(variable), similar to python objects
interface	// a type that holds value with specific defined method
pointers	// a type that stores the memory address of a variable,
channels	// a buffered or non-buffered pipe for sending data asynchronously


function	// function that can stored in a variable
*/

/*
GO Default Value for Variable or value of variable when it declared not assigned

(u)int/8/16/32/64	| 0
string				| ""
bool				| false
float32/64			| 0.0
slice				| nill
map 				| nill
struct				| {} with all value with zeroes values
interface 			| nill
pointers			| nill
channels			| nil
bytes				| 0
rune				| 0
array				| zero-length array
function			| nil

*/

/*

Go Variable Scope
1. Package scoped		// can seen by the entire package and declared outside a function
2. Function scoped		// can seen within {} which defines the function
3. Statement Scoped		// can be seen within {} of a statement in function (for loop, if/else)

Go variable Shadowing
local scope shadow outer scope variable if they have the same name.

*/

/*

GO Variable Declaration
*/
// var x int			// x = 0
// var name string		// name = ""
// var balance float32 	// balance = 0.0

// var array [3]int		// array = {0,0,0}
// var slice = [5]int{}

/*
Array

declaration
var array_name [size] type
exp:
var fixedSizeArray [5]int
*/
