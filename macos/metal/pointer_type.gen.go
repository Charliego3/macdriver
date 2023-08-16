// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PointerType] class.
var PointerTypeClass = _PointerTypeClass{objc.GetClass("MTLPointerType")}

type _PointerTypeClass struct {
	objc.Class
}

// An interface definition for the [PointerType] class.
type IPointerType interface {
	IType
	ElementArrayType() ArrayType
	ElementStructType() StructType
	Alignment() uint
	DataSize() uint
	ElementType() DataType
	Access() objc.Object
	ElementIsArgumentBuffer() bool
}

// A description of a pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype?language=objc
type PointerType struct {
	Type
}

func PointerTypeFrom(ptr unsafe.Pointer) PointerType {
	return PointerType{
		Type: TypeFrom(ptr),
	}
}

func (pc _PointerTypeClass) Alloc() PointerType {
	rv := objc.Call[PointerType](pc, objc.Sel("alloc"))
	return rv
}

func PointerType_Alloc() PointerType {
	return PointerTypeClass.Alloc()
}

func (pc _PointerTypeClass) New() PointerType {
	rv := objc.Call[PointerType](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPointerType() PointerType {
	return PointerTypeClass.New()
}

func (p_ PointerType) Init() PointerType {
	rv := objc.Call[PointerType](p_, objc.Sel("init"))
	return rv
}

// Provides a description of the underlying array when the pointer points to an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2915757-elementarraytype?language=objc
func (p_ PointerType) ElementArrayType() ArrayType {
	rv := objc.Call[ArrayType](p_, objc.Sel("elementArrayType"))
	return rv
}

// Provides a description of the underlying struct when the pointer points to a struct. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2915755-elementstructtype?language=objc
func (p_ PointerType) ElementStructType() StructType {
	rv := objc.Call[StructType](p_, objc.Sel("elementStructType"))
	return rv
}

// The required byte alignment in memory for the element data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2877458-alignment?language=objc
func (p_ PointerType) Alignment() uint {
	rv := objc.Call[uint](p_, objc.Sel("alignment"))
	return rv
}

// The size, in bytes, of the element data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2877451-datasize?language=objc
func (p_ PointerType) DataSize() uint {
	rv := objc.Call[uint](p_, objc.Sel("dataSize"))
	return rv
}

// The data type of the element data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2877457-elementtype?language=objc
func (p_ PointerType) ElementType() DataType {
	rv := objc.Call[DataType](p_, objc.Sel("elementType"))
	return rv
}

// The function’s read/write access to the element data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2877452-access?language=objc
func (p_ PointerType) Access() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("access"))
	return rv
}

// A Boolean value that indicates whether the element is an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpointertype/2915756-elementisargumentbuffer?language=objc
func (p_ PointerType) ElementIsArgumentBuffer() bool {
	rv := objc.Call[bool](p_, objc.Sel("elementIsArgumentBuffer"))
	return rv
}
