package cocoa

import (
	"github.com/PumpkinSeed/macdriver2/core"
)

type NSLayoutManager struct {
	gen_NSLayoutManager
}

func (lm NSLayoutManager) EnsureLayoutForTextContainer(tc NSTextContainerRef) {
	lm.EnsureLayoutForTextContainer_(tc)
}

func (lm NSLayoutManager) UsedRectForTextContainer(tc NSTextContainerRef) (rect core.NSRect) {
	return lm.UsedRectForTextContainer_(tc)
}
