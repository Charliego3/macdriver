// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSStyleDeclaration] class.
var DOMCSSStyleDeclarationClass = _DOMCSSStyleDeclarationClass{objc.GetClass("DOMCSSStyleDeclaration")}

type _DOMCSSStyleDeclarationClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSStyleDeclaration] class.
type IDOMCSSStyleDeclaration interface {
	IDOMObject
	SetBorderTop(borderTop string)
	SetZIndex(zIndex string)
	PauseBefore() string
	SetMinWidth(minWidth string)
	Content() string
	SetBorderRightStyle(borderRightStyle string)
	SetSpeechRate(speechRate string)
	Cursor() string
	SetCueBefore(cueBefore string)
	WhiteSpace() string
	Color() string
	SetLetterSpacing(letterSpacing string)
	FontSizeAdjust() string
	PlayDuring() string
	SetBackground(background string)
	Width() string
	SetPaddingTop(paddingTop string)
	SetAzimuth(azimuth string)
	PageBreakAfter() string
	PitchRange() string
	Direction() string
	Marks() string
	Elevation() string
	TextIndent() string
	SetFontFamily(fontFamily string)
	FontSize() string
	EmptyCells() string
	BorderBottomWidth() string
	Volume() string
	SetOverflow(overflow string)
	SetClip(clip string)
	UnicodeBidi() string
	PaddingBottom() string
	SetCssFloat(cssFloat string)
	BorderBottomColor() string
	VerticalAlign() string
	MinHeight() string
	SetTextShadow(textShadow string)
	SetPadding(padding string)
	SetBorderStyle(borderStyle string)
	BorderTopColor() string
	TextShadow() string
	SetEmptyCells(emptyCells string)
	SetCaptionSide(captionSide string)
	BorderTopWidth() string
	SetBorderBottomStyle(borderBottomStyle string)
	SetDisplay(display string)
	SetTextIndent(textIndent string)
	SetBorderCollapse(borderCollapse string)
	FontFamily() string
	SetBorder(border string)
	SetPitchRange(pitchRange string)
	BorderSpacing() string
	ListStyle() string
	TextTransform() string
	SetPitch(pitch string)
	SetSpeakHeader(speakHeader string)
	Background() string
	SetPlayDuring(playDuring string)
	PaddingTop() string
	SetPause(pause string)
	SetOrphans(orphans string)
	Clip() string
	SetWhiteSpace(whiteSpace string)
	SpeechRate() string
	SetFontStretch(fontStretch string)
	SetBorderLeftStyle(borderLeftStyle string)
	SetPosition(position string)
	Azimuth() string
	SetBorderBottomColor(borderBottomColor string)
	WordSpacing() string
	SetBorderBottomWidth(borderBottomWidth string)
	SetCue(cue string)
	MinWidth() string
	SetBorderColor(borderColor string)
	SetCounterIncrement(counterIncrement string)
	BorderTopStyle() string
	SetBorderWidth(borderWidth string)
	PageBreakInside() string
	Clear() string
	SetBackgroundImage(backgroundImage string)
	SetFontSize(fontSize string)
	SetBorderLeftColor(borderLeftColor string)
	SetBorderSpacing(borderSpacing string)
	SetTextDecoration(textDecoration string)
	BackgroundRepeat() string
	SetTextTransform(textTransform string)
	Left() string
	SetBorderLeftWidth(borderLeftWidth string)
	SetPageBreakBefore(pageBreakBefore string)
	SetContent(content string)
	VoiceFamily() string
	CssFloat() string
	Overflow() string
	SetListStyle(listStyle string)
	SetDirection(direction string)
	Cue() string
	Orphans() string
	SetPaddingBottom(paddingBottom string)
	SetRight(right string)
	SetVerticalAlign(verticalAlign string)
	SetStress(stress string)
	SetElevation(elevation string)
	BorderBottom() string
	SetSpeak(speak string)
	Height() string
	Quotes() string
	Margin() string
	SetBorderRightWidth(borderRightWidth string)
	SetSpeakPunctuation(speakPunctuation string)
	BorderRight() string
	SetBorderRightColor(borderRightColor string)
	SetMinHeight(minHeight string)
	SetFontVariant(fontVariant string)
	PaddingLeft() string
	Padding() string
	BorderTop() string
	SetBottom(bottom string)
	LetterSpacing() string
	TableLayout() string
	SetBorderBottom(borderBottom string)
	BorderBottomStyle() string
	MarginRight() string
	SetLeft(left string)
	Position() string
	ListStyleImage() string
	SetWidows(widows string)
	BackgroundColor() string
	Display() string
	CueBefore() string
	SetFontSizeAdjust(fontSizeAdjust string)
	SetFontWeight(fontWeight string)
	Outline() string
	Bottom() string
	SetRichness(richness string)
	SetOutlineWidth(outlineWidth string)
	SpeakPunctuation() string
	SetUnicodeBidi(unicodeBidi string)
	Widows() string
	BorderRightColor() string
	BackgroundAttachment() string
	SetPageBreakInside(pageBreakInside string)
	CounterReset() string
	SetOutlineColor(outlineColor string)
	BorderRightWidth() string
	SetFontStyle(fontStyle string)
	Font() string
	BackgroundImage() string
	BorderLeftColor() string
	SetVisibility(visibility string)
	SetPauseBefore(pauseBefore string)
	Stress() string
	SetClear(clear string)
	BorderLeftWidth() string
	SetBackgroundAttachment(backgroundAttachment string)
	PageBreakBefore() string
	SetHeight(height string)
	MaxHeight() string
	SetMargin(margin string)
	OutlineWidth() string
	SetQuotes(quotes string)
	ListStyleType() string
	SetPageBreakAfter(pageBreakAfter string)
	OutlineColor() string
	SetCounterReset(counterReset string)
	Top() string
	SpeakHeader() string
	SetListStylePosition(listStylePosition string)
	SetFont(font string)
	SetTextAlign(textAlign string)
	MaxWidth() string
	MarginTop() string
	SetSpeakNumeral(speakNumeral string)
	BorderCollapse() string
	FontStretch() string
	SetBorderTopWidth(borderTopWidth string)
	CueAfter() string
	Visibility() string
	SetBorderTopColor(borderTopColor string)
	CounterIncrement() string
	SetBackgroundRepeat(backgroundRepeat string)
	BorderStyle() string
	SpeakNumeral() string
	CaptionSide() string
	SetTop(top string)
	SetBackgroundColor(backgroundColor string)
	Speak() string
	FontWeight() string
	Right() string
	SetPage(page string)
	SetPaddingRight(paddingRight string)
	Richness() string
	SetMarginLeft(marginLeft string)
	SetBorderTopStyle(borderTopStyle string)
	SetVoiceFamily(voiceFamily string)
	LineHeight() string
	SetMarks(marks string)
	SetColor(color string)
	PaddingRight() string
	TextAlign() string
	SetOutline(outline string)
	SetWordSpacing(wordSpacing string)
	SetWidth(width string)
	SetMarginTop(marginTop string)
	BorderColor() string
	SetSize(size string)
	TextDecoration() string
	SetPauseAfter(pauseAfter string)
	Border() string
	BorderWidth() string
	SetBorderLeft(borderLeft string)
	BorderRightStyle() string
	BorderLeft() string
	PauseAfter() string
	SetListStyleType(listStyleType string)
	SetTableLayout(tableLayout string)
	SetOutlineStyle(outlineStyle string)
	SetMarginBottom(marginBottom string)
	SetMaxWidth(maxWidth string)
	SetMarginRight(marginRight string)
	Pause() string
	SetVolume(volume string)
	SetCueAfter(cueAfter string)
	SetMaxHeight(maxHeight string)
	MarkerOffset() string
	Page() string
	ZIndex() string
	Pitch() string
	SetListStyleImage(listStyleImage string)
	OutlineStyle() string
	Size() string
	MarginBottom() string
	SetLineHeight(lineHeight string)
	BackgroundPosition() string
	MarginLeft() string
	SetCursor(cursor string)
	BorderLeftStyle() string
	SetBackgroundPosition(backgroundPosition string)
	SetBorderRight(borderRight string)
	FontStyle() string
	ListStylePosition() string
	SetPaddingLeft(paddingLeft string)
	SetMarkerOffset(markerOffset string)
	FontVariant() string
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration?language=objc
type DOMCSSStyleDeclaration struct {
	DOMObject
}

func DOMCSSStyleDeclarationFrom(ptr unsafe.Pointer) DOMCSSStyleDeclaration {
	return DOMCSSStyleDeclaration{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMCSSStyleDeclarationClass) Alloc() DOMCSSStyleDeclaration {
	rv := objc.Call[DOMCSSStyleDeclaration](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSStyleDeclaration_Alloc() DOMCSSStyleDeclaration {
	return DOMCSSStyleDeclarationClass.Alloc()
}

func (dc _DOMCSSStyleDeclarationClass) New() DOMCSSStyleDeclaration {
	rv := objc.Call[DOMCSSStyleDeclaration](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSStyleDeclaration() DOMCSSStyleDeclaration {
	return DOMCSSStyleDeclarationClass.New()
}

func (d_ DOMCSSStyleDeclaration) Init() DOMCSSStyleDeclaration {
	rv := objc.Call[DOMCSSStyleDeclaration](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537037-setbordertop?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderTop(borderTop string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderTop:"), borderTop)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536811-setzindex?language=objc
func (d_ DOMCSSStyleDeclaration) SetZIndex(zIndex string) {
	objc.Call[objc.Void](d_, objc.Sel("setZIndex:"), zIndex)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537469-pausebefore?language=objc
func (d_ DOMCSSStyleDeclaration) PauseBefore() string {
	rv := objc.Call[string](d_, objc.Sel("pauseBefore"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537610-setminwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetMinWidth(minWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setMinWidth:"), minWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537719-content?language=objc
func (d_ DOMCSSStyleDeclaration) Content() string {
	rv := objc.Call[string](d_, objc.Sel("content"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537468-setborderrightstyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderRightStyle:"), borderRightStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537632-setspeechrate?language=objc
func (d_ DOMCSSStyleDeclaration) SetSpeechRate(speechRate string) {
	objc.Call[objc.Void](d_, objc.Sel("setSpeechRate:"), speechRate)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537597-cursor?language=objc
func (d_ DOMCSSStyleDeclaration) Cursor() string {
	rv := objc.Call[string](d_, objc.Sel("cursor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536519-setcuebefore?language=objc
func (d_ DOMCSSStyleDeclaration) SetCueBefore(cueBefore string) {
	objc.Call[objc.Void](d_, objc.Sel("setCueBefore:"), cueBefore)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537569-whitespace?language=objc
func (d_ DOMCSSStyleDeclaration) WhiteSpace() string {
	rv := objc.Call[string](d_, objc.Sel("whiteSpace"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536621-color?language=objc
func (d_ DOMCSSStyleDeclaration) Color() string {
	rv := objc.Call[string](d_, objc.Sel("color"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537534-setletterspacing?language=objc
func (d_ DOMCSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	objc.Call[objc.Void](d_, objc.Sel("setLetterSpacing:"), letterSpacing)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537039-fontsizeadjust?language=objc
func (d_ DOMCSSStyleDeclaration) FontSizeAdjust() string {
	rv := objc.Call[string](d_, objc.Sel("fontSizeAdjust"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537957-playduring?language=objc
func (d_ DOMCSSStyleDeclaration) PlayDuring() string {
	rv := objc.Call[string](d_, objc.Sel("playDuring"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537919-setbackground?language=objc
func (d_ DOMCSSStyleDeclaration) SetBackground(background string) {
	objc.Call[objc.Void](d_, objc.Sel("setBackground:"), background)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536411-width?language=objc
func (d_ DOMCSSStyleDeclaration) Width() string {
	rv := objc.Call[string](d_, objc.Sel("width"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537395-setpaddingtop?language=objc
func (d_ DOMCSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingTop:"), paddingTop)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536701-setazimuth?language=objc
func (d_ DOMCSSStyleDeclaration) SetAzimuth(azimuth string) {
	objc.Call[objc.Void](d_, objc.Sel("setAzimuth:"), azimuth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537432-pagebreakafter?language=objc
func (d_ DOMCSSStyleDeclaration) PageBreakAfter() string {
	rv := objc.Call[string](d_, objc.Sel("pageBreakAfter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537451-pitchrange?language=objc
func (d_ DOMCSSStyleDeclaration) PitchRange() string {
	rv := objc.Call[string](d_, objc.Sel("pitchRange"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537340-direction?language=objc
func (d_ DOMCSSStyleDeclaration) Direction() string {
	rv := objc.Call[string](d_, objc.Sel("direction"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536295-marks?language=objc
func (d_ DOMCSSStyleDeclaration) Marks() string {
	rv := objc.Call[string](d_, objc.Sel("marks"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537097-elevation?language=objc
func (d_ DOMCSSStyleDeclaration) Elevation() string {
	rv := objc.Call[string](d_, objc.Sel("elevation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537671-textindent?language=objc
func (d_ DOMCSSStyleDeclaration) TextIndent() string {
	rv := objc.Call[string](d_, objc.Sel("textIndent"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537485-setfontfamily?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontFamily(fontFamily string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontFamily:"), fontFamily)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538175-fontsize?language=objc
func (d_ DOMCSSStyleDeclaration) FontSize() string {
	rv := objc.Call[string](d_, objc.Sel("fontSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536606-emptycells?language=objc
func (d_ DOMCSSStyleDeclaration) EmptyCells() string {
	rv := objc.Call[string](d_, objc.Sel("emptyCells"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537966-borderbottomwidth?language=objc
func (d_ DOMCSSStyleDeclaration) BorderBottomWidth() string {
	rv := objc.Call[string](d_, objc.Sel("borderBottomWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538123-volume?language=objc
func (d_ DOMCSSStyleDeclaration) Volume() string {
	rv := objc.Call[string](d_, objc.Sel("volume"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537548-setoverflow?language=objc
func (d_ DOMCSSStyleDeclaration) SetOverflow(overflow string) {
	objc.Call[objc.Void](d_, objc.Sel("setOverflow:"), overflow)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536697-setclip?language=objc
func (d_ DOMCSSStyleDeclaration) SetClip(clip string) {
	objc.Call[objc.Void](d_, objc.Sel("setClip:"), clip)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537505-unicodebidi?language=objc
func (d_ DOMCSSStyleDeclaration) UnicodeBidi() string {
	rv := objc.Call[string](d_, objc.Sel("unicodeBidi"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538125-paddingbottom?language=objc
func (d_ DOMCSSStyleDeclaration) PaddingBottom() string {
	rv := objc.Call[string](d_, objc.Sel("paddingBottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538076-setcssfloat?language=objc
func (d_ DOMCSSStyleDeclaration) SetCssFloat(cssFloat string) {
	objc.Call[objc.Void](d_, objc.Sel("setCssFloat:"), cssFloat)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536522-borderbottomcolor?language=objc
func (d_ DOMCSSStyleDeclaration) BorderBottomColor() string {
	rv := objc.Call[string](d_, objc.Sel("borderBottomColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536515-verticalalign?language=objc
func (d_ DOMCSSStyleDeclaration) VerticalAlign() string {
	rv := objc.Call[string](d_, objc.Sel("verticalAlign"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537968-minheight?language=objc
func (d_ DOMCSSStyleDeclaration) MinHeight() string {
	rv := objc.Call[string](d_, objc.Sel("minHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537435-settextshadow?language=objc
func (d_ DOMCSSStyleDeclaration) SetTextShadow(textShadow string) {
	objc.Call[objc.Void](d_, objc.Sel("setTextShadow:"), textShadow)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537535-setpadding?language=objc
func (d_ DOMCSSStyleDeclaration) SetPadding(padding string) {
	objc.Call[objc.Void](d_, objc.Sel("setPadding:"), padding)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536214-setborderstyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderStyle:"), borderStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537557-bordertopcolor?language=objc
func (d_ DOMCSSStyleDeclaration) BorderTopColor() string {
	rv := objc.Call[string](d_, objc.Sel("borderTopColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537426-textshadow?language=objc
func (d_ DOMCSSStyleDeclaration) TextShadow() string {
	rv := objc.Call[string](d_, objc.Sel("textShadow"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536195-setemptycells?language=objc
func (d_ DOMCSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	objc.Call[objc.Void](d_, objc.Sel("setEmptyCells:"), emptyCells)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537601-setcaptionside?language=objc
func (d_ DOMCSSStyleDeclaration) SetCaptionSide(captionSide string) {
	objc.Call[objc.Void](d_, objc.Sel("setCaptionSide:"), captionSide)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536449-bordertopwidth?language=objc
func (d_ DOMCSSStyleDeclaration) BorderTopWidth() string {
	rv := objc.Call[string](d_, objc.Sel("borderTopWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537769-setborderbottomstyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderBottomStyle:"), borderBottomStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537000-setdisplay?language=objc
func (d_ DOMCSSStyleDeclaration) SetDisplay(display string) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplay:"), display)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536327-settextindent?language=objc
func (d_ DOMCSSStyleDeclaration) SetTextIndent(textIndent string) {
	objc.Call[objc.Void](d_, objc.Sel("setTextIndent:"), textIndent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537486-setbordercollapse?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderCollapse:"), borderCollapse)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537588-fontfamily?language=objc
func (d_ DOMCSSStyleDeclaration) FontFamily() string {
	rv := objc.Call[string](d_, objc.Sel("fontFamily"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538163-setborder?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorder(border string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorder:"), border)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537383-setpitchrange?language=objc
func (d_ DOMCSSStyleDeclaration) SetPitchRange(pitchRange string) {
	objc.Call[objc.Void](d_, objc.Sel("setPitchRange:"), pitchRange)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536488-borderspacing?language=objc
func (d_ DOMCSSStyleDeclaration) BorderSpacing() string {
	rv := objc.Call[string](d_, objc.Sel("borderSpacing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536240-liststyle?language=objc
func (d_ DOMCSSStyleDeclaration) ListStyle() string {
	rv := objc.Call[string](d_, objc.Sel("listStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537403-texttransform?language=objc
func (d_ DOMCSSStyleDeclaration) TextTransform() string {
	rv := objc.Call[string](d_, objc.Sel("textTransform"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536246-setpitch?language=objc
func (d_ DOMCSSStyleDeclaration) SetPitch(pitch string) {
	objc.Call[objc.Void](d_, objc.Sel("setPitch:"), pitch)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536545-setspeakheader?language=objc
func (d_ DOMCSSStyleDeclaration) SetSpeakHeader(speakHeader string) {
	objc.Call[objc.Void](d_, objc.Sel("setSpeakHeader:"), speakHeader)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537291-background?language=objc
func (d_ DOMCSSStyleDeclaration) Background() string {
	rv := objc.Call[string](d_, objc.Sel("background"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537715-setplayduring?language=objc
func (d_ DOMCSSStyleDeclaration) SetPlayDuring(playDuring string) {
	objc.Call[objc.Void](d_, objc.Sel("setPlayDuring:"), playDuring)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537698-paddingtop?language=objc
func (d_ DOMCSSStyleDeclaration) PaddingTop() string {
	rv := objc.Call[string](d_, objc.Sel("paddingTop"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536470-setpause?language=objc
func (d_ DOMCSSStyleDeclaration) SetPause(pause string) {
	objc.Call[objc.Void](d_, objc.Sel("setPause:"), pause)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537541-setorphans?language=objc
func (d_ DOMCSSStyleDeclaration) SetOrphans(orphans string) {
	objc.Call[objc.Void](d_, objc.Sel("setOrphans:"), orphans)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536966-clip?language=objc
func (d_ DOMCSSStyleDeclaration) Clip() string {
	rv := objc.Call[string](d_, objc.Sel("clip"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536604-setwhitespace?language=objc
func (d_ DOMCSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	objc.Call[objc.Void](d_, objc.Sel("setWhiteSpace:"), whiteSpace)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537016-speechrate?language=objc
func (d_ DOMCSSStyleDeclaration) SpeechRate() string {
	rv := objc.Call[string](d_, objc.Sel("speechRate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537280-setfontstretch?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontStretch(fontStretch string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontStretch:"), fontStretch)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537982-setborderleftstyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderLeftStyle:"), borderLeftStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537303-setposition?language=objc
func (d_ DOMCSSStyleDeclaration) SetPosition(position string) {
	objc.Call[objc.Void](d_, objc.Sel("setPosition:"), position)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537087-azimuth?language=objc
func (d_ DOMCSSStyleDeclaration) Azimuth() string {
	rv := objc.Call[string](d_, objc.Sel("azimuth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537304-setborderbottomcolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderBottomColor:"), borderBottomColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536882-wordspacing?language=objc
func (d_ DOMCSSStyleDeclaration) WordSpacing() string {
	rv := objc.Call[string](d_, objc.Sel("wordSpacing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536783-setborderbottomwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderBottomWidth:"), borderBottomWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537627-setcue?language=objc
func (d_ DOMCSSStyleDeclaration) SetCue(cue string) {
	objc.Call[objc.Void](d_, objc.Sel("setCue:"), cue)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537387-minwidth?language=objc
func (d_ DOMCSSStyleDeclaration) MinWidth() string {
	rv := objc.Call[string](d_, objc.Sel("minWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537641-setbordercolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderColor(borderColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderColor:"), borderColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537174-setcounterincrement?language=objc
func (d_ DOMCSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	objc.Call[objc.Void](d_, objc.Sel("setCounterIncrement:"), counterIncrement)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536636-bordertopstyle?language=objc
func (d_ DOMCSSStyleDeclaration) BorderTopStyle() string {
	rv := objc.Call[string](d_, objc.Sel("borderTopStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536431-setborderwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderWidth:"), borderWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538055-pagebreakinside?language=objc
func (d_ DOMCSSStyleDeclaration) PageBreakInside() string {
	rv := objc.Call[string](d_, objc.Sel("pageBreakInside"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536731-clear?language=objc
func (d_ DOMCSSStyleDeclaration) Clear() string {
	rv := objc.Call[string](d_, objc.Sel("clear"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538180-setbackgroundimage?language=objc
func (d_ DOMCSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundImage:"), backgroundImage)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536339-setfontsize?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontSize(fontSize string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontSize:"), fontSize)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537964-setborderleftcolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderLeftColor:"), borderLeftColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537743-setborderspacing?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderSpacing:"), borderSpacing)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537524-settextdecoration?language=objc
func (d_ DOMCSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	objc.Call[objc.Void](d_, objc.Sel("setTextDecoration:"), textDecoration)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537415-backgroundrepeat?language=objc
func (d_ DOMCSSStyleDeclaration) BackgroundRepeat() string {
	rv := objc.Call[string](d_, objc.Sel("backgroundRepeat"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537484-settexttransform?language=objc
func (d_ DOMCSSStyleDeclaration) SetTextTransform(textTransform string) {
	objc.Call[objc.Void](d_, objc.Sel("setTextTransform:"), textTransform)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537622-left?language=objc
func (d_ DOMCSSStyleDeclaration) Left() string {
	rv := objc.Call[string](d_, objc.Sel("left"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537422-setborderleftwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderLeftWidth:"), borderLeftWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537929-setpagebreakbefore?language=objc
func (d_ DOMCSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	objc.Call[objc.Void](d_, objc.Sel("setPageBreakBefore:"), pageBreakBefore)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537673-setcontent?language=objc
func (d_ DOMCSSStyleDeclaration) SetContent(content string) {
	objc.Call[objc.Void](d_, objc.Sel("setContent:"), content)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537517-voicefamily?language=objc
func (d_ DOMCSSStyleDeclaration) VoiceFamily() string {
	rv := objc.Call[string](d_, objc.Sel("voiceFamily"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537488-cssfloat?language=objc
func (d_ DOMCSSStyleDeclaration) CssFloat() string {
	rv := objc.Call[string](d_, objc.Sel("cssFloat"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537327-overflow?language=objc
func (d_ DOMCSSStyleDeclaration) Overflow() string {
	rv := objc.Call[string](d_, objc.Sel("overflow"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537434-setliststyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetListStyle(listStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setListStyle:"), listStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536956-setdirection?language=objc
func (d_ DOMCSSStyleDeclaration) SetDirection(direction string) {
	objc.Call[objc.Void](d_, objc.Sel("setDirection:"), direction)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536215-cue?language=objc
func (d_ DOMCSSStyleDeclaration) Cue() string {
	rv := objc.Call[string](d_, objc.Sel("cue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536200-orphans?language=objc
func (d_ DOMCSSStyleDeclaration) Orphans() string {
	rv := objc.Call[string](d_, objc.Sel("orphans"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537428-setpaddingbottom?language=objc
func (d_ DOMCSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingBottom:"), paddingBottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536389-setright?language=objc
func (d_ DOMCSSStyleDeclaration) SetRight(right string) {
	objc.Call[objc.Void](d_, objc.Sel("setRight:"), right)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537810-setverticalalign?language=objc
func (d_ DOMCSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	objc.Call[objc.Void](d_, objc.Sel("setVerticalAlign:"), verticalAlign)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536294-setstress?language=objc
func (d_ DOMCSSStyleDeclaration) SetStress(stress string) {
	objc.Call[objc.Void](d_, objc.Sel("setStress:"), stress)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536650-setelevation?language=objc
func (d_ DOMCSSStyleDeclaration) SetElevation(elevation string) {
	objc.Call[objc.Void](d_, objc.Sel("setElevation:"), elevation)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537136-borderbottom?language=objc
func (d_ DOMCSSStyleDeclaration) BorderBottom() string {
	rv := objc.Call[string](d_, objc.Sel("borderBottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536305-setspeak?language=objc
func (d_ DOMCSSStyleDeclaration) SetSpeak(speak string) {
	objc.Call[objc.Void](d_, objc.Sel("setSpeak:"), speak)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536925-height?language=objc
func (d_ DOMCSSStyleDeclaration) Height() string {
	rv := objc.Call[string](d_, objc.Sel("height"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537439-quotes?language=objc
func (d_ DOMCSSStyleDeclaration) Quotes() string {
	rv := objc.Call[string](d_, objc.Sel("quotes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538039-margin?language=objc
func (d_ DOMCSSStyleDeclaration) Margin() string {
	rv := objc.Call[string](d_, objc.Sel("margin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537779-setborderrightwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderRightWidth:"), borderRightWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536570-setspeakpunctuation?language=objc
func (d_ DOMCSSStyleDeclaration) SetSpeakPunctuation(speakPunctuation string) {
	objc.Call[objc.Void](d_, objc.Sel("setSpeakPunctuation:"), speakPunctuation)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537216-borderright?language=objc
func (d_ DOMCSSStyleDeclaration) BorderRight() string {
	rv := objc.Call[string](d_, objc.Sel("borderRight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536661-setborderrightcolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderRightColor:"), borderRightColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536334-setminheight?language=objc
func (d_ DOMCSSStyleDeclaration) SetMinHeight(minHeight string) {
	objc.Call[objc.Void](d_, objc.Sel("setMinHeight:"), minHeight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537586-setfontvariant?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontVariant(fontVariant string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontVariant:"), fontVariant)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537561-paddingleft?language=objc
func (d_ DOMCSSStyleDeclaration) PaddingLeft() string {
	rv := objc.Call[string](d_, objc.Sel("paddingLeft"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537208-padding?language=objc
func (d_ DOMCSSStyleDeclaration) Padding() string {
	rv := objc.Call[string](d_, objc.Sel("padding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537619-bordertop?language=objc
func (d_ DOMCSSStyleDeclaration) BorderTop() string {
	rv := objc.Call[string](d_, objc.Sel("borderTop"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536531-setbottom?language=objc
func (d_ DOMCSSStyleDeclaration) SetBottom(bottom string) {
	objc.Call[objc.Void](d_, objc.Sel("setBottom:"), bottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538024-letterspacing?language=objc
func (d_ DOMCSSStyleDeclaration) LetterSpacing() string {
	rv := objc.Call[string](d_, objc.Sel("letterSpacing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536287-tablelayout?language=objc
func (d_ DOMCSSStyleDeclaration) TableLayout() string {
	rv := objc.Call[string](d_, objc.Sel("tableLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536978-setborderbottom?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderBottom:"), borderBottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536865-borderbottomstyle?language=objc
func (d_ DOMCSSStyleDeclaration) BorderBottomStyle() string {
	rv := objc.Call[string](d_, objc.Sel("borderBottomStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537806-marginright?language=objc
func (d_ DOMCSSStyleDeclaration) MarginRight() string {
	rv := objc.Call[string](d_, objc.Sel("marginRight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536916-setleft?language=objc
func (d_ DOMCSSStyleDeclaration) SetLeft(left string) {
	objc.Call[objc.Void](d_, objc.Sel("setLeft:"), left)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537949-position?language=objc
func (d_ DOMCSSStyleDeclaration) Position() string {
	rv := objc.Call[string](d_, objc.Sel("position"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537725-liststyleimage?language=objc
func (d_ DOMCSSStyleDeclaration) ListStyleImage() string {
	rv := objc.Call[string](d_, objc.Sel("listStyleImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537307-setwidows?language=objc
func (d_ DOMCSSStyleDeclaration) SetWidows(widows string) {
	objc.Call[objc.Void](d_, objc.Sel("setWidows:"), widows)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536508-backgroundcolor?language=objc
func (d_ DOMCSSStyleDeclaration) BackgroundColor() string {
	rv := objc.Call[string](d_, objc.Sel("backgroundColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536272-display?language=objc
func (d_ DOMCSSStyleDeclaration) Display() string {
	rv := objc.Call[string](d_, objc.Sel("display"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537217-cuebefore?language=objc
func (d_ DOMCSSStyleDeclaration) CueBefore() string {
	rv := objc.Call[string](d_, objc.Sel("cueBefore"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536615-setfontsizeadjust?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontSizeAdjust:"), fontSizeAdjust)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536632-setfontweight?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontWeight(fontWeight string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontWeight:"), fontWeight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537652-outline?language=objc
func (d_ DOMCSSStyleDeclaration) Outline() string {
	rv := objc.Call[string](d_, objc.Sel("outline"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536221-bottom?language=objc
func (d_ DOMCSSStyleDeclaration) Bottom() string {
	rv := objc.Call[string](d_, objc.Sel("bottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537069-setrichness?language=objc
func (d_ DOMCSSStyleDeclaration) SetRichness(richness string) {
	objc.Call[objc.Void](d_, objc.Sel("setRichness:"), richness)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536742-setoutlinewidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setOutlineWidth:"), outlineWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537181-speakpunctuation?language=objc
func (d_ DOMCSSStyleDeclaration) SpeakPunctuation() string {
	rv := objc.Call[string](d_, objc.Sel("speakPunctuation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536789-setunicodebidi?language=objc
func (d_ DOMCSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	objc.Call[objc.Void](d_, objc.Sel("setUnicodeBidi:"), unicodeBidi)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536961-widows?language=objc
func (d_ DOMCSSStyleDeclaration) Widows() string {
	rv := objc.Call[string](d_, objc.Sel("widows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536630-borderrightcolor?language=objc
func (d_ DOMCSSStyleDeclaration) BorderRightColor() string {
	rv := objc.Call[string](d_, objc.Sel("borderRightColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537399-backgroundattachment?language=objc
func (d_ DOMCSSStyleDeclaration) BackgroundAttachment() string {
	rv := objc.Call[string](d_, objc.Sel("backgroundAttachment"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537608-setpagebreakinside?language=objc
func (d_ DOMCSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	objc.Call[objc.Void](d_, objc.Sel("setPageBreakInside:"), pageBreakInside)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537063-counterreset?language=objc
func (d_ DOMCSSStyleDeclaration) CounterReset() string {
	rv := objc.Call[string](d_, objc.Sel("counterReset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537412-setoutlinecolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setOutlineColor:"), outlineColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536486-borderrightwidth?language=objc
func (d_ DOMCSSStyleDeclaration) BorderRightWidth() string {
	rv := objc.Call[string](d_, objc.Sel("borderRightWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536843-setfontstyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetFontStyle(fontStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setFontStyle:"), fontStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536196-font?language=objc
func (d_ DOMCSSStyleDeclaration) Font() string {
	rv := objc.Call[string](d_, objc.Sel("font"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537747-backgroundimage?language=objc
func (d_ DOMCSSStyleDeclaration) BackgroundImage() string {
	rv := objc.Call[string](d_, objc.Sel("backgroundImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537504-borderleftcolor?language=objc
func (d_ DOMCSSStyleDeclaration) BorderLeftColor() string {
	rv := objc.Call[string](d_, objc.Sel("borderLeftColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537057-setvisibility?language=objc
func (d_ DOMCSSStyleDeclaration) SetVisibility(visibility string) {
	objc.Call[objc.Void](d_, objc.Sel("setVisibility:"), visibility)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538116-setpausebefore?language=objc
func (d_ DOMCSSStyleDeclaration) SetPauseBefore(pauseBefore string) {
	objc.Call[objc.Void](d_, objc.Sel("setPauseBefore:"), pauseBefore)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537394-stress?language=objc
func (d_ DOMCSSStyleDeclaration) Stress() string {
	rv := objc.Call[string](d_, objc.Sel("stress"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538129-setclear?language=objc
func (d_ DOMCSSStyleDeclaration) SetClear(clear string) {
	objc.Call[objc.Void](d_, objc.Sel("setClear:"), clear)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537527-borderleftwidth?language=objc
func (d_ DOMCSSStyleDeclaration) BorderLeftWidth() string {
	rv := objc.Call[string](d_, objc.Sel("borderLeftWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537695-setbackgroundattachment?language=objc
func (d_ DOMCSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundAttachment:"), backgroundAttachment)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537033-pagebreakbefore?language=objc
func (d_ DOMCSSStyleDeclaration) PageBreakBefore() string {
	rv := objc.Call[string](d_, objc.Sel("pageBreakBefore"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537760-setheight?language=objc
func (d_ DOMCSSStyleDeclaration) SetHeight(height string) {
	objc.Call[objc.Void](d_, objc.Sel("setHeight:"), height)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537059-maxheight?language=objc
func (d_ DOMCSSStyleDeclaration) MaxHeight() string {
	rv := objc.Call[string](d_, objc.Sel("maxHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537501-setmargin?language=objc
func (d_ DOMCSSStyleDeclaration) SetMargin(margin string) {
	objc.Call[objc.Void](d_, objc.Sel("setMargin:"), margin)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537819-outlinewidth?language=objc
func (d_ DOMCSSStyleDeclaration) OutlineWidth() string {
	rv := objc.Call[string](d_, objc.Sel("outlineWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536734-setquotes?language=objc
func (d_ DOMCSSStyleDeclaration) SetQuotes(quotes string) {
	objc.Call[objc.Void](d_, objc.Sel("setQuotes:"), quotes)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537498-liststyletype?language=objc
func (d_ DOMCSSStyleDeclaration) ListStyleType() string {
	rv := objc.Call[string](d_, objc.Sel("listStyleType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538004-setpagebreakafter?language=objc
func (d_ DOMCSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	objc.Call[objc.Void](d_, objc.Sel("setPageBreakAfter:"), pageBreakAfter)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536817-outlinecolor?language=objc
func (d_ DOMCSSStyleDeclaration) OutlineColor() string {
	rv := objc.Call[string](d_, objc.Sel("outlineColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536331-setcounterreset?language=objc
func (d_ DOMCSSStyleDeclaration) SetCounterReset(counterReset string) {
	objc.Call[objc.Void](d_, objc.Sel("setCounterReset:"), counterReset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537992-top?language=objc
func (d_ DOMCSSStyleDeclaration) Top() string {
	rv := objc.Call[string](d_, objc.Sel("top"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536788-speakheader?language=objc
func (d_ DOMCSSStyleDeclaration) SpeakHeader() string {
	rv := objc.Call[string](d_, objc.Sel("speakHeader"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537407-setliststyleposition?language=objc
func (d_ DOMCSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	objc.Call[objc.Void](d_, objc.Sel("setListStylePosition:"), listStylePosition)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536461-setfont?language=objc
func (d_ DOMCSSStyleDeclaration) SetFont(font string) {
	objc.Call[objc.Void](d_, objc.Sel("setFont:"), font)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536564-settextalign?language=objc
func (d_ DOMCSSStyleDeclaration) SetTextAlign(textAlign string) {
	objc.Call[objc.Void](d_, objc.Sel("setTextAlign:"), textAlign)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537577-maxwidth?language=objc
func (d_ DOMCSSStyleDeclaration) MaxWidth() string {
	rv := objc.Call[string](d_, objc.Sel("maxWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536405-margintop?language=objc
func (d_ DOMCSSStyleDeclaration) MarginTop() string {
	rv := objc.Call[string](d_, objc.Sel("marginTop"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536395-setspeaknumeral?language=objc
func (d_ DOMCSSStyleDeclaration) SetSpeakNumeral(speakNumeral string) {
	objc.Call[objc.Void](d_, objc.Sel("setSpeakNumeral:"), speakNumeral)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537897-bordercollapse?language=objc
func (d_ DOMCSSStyleDeclaration) BorderCollapse() string {
	rv := objc.Call[string](d_, objc.Sel("borderCollapse"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537452-fontstretch?language=objc
func (d_ DOMCSSStyleDeclaration) FontStretch() string {
	rv := objc.Call[string](d_, objc.Sel("fontStretch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536964-setbordertopwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderTopWidth:"), borderTopWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536557-cueafter?language=objc
func (d_ DOMCSSStyleDeclaration) CueAfter() string {
	rv := objc.Call[string](d_, objc.Sel("cueAfter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537329-visibility?language=objc
func (d_ DOMCSSStyleDeclaration) Visibility() string {
	rv := objc.Call[string](d_, objc.Sel("visibility"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537536-setbordertopcolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderTopColor:"), borderTopColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537595-counterincrement?language=objc
func (d_ DOMCSSStyleDeclaration) CounterIncrement() string {
	rv := objc.Call[string](d_, objc.Sel("counterIncrement"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536502-setbackgroundrepeat?language=objc
func (d_ DOMCSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundRepeat:"), backgroundRepeat)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537384-borderstyle?language=objc
func (d_ DOMCSSStyleDeclaration) BorderStyle() string {
	rv := objc.Call[string](d_, objc.Sel("borderStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536347-speaknumeral?language=objc
func (d_ DOMCSSStyleDeclaration) SpeakNumeral() string {
	rv := objc.Call[string](d_, objc.Sel("speakNumeral"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536185-captionside?language=objc
func (d_ DOMCSSStyleDeclaration) CaptionSide() string {
	rv := objc.Call[string](d_, objc.Sel("captionSide"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537420-settop?language=objc
func (d_ DOMCSSStyleDeclaration) SetTop(top string) {
	objc.Call[objc.Void](d_, objc.Sel("setTop:"), top)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536820-setbackgroundcolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundColor:"), backgroundColor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537827-speak?language=objc
func (d_ DOMCSSStyleDeclaration) Speak() string {
	rv := objc.Call[string](d_, objc.Sel("speak"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537495-fontweight?language=objc
func (d_ DOMCSSStyleDeclaration) FontWeight() string {
	rv := objc.Call[string](d_, objc.Sel("fontWeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536909-right?language=objc
func (d_ DOMCSSStyleDeclaration) Right() string {
	rv := objc.Call[string](d_, objc.Sel("right"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536179-setpage?language=objc
func (d_ DOMCSSStyleDeclaration) SetPage(page string) {
	objc.Call[objc.Void](d_, objc.Sel("setPage:"), page)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537083-setpaddingright?language=objc
func (d_ DOMCSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingRight:"), paddingRight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537821-richness?language=objc
func (d_ DOMCSSStyleDeclaration) Richness() string {
	rv := objc.Call[string](d_, objc.Sel("richness"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536467-setmarginleft?language=objc
func (d_ DOMCSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	objc.Call[objc.Void](d_, objc.Sel("setMarginLeft:"), marginLeft)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537004-setbordertopstyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderTopStyle:"), borderTopStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537862-setvoicefamily?language=objc
func (d_ DOMCSSStyleDeclaration) SetVoiceFamily(voiceFamily string) {
	objc.Call[objc.Void](d_, objc.Sel("setVoiceFamily:"), voiceFamily)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536560-lineheight?language=objc
func (d_ DOMCSSStyleDeclaration) LineHeight() string {
	rv := objc.Call[string](d_, objc.Sel("lineHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538135-setmarks?language=objc
func (d_ DOMCSSStyleDeclaration) SetMarks(marks string) {
	objc.Call[objc.Void](d_, objc.Sel("setMarks:"), marks)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537213-setcolor?language=objc
func (d_ DOMCSSStyleDeclaration) SetColor(color string) {
	objc.Call[objc.Void](d_, objc.Sel("setColor:"), color)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536223-paddingright?language=objc
func (d_ DOMCSSStyleDeclaration) PaddingRight() string {
	rv := objc.Call[string](d_, objc.Sel("paddingRight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536370-textalign?language=objc
func (d_ DOMCSSStyleDeclaration) TextAlign() string {
	rv := objc.Call[string](d_, objc.Sel("textAlign"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537237-setoutline?language=objc
func (d_ DOMCSSStyleDeclaration) SetOutline(outline string) {
	objc.Call[objc.Void](d_, objc.Sel("setOutline:"), outline)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537530-setwordspacing?language=objc
func (d_ DOMCSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	objc.Call[objc.Void](d_, objc.Sel("setWordSpacing:"), wordSpacing)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537062-setwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetWidth(width string) {
	objc.Call[objc.Void](d_, objc.Sel("setWidth:"), width)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537184-setmargintop?language=objc
func (d_ DOMCSSStyleDeclaration) SetMarginTop(marginTop string) {
	objc.Call[objc.Void](d_, objc.Sel("setMarginTop:"), marginTop)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537814-bordercolor?language=objc
func (d_ DOMCSSStyleDeclaration) BorderColor() string {
	rv := objc.Call[string](d_, objc.Sel("borderColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536231-setsize?language=objc
func (d_ DOMCSSStyleDeclaration) SetSize(size string) {
	objc.Call[objc.Void](d_, objc.Sel("setSize:"), size)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537483-textdecoration?language=objc
func (d_ DOMCSSStyleDeclaration) TextDecoration() string {
	rv := objc.Call[string](d_, objc.Sel("textDecoration"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537591-setpauseafter?language=objc
func (d_ DOMCSSStyleDeclaration) SetPauseAfter(pauseAfter string) {
	objc.Call[objc.Void](d_, objc.Sel("setPauseAfter:"), pauseAfter)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537457-border?language=objc
func (d_ DOMCSSStyleDeclaration) Border() string {
	rv := objc.Call[string](d_, objc.Sel("border"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537607-borderwidth?language=objc
func (d_ DOMCSSStyleDeclaration) BorderWidth() string {
	rv := objc.Call[string](d_, objc.Sel("borderWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537141-setborderleft?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderLeft:"), borderLeft)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538198-borderrightstyle?language=objc
func (d_ DOMCSSStyleDeclaration) BorderRightStyle() string {
	rv := objc.Call[string](d_, objc.Sel("borderRightStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536544-borderleft?language=objc
func (d_ DOMCSSStyleDeclaration) BorderLeft() string {
	rv := objc.Call[string](d_, objc.Sel("borderLeft"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536316-pauseafter?language=objc
func (d_ DOMCSSStyleDeclaration) PauseAfter() string {
	rv := objc.Call[string](d_, objc.Sel("pauseAfter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536391-setliststyletype?language=objc
func (d_ DOMCSSStyleDeclaration) SetListStyleType(listStyleType string) {
	objc.Call[objc.Void](d_, objc.Sel("setListStyleType:"), listStyleType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537940-settablelayout?language=objc
func (d_ DOMCSSStyleDeclaration) SetTableLayout(tableLayout string) {
	objc.Call[objc.Void](d_, objc.Sel("setTableLayout:"), tableLayout)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537915-setoutlinestyle?language=objc
func (d_ DOMCSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	objc.Call[objc.Void](d_, objc.Sel("setOutlineStyle:"), outlineStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537361-setmarginbottom?language=objc
func (d_ DOMCSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	objc.Call[objc.Void](d_, objc.Sel("setMarginBottom:"), marginBottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538133-setmaxwidth?language=objc
func (d_ DOMCSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	objc.Call[objc.Void](d_, objc.Sel("setMaxWidth:"), maxWidth)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537758-setmarginright?language=objc
func (d_ DOMCSSStyleDeclaration) SetMarginRight(marginRight string) {
	objc.Call[objc.Void](d_, objc.Sel("setMarginRight:"), marginRight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537631-pause?language=objc
func (d_ DOMCSSStyleDeclaration) Pause() string {
	rv := objc.Call[string](d_, objc.Sel("pause"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536928-setvolume?language=objc
func (d_ DOMCSSStyleDeclaration) SetVolume(volume string) {
	objc.Call[objc.Void](d_, objc.Sel("setVolume:"), volume)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537513-setcueafter?language=objc
func (d_ DOMCSSStyleDeclaration) SetCueAfter(cueAfter string) {
	objc.Call[objc.Void](d_, objc.Sel("setCueAfter:"), cueAfter)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537947-setmaxheight?language=objc
func (d_ DOMCSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	objc.Call[objc.Void](d_, objc.Sel("setMaxHeight:"), maxHeight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536310-markeroffset?language=objc
func (d_ DOMCSSStyleDeclaration) MarkerOffset() string {
	rv := objc.Call[string](d_, objc.Sel("markerOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537256-page?language=objc
func (d_ DOMCSSStyleDeclaration) Page() string {
	rv := objc.Call[string](d_, objc.Sel("page"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537553-zindex?language=objc
func (d_ DOMCSSStyleDeclaration) ZIndex() string {
	rv := objc.Call[string](d_, objc.Sel("zIndex"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537996-pitch?language=objc
func (d_ DOMCSSStyleDeclaration) Pitch() string {
	rv := objc.Call[string](d_, objc.Sel("pitch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538098-setliststyleimage?language=objc
func (d_ DOMCSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	objc.Call[objc.Void](d_, objc.Sel("setListStyleImage:"), listStyleImage)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538219-outlinestyle?language=objc
func (d_ DOMCSSStyleDeclaration) OutlineStyle() string {
	rv := objc.Call[string](d_, objc.Sel("outlineStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538103-size?language=objc
func (d_ DOMCSSStyleDeclaration) Size() string {
	rv := objc.Call[string](d_, objc.Sel("size"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536390-marginbottom?language=objc
func (d_ DOMCSSStyleDeclaration) MarginBottom() string {
	rv := objc.Call[string](d_, objc.Sel("marginBottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537602-setlineheight?language=objc
func (d_ DOMCSSStyleDeclaration) SetLineHeight(lineHeight string) {
	objc.Call[objc.Void](d_, objc.Sel("setLineHeight:"), lineHeight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537212-backgroundposition?language=objc
func (d_ DOMCSSStyleDeclaration) BackgroundPosition() string {
	rv := objc.Call[string](d_, objc.Sel("backgroundPosition"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536298-marginleft?language=objc
func (d_ DOMCSSStyleDeclaration) MarginLeft() string {
	rv := objc.Call[string](d_, objc.Sel("marginLeft"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537311-setcursor?language=objc
func (d_ DOMCSSStyleDeclaration) SetCursor(cursor string) {
	objc.Call[objc.Void](d_, objc.Sel("setCursor:"), cursor)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537899-borderleftstyle?language=objc
func (d_ DOMCSSStyleDeclaration) BorderLeftStyle() string {
	rv := objc.Call[string](d_, objc.Sel("borderLeftStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1538222-setbackgroundposition?language=objc
func (d_ DOMCSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundPosition:"), backgroundPosition)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536575-setborderright?language=objc
func (d_ DOMCSSStyleDeclaration) SetBorderRight(borderRight string) {
	objc.Call[objc.Void](d_, objc.Sel("setBorderRight:"), borderRight)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537576-fontstyle?language=objc
func (d_ DOMCSSStyleDeclaration) FontStyle() string {
	rv := objc.Call[string](d_, objc.Sel("fontStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537637-liststyleposition?language=objc
func (d_ DOMCSSStyleDeclaration) ListStylePosition() string {
	rv := objc.Call[string](d_, objc.Sel("listStylePosition"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537543-setpaddingleft?language=objc
func (d_ DOMCSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingLeft:"), paddingLeft)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1536269-setmarkeroffset?language=objc
func (d_ DOMCSSStyleDeclaration) SetMarkerOffset(markerOffset string) {
	objc.Call[objc.Void](d_, objc.Sel("setMarkerOffset:"), markerOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstyledeclaration/1537762-fontvariant?language=objc
func (d_ DOMCSSStyleDeclaration) FontVariant() string {
	rv := objc.Call[string](d_, objc.Sel("fontVariant"))
	return rv
}
