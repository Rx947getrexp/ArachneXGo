package iface

import "golang.org/x/net/html"

type ISpider interface {
	Run()
	Fetch(string)
	ParseLinks(*html.Node)
}
