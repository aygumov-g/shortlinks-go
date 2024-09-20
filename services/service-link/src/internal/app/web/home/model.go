package home

type Link struct {
	LinkId      int    `json:"link_id"`
	LinkAddrIn  string `json:"link_addr_in"`
	LinkAddrOut string `json:"link_addr_out"`
}
