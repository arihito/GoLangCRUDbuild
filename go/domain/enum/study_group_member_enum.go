package enum

type Kind int

// study group member type
//   0: administrator
//   1: member
const (
	ADMINISTRATOR Kind = iota
	MEMBER
)
