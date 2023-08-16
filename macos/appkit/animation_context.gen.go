// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AnimationContext] class.
var AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}

type _AnimationContextClass struct {
	objc.Class
}

// An interface definition for the [AnimationContext] class.
type IAnimationContext interface {
	objc.IObject
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
	CompletionHandler() func()
	SetCompletionHandler(value func())
	TimingFunction() quartzcore.MediaTimingFunction
	SetTimingFunction(value quartzcore.IMediaTimingFunction)
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
}

// An animation context, which contains information about environment and state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext?language=objc
type AnimationContext struct {
	objc.Object
}

func AnimationContextFrom(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := objc.Call[AnimationContext](ac, objc.Sel("alloc"))
	return rv
}

func AnimationContext_Alloc() AnimationContext {
	return AnimationContextClass.Alloc()
}

func (ac _AnimationContextClass) New() AnimationContext {
	rv := objc.Call[AnimationContext](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAnimationContext() AnimationContext {
	return AnimationContextClass.New()
}

func (a_ AnimationContext) Init() AnimationContext {
	rv := objc.Call[AnimationContext](a_, objc.Sel("init"))
	return rv
}

// Creates a new animation grouping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1530234-begingrouping?language=objc
func (ac _AnimationContextClass) BeginGrouping() {
	objc.Call[objc.Void](ac, objc.Sel("beginGrouping"))
}

// Creates a new animation grouping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1530234-begingrouping?language=objc
func AnimationContext_BeginGrouping() {
	AnimationContextClass.BeginGrouping()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/2967169-runanimationgroup?language=objc
func (ac _AnimationContextClass) RunAnimationGroup(changes func(context AnimationContext)) {
	objc.Call[objc.Void](ac, objc.Sel("runAnimationGroup:"), changes)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/2967169-runanimationgroup?language=objc
func AnimationContext_RunAnimationGroup(changes func(context AnimationContext)) {
	AnimationContextClass.RunAnimationGroup(changes)
}

// Ends the current animation grouping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1526927-endgrouping?language=objc
func (ac _AnimationContextClass) EndGrouping() {
	objc.Call[objc.Void](ac, objc.Sel("endGrouping"))
}

// Ends the current animation grouping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1526927-endgrouping?language=objc
func AnimationContext_EndGrouping() {
	AnimationContextClass.EndGrouping()
}

// Returns the current animation context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1535845-currentcontext?language=objc
func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := objc.Call[AnimationContext](ac, objc.Sel("currentContext"))
	return rv
}

// Returns the current animation context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1535845-currentcontext?language=objc
func AnimationContext_CurrentContext() AnimationContext {
	return AnimationContextClass.CurrentContext()
}

// Determine if animations are enabled or not for animations that occur as a result of another property change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1525870-allowsimplicitanimation?language=objc
func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := objc.Call[bool](a_, objc.Sel("allowsImplicitAnimation"))
	return rv
}

// Determine if animations are enabled or not for animations that occur as a result of another property change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1525870-allowsimplicitanimation?language=objc
func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAllowsImplicitAnimation:"), value)
}

// A completion Block that is called when the animations in the grouping are completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1531132-completionhandler?language=objc
func (a_ AnimationContext) CompletionHandler() func() {
	rv := objc.Call[func()](a_, objc.Sel("completionHandler"))
	return rv
}

// A completion Block that is called when the animations in the grouping are completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1531132-completionhandler?language=objc
func (a_ AnimationContext) SetCompletionHandler(value func()) {
	objc.Call[objc.Void](a_, objc.Sel("setCompletionHandler:"), value)
}

// The timing function used for all animations within this animation proxy group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1524985-timingfunction?language=objc
func (a_ AnimationContext) TimingFunction() quartzcore.MediaTimingFunction {
	rv := objc.Call[quartzcore.MediaTimingFunction](a_, objc.Sel("timingFunction"))
	return rv
}

// The timing function used for all animations within this animation proxy group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1524985-timingfunction?language=objc
func (a_ AnimationContext) SetTimingFunction(value quartzcore.IMediaTimingFunction) {
	objc.Call[objc.Void](a_, objc.Sel("setTimingFunction:"), objc.Ptr(value))
}

// The duration used by animations created as a result of setting new values for an animatable property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1526780-duration?language=objc
func (a_ AnimationContext) Duration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](a_, objc.Sel("duration"))
	return rv
}

// The duration used by animations created as a result of setting new values for an animatable property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/1526780-duration?language=objc
func (a_ AnimationContext) SetDuration(value foundation.TimeInterval) {
	objc.Call[objc.Void](a_, objc.Sel("setDuration:"), value)
}
