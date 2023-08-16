// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LinkedFunctions] class.
var LinkedFunctionsClass = _LinkedFunctionsClass{objc.GetClass("MTLLinkedFunctions")}

type _LinkedFunctionsClass struct {
	objc.Class
}

// An interface definition for the [LinkedFunctions] class.
type ILinkedFunctions interface {
	objc.IObject
	PrivateFunctions() []FunctionWrapper
	SetPrivateFunctions(value []PFunction)
	BinaryFunctions() []FunctionWrapper
	SetBinaryFunctions(value []PFunction)
	Groups() map[string][]FunctionWrapper
	SetGroups(value map[string][]PFunction)
	Functions() []FunctionWrapper
	SetFunctions(value []PFunction)
}

// A set of related functions that Metal links to when necessary to create the function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions?language=objc
type LinkedFunctions struct {
	objc.Object
}

func LinkedFunctionsFrom(ptr unsafe.Pointer) LinkedFunctions {
	return LinkedFunctions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LinkedFunctionsClass) Alloc() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](lc, objc.Sel("alloc"))
	return rv
}

func LinkedFunctions_Alloc() LinkedFunctions {
	return LinkedFunctionsClass.Alloc()
}

func (lc _LinkedFunctionsClass) New() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLinkedFunctions() LinkedFunctions {
	return LinkedFunctionsClass.New()
}

func (l_ LinkedFunctions) Init() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](l_, objc.Sel("init"))
	return rv
}

// Creates an empty linked functions object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3578281-linkedfunctions?language=objc
func (lc _LinkedFunctionsClass) LinkedFunctions() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](lc, objc.Sel("linkedFunctions"))
	return rv
}

// Creates an empty linked functions object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3578281-linkedfunctions?language=objc
func LinkedFunctions_LinkedFunctions() LinkedFunctions {
	return LinkedFunctionsClass.LinkedFunctions()
}

// An array of function objects to link to the new function, without exporting the functions publicly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3750551-privatefunctions?language=objc
func (l_ LinkedFunctions) PrivateFunctions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](l_, objc.Sel("privateFunctions"))
	return rv
}

// An array of function objects to link to the new function, without exporting the functions publicly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3750551-privatefunctions?language=objc
func (l_ LinkedFunctions) SetPrivateFunctions(value []PFunction) {
	objc.Call[objc.Void](l_, objc.Sel("setPrivateFunctions:"), value)
}

// An array of function objects already compiled to a binary representation to link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3566560-binaryfunctions?language=objc
func (l_ LinkedFunctions) BinaryFunctions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](l_, objc.Sel("binaryFunctions"))
	return rv
}

// An array of function objects already compiled to a binary representation to link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3566560-binaryfunctions?language=objc
func (l_ LinkedFunctions) SetBinaryFunctions(value []PFunction) {
	objc.Call[objc.Void](l_, objc.Sel("setBinaryFunctions:"), value)
}

// An optional list of groups specifying which functions your shader can call at each call site. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3554049-groups?language=objc
func (l_ LinkedFunctions) Groups() map[string][]FunctionWrapper {
	rv := objc.Call[map[string][]FunctionWrapper](l_, objc.Sel("groups"))
	return rv
}

// An optional list of groups specifying which functions your shader can call at each call site. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3554049-groups?language=objc
func (l_ LinkedFunctions) SetGroups(value map[string][]PFunction) {
	objc.Call[objc.Void](l_, objc.Sel("setGroups:"), value)
}

// An array of function objects to link to the new function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3554048-functions?language=objc
func (l_ LinkedFunctions) Functions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](l_, objc.Sel("functions"))
	return rv
}

// An array of function objects to link to the new function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllinkedfunctions/3554048-functions?language=objc
func (l_ LinkedFunctions) SetFunctions(value []PFunction) {
	objc.Call[objc.Void](l_, objc.Sel("setFunctions:"), value)
}
