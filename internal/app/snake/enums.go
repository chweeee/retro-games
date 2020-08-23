/**
PUT ALL ENUMS HERE
1) Declare the type of Enum
2) Declare the enum constant with a type and const values of the same type using iota
3) Declare a variable string var that could represent the strings for each one of the constant values
4) Declare a String() function on the declared enum Type
**/

package snake

type keyboardEventType int

// possible user input
const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
)

var userInputs = [...]string{"MOVE", "RETRY", "END"}

// Usage: fmt.Println(keyboardEventType(1))
func (k keyboardEventType) keyboardEventTypeString() string {
	return userInputs[k-1]
}

type directionType int

// Valid snake movement directions
const (
	RIGHT directionType = 1 + iota
	UP
	LEFT
	DOWN
)

var directions = [...]string{"RIGHT", "UP", "LEFT", "DOWN"}

// Usage: fmt.Println(directionType(1))
func (d directionType) directionTypeString() string {
	return directions[d-1]
}
