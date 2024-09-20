package home

type storage interface {
	LinkCreate(link Link) (Link, error)
	LinkSearch(link_addr_in string) (Link, error)
}
