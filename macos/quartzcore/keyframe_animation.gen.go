// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [KeyframeAnimation] class.
var KeyframeAnimationClass = _KeyframeAnimationClass{objc.GetClass("CAKeyframeAnimation")}

type _KeyframeAnimationClass struct {
	objc.Class
}

// An interface definition for the [KeyframeAnimation] class.
type IKeyframeAnimation interface {
	IPropertyAnimation
	Path() unsafe.Pointer
	SetPath(value unsafe.Pointer)
	RotationMode() AnimationRotationMode
	SetRotationMode(value AnimationRotationMode)
	BiasValues() []foundation.Number
	SetBiasValues(value []foundation.INumber)
	TensionValues() []foundation.Number
	SetTensionValues(value []foundation.INumber)
	KeyTimes() []foundation.Number
	SetKeyTimes(value []foundation.INumber)
	CalculationMode() AnimationCalculationMode
	SetCalculationMode(value AnimationCalculationMode)
	ContinuityValues() []foundation.Number
	SetContinuityValues(value []foundation.INumber)
	Values() []objc.Object
	SetValues(value []objc.IObject)
	TimingFunctions() []MediaTimingFunction
	SetTimingFunctions(value []IMediaTimingFunction)
}

// An object that provides keyframe animation capabilities for a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation?language=objc
type KeyframeAnimation struct {
	PropertyAnimation
}

func KeyframeAnimationFrom(ptr unsafe.Pointer) KeyframeAnimation {
	return KeyframeAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}

func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := objc.Call[KeyframeAnimation](kc, objc.Sel("alloc"))
	return rv
}

func KeyframeAnimation_Alloc() KeyframeAnimation {
	return KeyframeAnimationClass.Alloc()
}

func (kc _KeyframeAnimationClass) New() KeyframeAnimation {
	rv := objc.Call[KeyframeAnimation](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKeyframeAnimation() KeyframeAnimation {
	return KeyframeAnimationClass.New()
}

func (k_ KeyframeAnimation) Init() KeyframeAnimation {
	rv := objc.Call[KeyframeAnimation](k_, objc.Sel("init"))
	return rv
}

func (kc _KeyframeAnimationClass) AnimationWithKeyPath(path string) KeyframeAnimation {
	rv := objc.Call[KeyframeAnimation](kc, objc.Sel("animationWithKeyPath:"), path)
	return rv
}

// Creates and returns an CAPropertyAnimation instance for the specified key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412534-animationwithkeypath?language=objc
func KeyframeAnimation_AnimationWithKeyPath(path string) KeyframeAnimation {
	return KeyframeAnimationClass.AnimationWithKeyPath(path)
}

func (kc _KeyframeAnimationClass) Animation() KeyframeAnimation {
	rv := objc.Call[KeyframeAnimation](kc, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func KeyframeAnimation_Animation() KeyframeAnimation {
	return KeyframeAnimationClass.Animation()
}

// The path for a point-based property to follow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412474-path?language=objc
func (k_ KeyframeAnimation) Path() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](k_, objc.Sel("path"))
	return rv
}

// The path for a point-based property to follow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412474-path?language=objc
func (k_ KeyframeAnimation) SetPath(value unsafe.Pointer) {
	objc.Call[objc.Void](k_, objc.Sel("setPath:"), value)
}

// Determines whether objects animating along the path rotate to match the path tangent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412454-rotationmode?language=objc
func (k_ KeyframeAnimation) RotationMode() AnimationRotationMode {
	rv := objc.Call[AnimationRotationMode](k_, objc.Sel("rotationMode"))
	return rv
}

// Determines whether objects animating along the path rotate to match the path tangent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412454-rotationmode?language=objc
func (k_ KeyframeAnimation) SetRotationMode(value AnimationRotationMode) {
	objc.Call[objc.Void](k_, objc.Sel("setRotationMode:"), value)
}

// An array of NSNumber objects that define the position of the curve relative to a control point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412485-biasvalues?language=objc
func (k_ KeyframeAnimation) BiasValues() []foundation.Number {
	rv := objc.Call[[]foundation.Number](k_, objc.Sel("biasValues"))
	return rv
}

// An array of NSNumber objects that define the position of the curve relative to a control point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412485-biasvalues?language=objc
func (k_ KeyframeAnimation) SetBiasValues(value []foundation.INumber) {
	objc.Call[objc.Void](k_, objc.Sel("setBiasValues:"), value)
}

// An array of NSNumber objects that define the tightness of the curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412475-tensionvalues?language=objc
func (k_ KeyframeAnimation) TensionValues() []foundation.Number {
	rv := objc.Call[[]foundation.Number](k_, objc.Sel("tensionValues"))
	return rv
}

// An array of NSNumber objects that define the tightness of the curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412475-tensionvalues?language=objc
func (k_ KeyframeAnimation) SetTensionValues(value []foundation.INumber) {
	objc.Call[objc.Void](k_, objc.Sel("setTensionValues:"), value)
}

// An optional array of NSNumber objects that define the time at which to apply a given keyframe segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412522-keytimes?language=objc
func (k_ KeyframeAnimation) KeyTimes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](k_, objc.Sel("keyTimes"))
	return rv
}

