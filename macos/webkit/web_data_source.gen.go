// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebDataSource] class.
var WebDataSourceClass = _WebDataSourceClass{objc.GetClass("WebDataSource")}

type _WebDataSourceClass struct {
	objc.Class
}

// An interface definition for the [WebDataSource] class.
type IWebDataSource interface {
	objc.IObject
}

// WebDataSource encapsulates the web content to be displayed in a web frame view. A WebDataSource object has a representation object, conforming to the WebDocumentRepresentation protocol, that holds the data in an appropriate format depending on the MIME type. You can extend WebKit to support new MIME types by implementing your own view and representation classes, and specifying the mapping between them using the  registerViewClass:representationClass:forMIMEType: WebView class method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdatasource?language=objc
type WebDataSource struct {
	objc.Object
}

func WebDataSourceFrom(ptr unsafe.Pointer) WebDataSource {
	return WebDataSource{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebDataSourceClass) Alloc() WebDataSource {
	rv := objc.Call[WebDataSource](wc, objc.Sel("alloc"))
	return rv
}

func WebDataSource_Alloc() WebDataSource {
	return WebDataSourceClass.Alloc()
}

func (wc _WebDataSourceClass) New() WebDataSource {
	rv := objc.Call[WebDataSource](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebDataSource() WebDataSource {
	return WebDataSourceClass.New()
}

func (w_ WebDataSource) Init() WebDataSource {
	rv := objc.Call[WebDataSource](w_, objc.Sel("init"))
	return rv
}
