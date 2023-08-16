// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ClickGestureRecognizer] class.
var ClickGestureRecognizerClass = _ClickGestureRecognizerClass{objc.GetClass("NSClickGestureRecognizer")}

type _ClickGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [ClickGestureRecognizer] class.
type IClickGestureRecognizer interface {
	IGestureRecognizer
	ButtonMask() uint
	SetButtonMask(value uint)
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
	NumberOfClicksRequired() int
	SetNumberOfClicksRequired(value int)
}

// A discrete gesture recognizer that tracks a specified number of mouse clicks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer?language=objc
type ClickGestureRecognizer struct {
	GestureRecognizer
}

func ClickGestureRecognizerFrom(ptr unsafe.Pointer) ClickGestureRecognizer {
	return ClickGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

func (cc _ClickGestureRecognizerClass) Alloc() ClickGestureRecognizer {
	rv := objc.Call[ClickGestureRecognizer](cc, objc.Sel("alloc"))
	return rv
}

func ClickGestureRecognizer_Alloc() ClickGestureRecognizer {
	return ClickGestureRecognizerClass.Alloc()
}

func (cc _ClickGestureRecognizerClass) New() ClickGestureRecognizer {
	rv := objc.Call[ClickGestureRecognizer](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewClickGestureRecognizer() ClickGestureRecognizer {
	return ClickGestureRecognizerClass.New()
}

func (c_ ClickGestureRecognizer) Init() ClickGestureRecognizer {
	rv := objc.Call[ClickGestureRecognizer](c_, objc.Sel("init"))
	return rv
}

func (c_ ClickGestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) ClickGestureRecognizer {
	rv := objc.Call[ClickGestureRecognizer](c_, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Initializes the gesture recognizer with the specified target and action information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535012-initwithtarget?language=objc
func ClickGestureRecognizer_InitWithTargetAction(target objc.IObject, action objc.Selector) ClickGestureRecognizer {
	return ClickGestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
}

// A bit mask of the button (or buttons) required to recognize this click. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/1530136-buttonmask?language=objc
func (c_ ClickGestureRecognizer) ButtonMask() uint {
	rv := objc.Call[uint](c_, objc.Sel("buttonMask"))
	return rv
}

// A bit mask of the button (or buttons) required to recognize this click. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/1530136-buttonmask?language=objc
func (c_ ClickGestureRecognizer) SetButtonMask(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setButtonMask:"), value)
}

// The number of touches required in an NSTouchBar object for the gesture recognizer to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/2544764-numberoftouchesrequired?language=objc
func (c_ ClickGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfTouchesRequired"))
	return rv
}

// The number of touches required in an NSTouchBar object for the gesture recognizer to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/2544764-numberoftouchesrequired?language=objc
func (c_ ClickGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfTouchesRequired:"), value)
}

// The number of clicks required to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/1534485-numberofclicksrequired?language=objc
func (c_ ClickGestureRecognizer) NumberOfClicksRequired() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfClicksRequired"))
	return rv
}

// The number of clicks required to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/1534485-numberofclicksrequired?language=objc
func (c_ ClickGestureRecognizer) SetNumberOfClicksRequired(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfClicksRequired:"), value)
}
