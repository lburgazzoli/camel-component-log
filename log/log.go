package log

import (
	"fmt"

	"github.com/lburgazzoli/camel-go/camel"
	"github.com/lburgazzoli/camel-go/core"
)

// ==========================
//
//
//
// ==========================

// Component --
type Component struct {
	core.DefaultComponent
}

// NewComponent --
func NewComponent() camel.Component {
	component := &Component{}

	return component
}

// Process --
func (component *Component) Process(message string) {
	fmt.Printf("%s\n", message)
}
