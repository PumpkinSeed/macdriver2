package webkit

import (
	"github.com/PumpkinSeed/macdriver2/core"
	"github.com/PumpkinSeed/macdriver2/objc"
)

type WKPreferences struct {
	gen_WKPreferences
}

func (p WKPreferences) SetValueForKey(value objc.Object, key core.NSStringRef) {
	p.SetValue_forKey_(value, key)
}
