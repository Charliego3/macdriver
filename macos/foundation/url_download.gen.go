// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLDownload] class.
var URLDownloadClass = _URLDownloadClass{objc.GetClass("NSURLDownload")}

type _URLDownloadClass struct {
	objc.Class
}

// An interface definition for the [URLDownload] class.
type IURLDownload interface {
	objc.IObject
	SetDestinationAllowOverwrite(path string, allowOverwrite bool)
	Cancel()
	Request() URLRequest
	ResumeData() []byte
	DeletesFileUponFailure() bool
	SetDeletesFileUponFailure(value bool)
}

// An object that downloads a resource asynchronously and saves the data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload?language=objc
type URLDownload struct {
	objc.Object
}

func URLDownloadFrom(ptr unsafe.Pointer) URLDownload {
	return URLDownload{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLDownloadClass) Alloc() URLDownload {
	rv := objc.Call[URLDownload](uc, objc.Sel("alloc"))
	return rv
}

func URLDownload_Alloc() URLDownload {
	return URLDownloadClass.Alloc()
}

func (uc _URLDownloadClass) New() URLDownload {
	rv := objc.Call[URLDownload](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLDownload() URLDownload {
	return URLDownloadClass.New()
}

func (u_ URLDownload) Init() URLDownload {
	rv := objc.Call[URLDownload](u_, objc.Sel("init"))
	return rv
}

// Sets the destination path of the downloaded file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1412923-setdestination?language=objc
func (u_ URLDownload) SetDestinationAllowOverwrite(path string, allowOverwrite bool) {
	objc.Call[objc.Void](u_, objc.Sel("setDestination:allowOverwrite:"), path, allowOverwrite)
}

// Returns whether a URL download object can resume a download that was decoded with the specified MIME type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1409113-canresumedownloaddecodedwithenco?language=objc
func (uc _URLDownloadClass) CanResumeDownloadDecodedWithEncodingMIMEType(MIMEType string) bool {
	rv := objc.Call[bool](uc, objc.Sel("canResumeDownloadDecodedWithEncodingMIMEType:"), MIMEType)
	return rv
}

// Returns whether a URL download object can resume a download that was decoded with the specified MIME type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1409113-canresumedownloaddecodedwithenco?language=objc
func URLDownload_CanResumeDownloadDecodedWithEncodingMIMEType(MIMEType string) bool {
	return URLDownloadClass.CanResumeDownloadDecodedWithEncodingMIMEType(MIMEType)
}

// Cancels the receiver’s download and deletes the downloaded file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1415299-cancel?language=objc
func (u_ URLDownload) Cancel() {
	objc.Call[objc.Void](u_, objc.Sel("cancel"))
}

// Returns the request that initiated the receiver’s download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1416157-request?language=objc
func (u_ URLDownload) Request() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("request"))
	return rv
}

// Returns the resume data for a download that is not yet complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1413244-resumedata?language=objc
func (u_ URLDownload) ResumeData() []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("resumeData"))
	return rv
}

// Returns whether the receiver deletes partially downloaded files when a download stops prematurely. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1409172-deletesfileuponfailure?language=objc
func (u_ URLDownload) DeletesFileUponFailure() bool {
	rv := objc.Call[bool](u_, objc.Sel("deletesFileUponFailure"))
	return rv
}

// Returns whether the receiver deletes partially downloaded files when a download stops prematurely. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/1409172-deletesfileuponfailure?language=objc
func (u_ URLDownload) SetDeletesFileUponFailure(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setDeletesFileUponFailure:"), value)
}
