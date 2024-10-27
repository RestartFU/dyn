module main

import rsh

fn main() {
	res := rsh.parse_script("./dyn-pkg/go/DYNPKG")
	res.run("install")
}
