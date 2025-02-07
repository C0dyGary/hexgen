package domain

type Method struct {
	NameMethod string
	Args       []Arg
	Returns    []string
}

type Arg struct {
	NameArg string
	TypeArg string
}
