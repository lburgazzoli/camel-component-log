package main

import (
	"github.com/lburgazzoli/camel-go/camel"
	"github.com/lburgazzoli/camel-go-component-log/log"
)

// ========================================
// plugin entry-pooint
// ========================================

// Create --
func Create() camel.Component {
	return &log.LogComponent{}
}