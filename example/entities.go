package example

type author struct {
	name string
}

type book struct {
	Name   string
	ISBN   string
	Author *author
}
