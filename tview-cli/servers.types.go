package main

type ServerGroup struct {
	ID int
	Name  string
	Desc string
}

type Server struct {
	ID int
	GroupdId int
	IP string
	Name string
	Desc string
	Mac string
	Os string
	Arch string
}