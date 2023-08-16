// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Animation] class.
var AnimationClass = _AnimationClass{objc.GetClass("CAAnimation")}

type _AnimationClass struct {
	objc.Class
}

// An interface definition for the [Animation] class.
type IAnimation interface {
	objc.IObject
	ShouldArchiveValueForKey(key string) bool
	IsRemovedOnCompletion() bool
	SetRemovedOnCompletion(value bool)
	UsesSceneTimeBase() bool
	SetUsesSceneTimeBase(value bool)
	Delegate() AnimationDelegateWrapper
	SetDelegate(value PAnimationDelegate)
	SetDelegateObject(valueObject objc.IObject)
	FadeInDuration() float64
	SetFadeInDuration(value float64)
	FadeOutDuration() float64
	SetFadeOutDuration(value float64)
	TimingFunction() MediaTimingFunction
	SetTimingFunction(value IMediaTimingFunction)
	PreferredFrameRateRange() FrameRateRange
	SetPreferredFrameRateRange(value FrameRateRange)
}

// The abstract superclass for animations in Core Animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation?language=objc
type Animation struct {
	objc.Object
}

func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AnimationClass) Animation() Animation {
	rv := objc.Call[Animation](ac, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func Animation_Animation() Animation {
	return AnimationClass.Animation()
}

func (ac _AnimationClass) Alloc() Animation {
	rv := objc.Call[Animation](ac, objc.Sel("alloc"))
	return rv
}

func Animation_Alloc() Animation {
	return AnimationClass.Alloc()
}

func (ac _AnimationClass) New() Animation {
	rv := objc.Call[Animation](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := objc.Call[Animation](a_, objc.Sel("init"))
	return rv
}

// Specifies the default value of the property with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412530-defaultvalueforkey?language=objc
func (ac _AnimationClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](ac, objc.Sel("defaultValueForKey:"), key)
	return rv
}

// Specifies the default value of the property with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412530-defaultvalueforkey?language=objc
func Animation_DefaultValueForKey(key string) objc.Object {
	return AnimationClass.DefaultValueForKey(key)
}

// Specifies whether the value of the property for a given key is archived. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412525-shouldarchivevalueforkey?language=objc
func (a_ Animation) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Call[bool](a_, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}

// Determines if the animation is removed from the target layer’s animations upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412458-removedoncompletion?language=objc
func (a_ Animation) IsRemovedOnCompletion() bool {
	rv := objc.Call[bool](a_, objc.Sel("isRemovedOnCompletion"))
	return rv
}

// Determines if the animation is removed from the target layer’s animations upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412458-removedoncompletion?language=objc
func (a_ Animation) SetRemovedOnCompletion(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setRemovedOnCompletion:"), value)
}

// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1523819-usesscenetimebase?language=objc
func (a_ Animation) UsesSceneTimeBase() bool {
	rv := objc.Call[bool](a_, objc.Sel("usesSceneTimeBase"))
	return rv
}

// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1523819-usesscenetimebase?language=objc
func (a_ Animation) SetUsesSceneTimeBase(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setUsesSceneTimeBase:"), value)
}

// Specifies the receiver’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412490-delegate?language=objc
func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := objc.Call[AnimationDelegateWrapper](a_, objc.Sel("delegate"))
	return rv
}

// Specifies the receiver’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412490-delegate?language=objc
func (a_ Animation) SetDelegate(value PAnimationDelegate) {
	po0 := objc.WrapAsProtocol("CAAnimationDelegate", value)
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), po0)
}

// Specifies the receiver’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412490-delegate?language=objc
func (a_ Animation) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1523370-fadeinduration?language=objc
func (a_ Animation) FadeInDuration() float64 {
	rv := objc.Call[float64](a_, objc.Sel("fadeInDuration"))
	return rv
}

// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1523370-fadeinduration?language=objc
func (a_ Animation) SetFadeInDuration(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setFadeInDuration:"), value)
}

// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1522959-fadeoutduration?language=objc
func (a_ Animation) FadeOutDuration() float64 {
	rv := objc.Call[float64](a_, objc.Sel("fadeOutDuration"))
	return rv
}

// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1522959-fadeoutduration?language=objc
func (a_ Animation) SetFadeOutDuration(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setFadeOutDuration:"), value)
}

// An optional timing function defining the pacing of the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412456-timingfunction?language=objc
func (a_ Animation) TimingFunction() MediaTimingFunction {
	rv := objc.Call[MediaTimingFunction](a_, objc.Sel("timingFunction"))
	return rv
}

// An optional timing function defining the pacing of the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412456-timingfunction?language=objc
func (a_ Animation) SetTimingFunction(value IMediaTimingFunction) {
	objc.Call[objc.Void](a_, objc.Sel("setTimingFunction:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/3875342-preferredframeraterange?language=objc
func (a_ Animation) PreferredFrameRateRange() FrameRateRange {
	rv := objc.Call[FrameRateRange](a_, objc.Sel("preferredFrameRateRange"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/3875342-preferredframeraterange?language=objc
func (a_ Animation) SetPreferredFrameRateRange(value FrameRateRange) {
	objc.Call[objc.Void](a_, objc.Sel("setPreferredFrameRateRange:"), value)
}
