package iface

type ISpider interface {
	Run()
	Fetch(string)
}
