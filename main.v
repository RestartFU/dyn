module main

import rsh
import os

fn main() {
	res := rsh.parse_script("./dyn-pkg/go/DYNPKG")
	res.run("install")

}
