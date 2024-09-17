module rsh

const (
	tok_eof = TokenKind{identifier: "EOF"}
	tok_invalid = TokenKind{identifier: "invalid"}
	tok_string = TokenKind{identifier: "string"}
	tok_variable = TokenKind{identifier: "variable"}
	tok_to = TokenKind{identifier: "to"}
	tok_function = TokenKind{identifier: "function"}
	tok_comment = TokenKind{identifier: "comment"}
	tok_identifier = TokenKind{identifier: "identifier"}
	tok_equals = TokenKind{identifier: "equals"}

	tok_left_bracket = TokenKind{identifier: "left bracket"}
	tok_right_bracket = TokenKind{identifier: "right bracket"}
	tok_require = TokenKind{identifier: "require"}
)
