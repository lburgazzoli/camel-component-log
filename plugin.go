package main

import (
	"github.com/lburgazzoli/camel-go-component-log/log"
	"github.com/lburgazzoli/camel-go/camel"
)

// ==========================
//
// plugin entry-pooint
//
// ==========================

// Create --
func Create() camel.Component {
	return log.NewComponent()
}
