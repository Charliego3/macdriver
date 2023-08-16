// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Kernel] class.
var KernelClass = _KernelClass{objc.GetClass("CIKernel")}

type _KernelClass struct {
	objc.Class
}

// An interface definition for the [Kernel] class.
type IKernel interface {
	objc.IObject
	SetROISelector(method objc.Selector)
	ApplyWithExtentRoiCallbackArguments(extent coregraphics.Rect, callback KernelROICallback, args []objc.IObject) Image
	Name() string
}

// A GPU-based image processing routine used to create custom Core Image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel?language=objc
type Kernel struct {
	objc.Object
}

func KernelFrom(ptr unsafe.Pointer) Kernel {
	return Kernel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (kc _KernelClass) KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) Kernel {
	rv := objc.Call[Kernel](kc, objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), name, data, format, objc.Ptr(error))
	return rv
}

// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/2880195-kernelwithfunctionname?language=objc
func Kernel_KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data []byte, format Format, error foundation.IError) Kernel {
	return KernelClass.KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name, data, format, error)
}

func (kc _KernelClass) Alloc() Kernel {
	rv := objc.Call[Kernel](kc, objc.Sel("alloc"))
	return rv
}

func Kernel_Alloc() Kernel {
	return KernelClass.Alloc()
}

func (kc _KernelClass) New() Kernel {
	rv := objc.Call[Kernel](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKernel() Kernel {
	return KernelClass.New()
}

func (k_ Kernel) Init() Kernel {
	rv := objc.Call[Kernel](k_, objc.Sel("init"))
	return rv
}

// Sets the selector Core Image uses to query the region of interest for image processing with the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/1437691-setroiselector?language=objc
func (k_ Kernel) SetROISelector(method objc.Selector) {
	objc.Call[objc.Void](k_, objc.Sel("setROISelector:"), method)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/3857565-kernelswithmetalstring?language=objc
func (kc _KernelClass) KernelsWithMetalStringError(source string, error foundation.IError) []Kernel {
	rv := objc.Call[[]Kernel](kc, objc.Sel("kernelsWithMetalString:error:"), source, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/3857565-kernelswithmetalstring?language=objc
func Kernel_KernelsWithMetalStringError(source string, error foundation.IError) []Kernel {
	return KernelClass.KernelsWithMetalStringError(source, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/3577535-kernelnamesfrommetallibrarydata?language=objc
func (kc _KernelClass) KernelNamesFromMetalLibraryData(data []byte) []string {
	rv := objc.Call[[]string](kc, objc.Sel("kernelNamesFromMetalLibraryData:"), data)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/3577535-kernelnamesfrommetallibrarydata?language=objc
func Kernel_KernelNamesFromMetalLibraryData(data []byte) []string {
	return KernelClass.KernelNamesFromMetalLibraryData(data)
}

// Creates a new image using the kernel and specified arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/1438243-applywithextent?language=objc
func (k_ Kernel) ApplyWithExtentRoiCallbackArguments(extent coregraphics.Rect, callback KernelROICallback, args []objc.IObject) Image {
	rv := objc.Call[Image](k_, objc.Sel("applyWithExtent:roiCallback:arguments:"), extent, callback, args)
	return rv
}

// The name of the kernel routine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikernel/1438067-name?language=objc
func (k_ Kernel) Name() string {
	rv := objc.Call[string](k_, objc.Sel("name"))
	return rv
}
