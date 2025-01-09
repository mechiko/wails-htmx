package domain

type PageInfo struct {
	Name      string
	Url       string
	MenuTitle string
	Desc      string
	Svg       string
}

type MapPageInfo map[string]*PageInfo
