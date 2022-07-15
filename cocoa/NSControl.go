package cocoa

import (
	"github.com/PumpkinSeed/macdriver2/core"
)

type NSControl struct {
	gen_NSControl
}

func NSControl_Init(frame core.NSRect) NSControl {
	return NSControl_alloc().InitWithFrame__asNSControl(frame)
}
