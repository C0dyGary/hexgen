package domain

type Entity struct {
	Name   string
	Fields []Field
}

type EntityName struct {
	NameLower string
	Name      string
	Project   string
}

type EntityRepository struct {
	NameLower string
	TypeDB    string
}
