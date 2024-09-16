module rsh

import os
import net.http

struct Function {
mut:
	actions []fn ()
}

fn delete(path string) {
	sh("rm -rf ${path}")
}

fn sh(s string) {
	println(s)
	os.system(s)
}

fn (p Parser) internal(line int, cursor int, s string, res &Script) {
	fnc := res.functions[s] or {
		println("${p.filename}: ${line}:${cursor}: no function with the name ${s} found")
		exit(0)
	}
	for act in fnc.actions {
		act()
	}
}

fn move(input string, output string) {
	sh("mv ${input} ${output}")
}

fn (p Parser) download(line int, url string, output string) {
	http.download_file_with_progress(url, output) or {
		println('${p.filename}: ${line}:0 could not download file from ${url} to ${output}')
		exit(0)
	}
}