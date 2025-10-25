package lexer

import (
	"fmt"
	"slices"
)

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSSIGNMENT
	EQUALS
	NOT
	NOT_EQUALS

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	OR
	AND

	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA

	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS

	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// Keywords Reserved
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
)

var reserved_lu map[string]TokenKind = map[string]TokenKind {
	"let":   LET,
	"const": CONST,
	"class": CLASS,
	"new":   NEW,
	"from":  FROM,
	"fn":    FN,
	"if":    IF,
	"else":  ELSE,
	"foreach": FOREACH,
	"while": WHILE,
	"for":   FOR,
	"import": IMPORT,
	"export": EXPORT,
	"typeof": TYPEOF,
	"in":    IN,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) isOneOfMany (expectedTokens ...TokenKind) bool {
	return slices.Contains(expectedTokens, token.Kind)
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSSIGNMENT:
		return "asssignment"
	case EQUALS:
		return "equals"
	case NOT:
		return "not"
	case NOT_EQUALS:
		return "not_equals"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOT_DOT:
		return "dot_dot"
	case SEMI_COLON:
		return "semi_colon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case PLUS:
		return "plus"
	case DASH:
		return "dash"
	case SLASH:
		return "slash"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:
		return "new"
	case FROM:
		return "from"
	case FN:
		return "fn"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case WHILE:
		return "while"
	case FOR:
		return "for"
	case IMPORT:
		return "import"
	case EXPORT:
		return "export"
	case TYPEOF:
		return "typeof"
	case IN:
		return "in"
	default:
		return fmt.Sprintf("unknow(%d)", kind)
	}
}