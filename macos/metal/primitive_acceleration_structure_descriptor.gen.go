// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PrimitiveAccelerationStructureDescriptor] class.
var PrimitiveAccelerationStructureDescriptorClass = _PrimitiveAccelerationStructureDescriptorClass{objc.GetClass("MTLPrimitiveAccelerationStructureDescriptor")}

type _PrimitiveAccelerationStructureDescriptorClass struct {
	objc.Class
}

// An interface definition for the [PrimitiveAccelerationStructureDescriptor] class.
type IPrimitiveAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
	MotionEndTime() float64
	SetMotionEndTime(value float64)
	MotionStartBorderMode() MotionBorderMode
	SetMotionStartBorderMode(value MotionBorderMode)
	MotionEndBorderMode() MotionBorderMode
	SetMotionEndBorderMode(value MotionBorderMode)
	MotionStartTime() float64
	SetMotionStartTime(value float64)
	GeometryDescriptors() []AccelerationStructureGeometryDescriptor
	SetGeometryDescriptors(value []IAccelerationStructureGeometryDescriptor)
	MotionKeyframeCount() uint
	SetMotionKeyframeCount(value uint)
}

// A description of an acceleration structure that contains geometry primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor?language=objc
type PrimitiveAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

func PrimitiveAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) PrimitiveAccelerationStructureDescriptor {
	return PrimitiveAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

func (pc _PrimitiveAccelerationStructureDescriptorClass) Descriptor() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Call[PrimitiveAccelerationStructureDescriptor](pc, objc.Sel("descriptor"))
	return rv
}

// Creates a new primitive descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3553891-descriptor?language=objc
func PrimitiveAccelerationStructureDescriptor_Descriptor() PrimitiveAccelerationStructureDescriptor {
	return PrimitiveAccelerationStructureDescriptorClass.Descriptor()
}

func (pc _PrimitiveAccelerationStructureDescriptorClass) Alloc() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Call[PrimitiveAccelerationStructureDescriptor](pc, objc.Sel("alloc"))
	return rv
}

func PrimitiveAccelerationStructureDescriptor_Alloc() PrimitiveAccelerationStructureDescriptor {
	return PrimitiveAccelerationStructureDescriptorClass.Alloc()
}

func (pc _PrimitiveAccelerationStructureDescriptorClass) New() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Call[PrimitiveAccelerationStructureDescriptor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPrimitiveAccelerationStructureDescriptor() PrimitiveAccelerationStructureDescriptor {
	return PrimitiveAccelerationStructureDescriptorClass.New()
}

func (p_ PrimitiveAccelerationStructureDescriptor) Init() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Call[PrimitiveAccelerationStructureDescriptor](p_, objc.Sel("init"))
	return rv
}

// The end time for the range of motion that the keyframe data describes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750510-motionendtime?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) MotionEndTime() float64 {
	rv := objc.Call[float64](p_, objc.Sel("motionEndTime"))
	return rv
}

// The end time for the range of motion that the keyframe data describes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750510-motionendtime?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMotionEndTime:"), value)
}

// The mode to use when handling timestamps before the start time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750512-motionstartbordermode?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() MotionBorderMode {
	rv := objc.Call[MotionBorderMode](p_, objc.Sel("motionStartBorderMode"))
	return rv
}

// The mode to use when handling timestamps before the start time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750512-motionstartbordermode?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value MotionBorderMode) {
	objc.Call[objc.Void](p_, objc.Sel("setMotionStartBorderMode:"), value)
}

// The mode to use when handling timestamps after the end time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750509-motionendbordermode?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() MotionBorderMode {
	rv := objc.Call[MotionBorderMode](p_, objc.Sel("motionEndBorderMode"))
	return rv
}

// The mode to use when handling timestamps after the end time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750509-motionendbordermode?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value MotionBorderMode) {
	objc.Call[objc.Void](p_, objc.Sel("setMotionEndBorderMode:"), value)
}

// The start time for the range of motion that the keyframe data describes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750513-motionstarttime?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) MotionStartTime() float64 {
	rv := objc.Call[float64](p_, objc.Sel("motionStartTime"))
	return rv
}

// The start time for the range of motion that the keyframe data describes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750513-motionstarttime?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMotionStartTime:"), value)
}

// An array that contains the individual pieces of geometry that compose the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3553892-geometrydescriptors?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) GeometryDescriptors() []AccelerationStructureGeometryDescriptor {
	rv := objc.Call[[]AccelerationStructureGeometryDescriptor](p_, objc.Sel("geometryDescriptors"))
	return rv
}

// An array that contains the individual pieces of geometry that compose the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3553892-geometrydescriptors?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value []IAccelerationStructureGeometryDescriptor) {
	objc.Call[objc.Void](p_, objc.Sel("setGeometryDescriptors:"), value)
}

// The number of keyframes in the geometry data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750511-motionkeyframecount?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() uint {
	rv := objc.Call[uint](p_, objc.Sel("motionKeyframeCount"))
	return rv
}

// The number of keyframes in the geometry data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/3750511-motionkeyframecount?language=objc
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setMotionKeyframeCount:"), value)
}
