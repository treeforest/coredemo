package main

import "github.com/treeforest/coredemo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
