module rsh

import os

struct Script {
	mut:
	functions map[string]Function
	requires []string
	variables map[string]string
}

pub fn (s Script) variable(identifier string) string {
	return s.variables[identifier]
}

pub fn (s Script) run(function string) {
	for r in s.requires {
		os.find_abs_path_of_executable(r) or {
			println("please install '${r}'")
			exit(0)
		}
	}

	func := s.functions[function] or {
		println("no function found with the name ${function}")
		exit(0)
	}
	for act in func.actions {
		act()
	}
}