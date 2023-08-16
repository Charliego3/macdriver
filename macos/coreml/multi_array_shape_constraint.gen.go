// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MultiArrayShapeConstraint] class.
var MultiArrayShapeConstraintClass = _MultiArrayShapeConstraintClass{objc.GetClass("MLMultiArrayShapeConstraint")}

type _MultiArrayShapeConstraintClass struct {
	objc.Class
}

// An interface definition for the [MultiArrayShapeConstraint] class.
type IMultiArrayShapeConstraint interface {
	objc.IObject
	EnumeratedShapes() [][]foundation.Number
	SizeRangeForDimension() []foundation.Value
	Type() MultiArrayShapeConstraintType
}

// The lists of shapes or ranges of shapes that constrain a multiarray feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint?language=objc
type MultiArrayShapeConstraint struct {
	objc.Object
}

func MultiArrayShapeConstraintFrom(ptr unsafe.Pointer) MultiArrayShapeConstraint {
	return MultiArrayShapeConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MultiArrayShapeConstraintClass) Alloc() MultiArrayShapeConstraint {
	rv := objc.Call[MultiArrayShapeConstraint](mc, objc.Sel("alloc"))
	return rv
}

func MultiArrayShapeConstraint_Alloc() MultiArrayShapeConstraint {
	return MultiArrayShapeConstraintClass.Alloc()
}

func (mc _MultiArrayShapeConstraintClass) New() MultiArrayShapeConstraint {
	rv := objc.Call[MultiArrayShapeConstraint](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMultiArrayShapeConstraint() MultiArrayShapeConstraint {
	return MultiArrayShapeConstraintClass.New()
}

func (m_ MultiArrayShapeConstraint) Init() MultiArrayShapeConstraint {
	rv := objc.Call[MultiArrayShapeConstraint](m_, objc.Sel("init"))
	return rv
}

// Array of allowed shapes for a multiarray feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/2994313-enumeratedshapes?language=objc
func (m_ MultiArrayShapeConstraint) EnumeratedShapes() [][]foundation.Number {
	rv := objc.Call[[][]foundation.Number](m_, objc.Sel("enumeratedShapes"))
	return rv
}

// The allowable range for a dimention of the multiarray. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/2994314-sizerangefordimension?language=objc
func (m_ MultiArrayShapeConstraint) SizeRangeForDimension() []foundation.Value {
	rv := objc.Call[[]foundation.Value](m_, objc.Sel("sizeRangeForDimension"))
	return rv
}

// The type of the shape constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/2994315-type?language=objc
func (m_ MultiArrayShapeConstraint) Type() MultiArrayShapeConstraintType {
	rv := objc.Call[MultiArrayShapeConstraintType](m_, objc.Sel("type"))
	return rv
}
