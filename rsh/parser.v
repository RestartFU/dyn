module rsh

import os

struct Parser {
	filename string
mut:
	tokenizer Tokenizer
	result Script
}

fn (mut p Parser) expect(kind TokenKind) Token {
	mut tok := p.tokenizer.token()
	if tok.kind != kind {
		println("${p.filename}: ${p.tokenizer.line}:${p.tokenizer.cursor}: expected token kind '${kind.identifier}' but got '${tok.kind.identifier}' instead")
		exit(0)
	}
	if tok.kind == tok_string {
		for id, v in p.result.variables {
			println(tok.text)
			tok.text = tok.text.replace("$" + "{" + id + "}", v)
			println(tok.text)
		}
	}

	return tok
}

fn (mut p Parser) parse_function() {
	identifier := p.expect(tok_identifier).text
	p.result.functions[identifier] = Function{}

	p.expect(tok_left_bracket)

	mut tok := p.tokenizer.token()
	for tok.kind != tok_right_bracket {
		line := tok.line
		match tok.text {
			"download" {
				url := p.expect(tok_string).text
				p.expect(tok_to)
				output := p.expect(tok_string).text

				p.result.functions[identifier].actions << fn [p, url, output, line] () { p.download(line, url, output) }
			}
			"sh" {
				s := p.expect(tok_string).text
				p.result.functions[identifier].actions << fn[s] () { sh(s) }
			}
			"internal" {
				res := &p.result
				func := p.expect(tok_identifier)
				p.result.functions[identifier].actions << fn[func, res, p, line] () { p.internal(line, func.cursor,func.text, res) }
			}
			"delete" {
				s := p.expect(tok_string).text
				p.result.functions[identifier].actions << fn[s] () { delete(s) }
			}
			"move" {
				from := p.expect(tok_string).text
				p.expect(tok_to)
				to := p.expect(tok_string).text

				p.result.functions[identifier].actions << fn [from, to] () { move(from, to) }
			}
			else {}
		}
		tok = p.tokenizer.token()
	}
}

fn (mut p Parser) parse_variable() {
	identifier := p.expect(tok_identifier).text
	p.expect(tok_equals)
	value := p.expect(tok_string).text
	p.result.variables[identifier] = value
}

pub fn parse_script(filename string) Script {
	data := os.read_file(filename) or {
		panic(err)
	}

	mut parser := &Parser{
		filename: os.file_name(filename)
		tokenizer: Tokenizer{data: data}
	}

	mut tok := parser.tokenizer.token()
	for tok.kind != tok_eof {
		match tok.kind {
			tok_variable {
				parser.parse_variable()
			}
			tok_function {
				parser.parse_function()
			}
			else {}
		}
		tok = parser.tokenizer.token()
	}

	return parser.result
}