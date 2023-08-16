// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PDFInfo] class.
var PDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}

type _PDFInfoClass struct {
	objc.Class
}

// An interface definition for the [PDFInfo] class.
type IPDFInfo interface {
	objc.IObject
	TagNames() []string
	SetTagNames(value []string)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	IsFileExtensionHidden() bool
	SetFileExtensionHidden(value bool)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	Attributes() foundation.MutableDictionary
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
}

// An object that stores information associated with the creation of a PDF file, such as its URL, tag names, page orientation, and paper size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo?language=objc
type PDFInfo struct {
	objc.Object
}

func PDFInfoFrom(ptr unsafe.Pointer) PDFInfo {
	return PDFInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PDFInfoClass) Alloc() PDFInfo {
	rv := objc.Call[PDFInfo](pc, objc.Sel("alloc"))
	return rv
}

func PDFInfo_Alloc() PDFInfo {
	return PDFInfoClass.Alloc()
}

func (pc _PDFInfoClass) New() PDFInfo {
	rv := objc.Call[PDFInfo](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPDFInfo() PDFInfo {
	return PDFInfoClass.New()
}

func (p_ PDFInfo) Init() PDFInfo {
	rv := objc.Call[PDFInfo](p_, objc.Sel("init"))
	return rv
}

// An array of tag names that should be applied to the PDF file after it’s created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1525418-tagnames?language=objc
func (p_ PDFInfo) TagNames() []string {
	rv := objc.Call[[]string](p_, objc.Sel("tagNames"))
	return rv
}

// An array of tag names that should be applied to the PDF file after it’s created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1525418-tagnames?language=objc
func (p_ PDFInfo) SetTagNames(value []string) {
	objc.Call[objc.Void](p_, objc.Sel("setTagNames:"), value)
}

// The paper size to use when exporting content as a PDF file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1532272-papersize?language=objc
func (p_ PDFInfo) PaperSize() foundation.Size {
	rv := objc.Call[foundation.Size](p_, objc.Sel("paperSize"))
	return rv
}

// The paper size to use when exporting content as a PDF file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1532272-papersize?language=objc
func (p_ PDFInfo) SetPaperSize(value foundation.Size) {
	objc.Call[objc.Void](p_, objc.Sel("setPaperSize:"), value)
}

// A Boolean value that indicates whether the file extension should appear after the filename. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1527208-fileextensionhidden?language=objc
func (p_ PDFInfo) IsFileExtensionHidden() bool {
	rv := objc.Call[bool](p_, objc.Sel("isFileExtensionHidden"))
	return rv
}

// A Boolean value that indicates whether the file extension should appear after the filename. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1527208-fileextensionhidden?language=objc
func (p_ PDFInfo) SetFileExtensionHidden(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setFileExtensionHidden:"), value)
}

// The URL identifying the location at which the PDF file will be created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1528885-url?language=objc
func (p_ PDFInfo) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// The URL identifying the location at which the PDF file will be created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1528885-url?language=objc
func (p_ PDFInfo) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), objc.Ptr(value))
}

// A dictionary of additional attributes that describe how to export content as a PDF file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1528715-attributes?language=objc
func (p_ PDFInfo) Attributes() foundation.MutableDictionary {
	rv := objc.Call[foundation.MutableDictionary](p_, objc.Sel("attributes"))
	return rv
}

// The paper orientation to use when exporting content as a PDF file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1524848-orientation?language=objc
func (p_ PDFInfo) Orientation() PaperOrientation {
	rv := objc.Call[PaperOrientation](p_, objc.Sel("orientation"))
	return rv
}

// The paper orientation to use when exporting content as a PDF file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/1524848-orientation?language=objc
func (p_ PDFInfo) SetOrientation(value PaperOrientation) {
	objc.Call[objc.Void](p_, objc.Sel("setOrientation:"), value)
}
