package main

var customerDatabase = []Customer{
	{
		1,
		"Logan Roy",
		"CEO",
		"logan@waystar.com",
		"212-111-1111",
		false,
	},
	{
		2,
		"Tom Wambsgans",
		"Head of Media",
		"twambsgans@waystar.com",
		"212-222-2222",
		true,
	},
	{
		3,
		"Nan Pierce",
		"CEO",
		"nan@pgm.com",
		"212-333-3333",
		false,
	},
}

type Customer struct {
	id        uint
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}
