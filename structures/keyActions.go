package structures

// KeyActions is a list of KeyAction
type KeyActions []KeyAction

// KeyAction is one goal to achieve an key result.
type KeyAction struct {
	Description
	Grade
}
