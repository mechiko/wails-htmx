package domain

type menu struct {
	Items MenuItemSlice
}

type MenuItem struct {
	Name   string
	Title  string
	Active bool
	Decs   string
	Svg    string
}

type MenuItemSlice []*MenuItem
