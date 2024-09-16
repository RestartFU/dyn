module rsh

struct Script {
	mut:
	functions map[string]Function
	variables map[string]string
}

pub fn (s Script) variable(identifier string) string {
	return s.variables[identifier]
}

pub fn (s Script) run(function string) {
	func := s.functions[function] or {
		panic("no function found")
	}
	for act in func.actions {
		act()
	}
}