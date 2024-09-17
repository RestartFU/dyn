module rsh

import os
import net.http
import compress.gzip
import compress.szip

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

fn link(input string, output string) {
	sh("ln -s $(realpath ${input}) $output")
}

fn unzip(input string, output string) {
	szip.extract_zip_to_dir(input, output) or {
		exit(0)
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