// An optional array of NSNumber objects that define the time at which to apply a given keyframe segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412522-keytimes?language=objc
func (k_ KeyframeAnimation) SetKeyTimes(value []foundation.INumber) {
	objc.Call[objc.Void](k_, objc.Sel("setKeyTimes:"), value)
}

// Specifies how intermediate keyframe values are calculated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412500-calculationmode?language=objc
func (k_ KeyframeAnimation) CalculationMode() AnimationCalculationMode {
	rv := objc.Call[AnimationCalculationMode](k_, objc.Sel("calculationMode"))
	return rv
}

// Specifies how intermediate keyframe values are calculated by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412500-calculationmode?language=objc
func (k_ KeyframeAnimation) SetCalculationMode(value AnimationCalculationMode) {
	objc.Call[objc.Void](k_, objc.Sel("setCalculationMode:"), value)
}

// An array of NSNumber objects that define the sharpness of the timing curve’s corners. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412491-continuityvalues?language=objc
func (k_ KeyframeAnimation) ContinuityValues() []foundation.Number {
	rv := objc.Call[[]foundation.Number](k_, objc.Sel("continuityValues"))
	return rv
}

// An array of NSNumber objects that define the sharpness of the timing curve’s corners. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412491-continuityvalues?language=objc
func (k_ KeyframeAnimation) SetContinuityValues(value []foundation.INumber) {
	objc.Call[objc.Void](k_, objc.Sel("setContinuityValues:"), value)
}

// An array of objects that specify the keyframe values to use for the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412498-values?language=objc
func (k_ KeyframeAnimation) Values() []objc.Object {
	rv := objc.Call[[]objc.Object](k_, objc.Sel("values"))
	return rv
}

// An array of objects that specify the keyframe values to use for the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412498-values?language=objc
func (k_ KeyframeAnimation) SetValues(value []objc.IObject) {
	objc.Call[objc.Void](k_, objc.Sel("setValues:"), value)
}

// An optional array of CAMediaTimingFunction objects that define the pacing for each keyframe segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412465-timingfunctions?language=objc
func (k_ KeyframeAnimation) TimingFunctions() []MediaTimingFunction {
	rv := objc.Call[[]MediaTimingFunction](k_, objc.Sel("timingFunctions"))
	return rv
}

// An optional array of CAMediaTimingFunction objects that define the pacing for each keyframe segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/1412465-timingfunctions?language=objc
func (k_ KeyframeAnimation) SetTimingFunctions(value []IMediaTimingFunction) {
	objc.Call[objc.Void](k_, objc.Sel("setTimingFunctions:"), value)
}
