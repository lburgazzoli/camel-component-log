package log

import (
	"fmt"
	"github.com/lburgazzoli/camel-go/camel"
)

// *****************************************************************************
//
// 
//
// *****************************************************************************

// LogComponent --
type LogComponent struct {
	context camel.Context
}

// Start --
func (component *LogComponent) Start() {
}

// Stop --
func (component *LogComponent) Stop() {
}

// SetContext --
func (component *LogComponent) SetContext(context camel.Context) {
	component.context = context
}

// Context --
func (component *LogComponent) Context() camel.Context {
	return component.context
}


// Process --
func (component *LogComponent) Process(message string) {
	fmt.Printf("%s\n", message)
}
