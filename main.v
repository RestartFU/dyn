module main

import rsh

fn main() {
	res := rsh.parse_script("./dyn-pkg/v/DYNPKG")
	res.run("install")

}
