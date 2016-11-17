package gitignore

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	EOL
	WHITESPACE
	COMMENT
	SEPARATOR
	NEGATION
	PATTERN
	WILDCARD
	ANY
	BAD
)

// String returns a string representation of the Token type.
func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case EOL:
		return "EOL"
	case WHITESPACE:
		return "WHITESPACE"
	case COMMENT:
		return "COMMENT"
	case SEPARATOR:
		return "SEPARATOR"
	case NEGATION:
		return "NEGATION"
	case PATTERN:
		return "PATTERN"
	case WILDCARD:
		return "WILDCARD"
	case ANY:
		return "ANY"
	default:
		return "BAD TOKEN"
	}
} // String()
