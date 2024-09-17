module rsh

import os

struct Position {
mut:
	offset int
	line int = 1
	cursor int
}

struct Token {
	Position
mut:
	kind TokenKind
	text string
}

struct TokenKind {
	identifier string
}

struct Tokenizer {
	Position
mut:
	r rune
	data string
}


fn (mut t Tokenizer) next() rune {
	t.offset++
	t.cursor++

	dat := t.data[t.offset-1..]
	if t.offset > t.data.len  || dat.len <= 0 {
		return 0
	}

	r := dat[0]
	if r == `\n` {
		t.line++
		t.cursor=0
	}
	return r
}

fn (mut t Tokenizer) previous() rune {
	t.offset--
	t.cursor--

	dat := t.data[t.offset..]
	if t.offset > t.data.len  || dat.len <= 0 {
		return 0
	}

	r := dat[0]
	if r == `\n` {
		t.line--
	}
	return r
}

fn (mut t Tokenizer) skip_white_space() {
	t.r = t.next()
	for is_any(t.r, `\n`, ` `, `\t`, `\r`, `\f`) {
		t.r = t.next()
	}
}

fn (mut t Tokenizer) token() Token {
	t.skip_white_space()

	mut tok := Token{}
	tok.Position = t.Position
	tok.kind = tok_invalid

	match t.r {
		0 {
			tok.kind = tok_eof
			return tok
		}
		`$` {
			t.r = t.next()
			if t.r == `(` {
				tok.kind = tok_string
				for t.offset < t.data.len {
					t.r = t.next()
					if t.r == `)` {
						break
					}
					tok.text += t.r.str()
				}
				tok.text = os.raw_execute(tok.text).output.trim_space()
			}
		}
		`#` {
			tok.kind = tok_comment

			for t.offset < t.data.len {
				t.r = t.next()
				if t.r == `\n` {
					break
				}
				tok.text += t.r.str()
			}
		}
		`=` {
			tok.kind = tok_equals
			tok.text = "="
		}
		`{` {
			tok.kind = tok_left_bracket
			tok.text = "{"
		}
		`}` {
			tok.kind = tok_right_bracket
			tok.text = "}"
		}
		`"` {
			tok.kind = tok_string
			for t.offset < t.data.len {
				t.r = t.next()
				match t.r {
					`$` {
						t.r = t.next()
						if t.r == `(` {
							mut shcmd := ""
							for t.offset < t.data.len {
								if t.r == `)` {
									tok.text += os.raw_execute(shcmd[1..]).output.trim_space()
									t.r = t.next()
									break
								}
								shcmd += t.r.str()
								t.r = t.next()
							}
						} else {
							tok.text += "$"
						}
					}
					`"` {
						break
					}
					else {}
				}
				tok.text += t.r.str()
			}
		}
		else {
			tok.kind = tok_identifier
			for t.offset < t.data.len {
				if is_any(t.r,`\n`, ` `, `\t`, `\r`, `\f`) {
					break
				}

				tok.text += t.r.str()
				t.r = t.next()

				mut should_break := true

				match tok.text{
					"to" {
						tok.kind = tok_to
					}
					"variable"{
						tok.kind = tok_variable
					}
					"function" {
						tok.kind = tok_function
					}
					"require" {
						tok.kind = tok_require
					}
					else {
						should_break = false
					}
				}

				if should_break {
					break
				}
			}
		}
	}

	println("${tok.line}:${tok.cursor} -> '${tok.text}' len[${tok.text.len}] kind[${tok.kind.identifier}]")
	return tok
}

fn is_any(r rune, runes ...rune) bool {
	for v in runes {
		if r == v {
			return true
		}
	}
	return false
}