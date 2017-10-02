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
	camel.Component
}

// NewComponent --
func NewComponent() camel.Component {
	return &Component{
		Component: core.NewComponent(1),
	}
}

// Process --
func (component *Component) Process(message string) {
	fmt.Printf("%s\n", message)
}
