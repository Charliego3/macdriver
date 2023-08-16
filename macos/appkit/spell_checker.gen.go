// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SpellChecker] class.
var SpellCheckerClass = _SpellCheckerClass{objc.GetClass("NSSpellChecker")}

type _SpellCheckerClass struct {
	objc.Class
}

// An interface definition for the [SpellChecker] class.
type ISpellChecker interface {
	objc.IObject
	LanguageForWordRangeInStringOrthography(range_ foundation.Range, string_ string, orthography foundation.IOrthography) string
	CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, tag int, orthography foundation.IOrthography, wordCount *int) []foundation.TextCheckingResult
	DismissCorrectionIndicatorForView(view IView)
	CountWordsInStringLanguage(stringToCount string, language string) int
	MenuForResultStringOptionsAtLocationInView(result foundation.ITextCheckingResult, checkedString string, options map[TextCheckingOptionKey]objc.IObject, location foundation.Point, view IView) Menu
	Language() string
	LearnWord(word string)
	CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.Range, string_ string, language string, tag int) []string
	CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, details []map[string]objc.IObject) foundation.Range
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, tag int, completionHandler func(sequenceNumber int, candidates []foundation.TextCheckingResult)) int
	ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler(type_ CorrectionIndicatorType, primaryString string, alternativeStrings []string, rectOfTypedString foundation.Rect, view IView, completionBlock func(acceptedString string))
	UnlearnWord(word string)
	RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response CorrectionResponse, correction string, word string, language string, tag int)
	RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, tag int, completionHandler func(sequenceNumber int, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int)) int
	UpdateSpellingPanelWithMisspelledWord(word string)
	CheckSpellingOfStringStartingAt(stringToCheck string, startingOffset int) foundation.Range
	UpdatePanels()
	SetLanguage(language string) bool
	SetIgnoredWordsInSpellDocumentWithTag(words []string, tag int)
	IgnoreWordInSpellDocumentWithTag(wordToIgnore string, tag int)
	UpdateSpellingPanelWithGrammarStringDetail(string_ string, detail map[string]objc.IObject)
	CloseSpellDocumentWithTag(tag int)
	UserQuotesArrayForLanguage(language string) []string
	SetWordFieldStringValue(string_ string)
	PreventsAutocorrectionBeforeStringLanguage(string_ string, language string) bool
	HasLearnedWord(word string) bool
	IgnoredWordsInSpellDocumentWithTag(tag int) []string
	CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.Range, string_ string, language string, tag int) string
	DeletesAutospaceBetweenStringAndStringLanguage(precedingString string, followingString string, language string) bool
	GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.Range, string_ string, language string, tag int) []string
	UserReplacementsDictionary() map[string]string
	UserPreferredLanguages() []string
	SubstitutionsPanel() Panel
	SpellingPanel() Panel
	SubstitutionsPanelAccessoryViewController() ViewController
	SetSubstitutionsPanelAccessoryViewController(value IViewController)
	AccessoryView() View
	SetAccessoryView(value IView)
	AvailableLanguages() []string
	AutomaticallyIdentifiesLanguages() bool
	SetAutomaticallyIdentifiesLanguages(value bool)
}

// An interface to the Cocoa spell-checking service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker?language=objc
type SpellChecker struct {
	objc.Object
}

func SpellCheckerFrom(ptr unsafe.Pointer) SpellChecker {
	return SpellChecker{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SpellCheckerClass) Alloc() SpellChecker {
	rv := objc.Call[SpellChecker](sc, objc.Sel("alloc"))
	return rv
}

func SpellChecker_Alloc() SpellChecker {
	return SpellCheckerClass.Alloc()
}

func (sc _SpellCheckerClass) New() SpellChecker {
	rv := objc.Call[SpellChecker](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSpellChecker() SpellChecker {
	return SpellCheckerClass.New()
}

func (s_ SpellChecker) Init() SpellChecker {
	rv := objc.Call[SpellChecker](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1530303-languageforwordrange?language=objc
func (s_ SpellChecker) LanguageForWordRangeInStringOrthography(range_ foundation.Range, string_ string, orthography foundation.IOrthography) string {
	rv := objc.Call[string](s_, objc.Sel("languageForWordRange:inString:orthography:"), range_, string_, objc.Ptr(orthography))
	return rv
}

// Requests unified text checking for the given range of the given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535739-checkstring?language=objc
func (s_ SpellChecker) CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, tag int, orthography foundation.IOrthography, wordCount *int) []foundation.TextCheckingResult {
	rv := objc.Call[[]foundation.TextCheckingResult](s_, objc.Sel("checkString:range:types:options:inSpellDocumentWithTag:orthography:wordCount:"), stringToCheck, range_, checkingTypes, options, tag, objc.Ptr(orthography), wordCount)
	return rv
}

// Dismisses the correction indicator for the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535527-dismisscorrectionindicatorforvie?language=objc
func (s_ SpellChecker) DismissCorrectionIndicatorForView(view IView) {
	objc.Call[objc.Void](s_, objc.Sel("dismissCorrectionIndicatorForView:"), objc.Ptr(view))
}

// Returns the number of words in the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1528008-countwordsinstring?language=objc
func (s_ SpellChecker) CountWordsInStringLanguage(stringToCount string, language string) int {
	rv := objc.Call[int](s_, objc.Sel("countWordsInString:language:"), stringToCount, language)
	return rv
}

// Provides a menu containing contextual menu items suitable for certain kinds of detected results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531603-menuforresult?language=objc
func (s_ SpellChecker) MenuForResultStringOptionsAtLocationInView(result foundation.ITextCheckingResult, checkedString string, options map[TextCheckingOptionKey]objc.IObject, location foundation.Point, view IView) Menu {
	rv := objc.Call[Menu](s_, objc.Sel("menuForResult:string:options:atLocation:inView:"), objc.Ptr(result), checkedString, options, location, objc.Ptr(view))
	return rv
}

// Returns the current language used in spell checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535648-language?language=objc
func (s_ SpellChecker) Language() string {
	rv := objc.Call[string](s_, objc.Sel("language"))
	return rv
}

// Adds the word to the spell checker dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1534837-learnword?language=objc
func (s_ SpellChecker) LearnWord(word string) {
	objc.Call[objc.Void](s_, objc.Sel("learnWord:"), word)
}

// Provides a list of complete words that the user might be trying to type based on a partial word in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531406-completionsforpartialwordrange?language=objc
func (s_ SpellChecker) CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.Range, string_ string, language string, tag int) []string {
	rv := objc.Call[[]string](s_, objc.Sel("completionsForPartialWordRange:inString:language:inSpellDocumentWithTag:"), range_, string_, language, tag)
	return rv
}

// Initiates a grammatical analysis of a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1528441-checkgrammarofstring?language=objc
func (s_ SpellChecker) CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, details []map[string]objc.IObject) foundation.Range {
	rv := objc.Call[foundation.Range](s_, objc.Sel("checkGrammarOfString:startingAt:language:wrap:inSpellDocumentWithTag:details:"), stringToCheck, startingOffset, language, wrapFlag, tag, details)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2560997-requestcandidatesforselectedrang?language=objc
func (s_ SpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.Range, stringToCheck string, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, tag int, completionHandler func(sequenceNumber int, candidates []foundation.TextCheckingResult)) int {
	rv := objc.Call[int](s_, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, stringToCheck, checkingTypes, options, tag, completionHandler)
	return rv
}

// Display a suitable user interface to indicate a correction may need to be made. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1524316-showcorrectionindicatoroftype?language=objc
func (s_ SpellChecker) ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler(type_ CorrectionIndicatorType, primaryString string, alternativeStrings []string, rectOfTypedString foundation.Rect, view IView, completionBlock func(acceptedString string)) {
	objc.Call[objc.Void](s_, objc.Sel("showCorrectionIndicatorOfType:primaryString:alternativeStrings:forStringInRect:view:completionHandler:"), type_, primaryString, alternativeStrings, rectOfTypedString, objc.Ptr(view), completionBlock)
}

// Tells the spell checker to unlearn a given word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1525147-unlearnword?language=objc
func (s_ SpellChecker) UnlearnWord(word string) {
	objc.Call[objc.Void](s_, objc.Sel("unlearnWord:"), word)
}

// Records the user response to the correction indicator being displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535568-recordresponse?language=objc
func (s_ SpellChecker) RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response CorrectionResponse, correction string, word string, language string, tag int) {
	objc.Call[objc.Void](s_, objc.Sel("recordResponse:toCorrection:forWord:language:inSpellDocumentWithTag:"), response, correction, word, language, tag)
}

// Requests that the string be checked in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1526261-requestcheckingofstring?language=objc
func (s_ SpellChecker) RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, tag int, completionHandler func(sequenceNumber int, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int)) int {
	rv := objc.Call[int](s_, objc.Sel("requestCheckingOfString:range:types:options:inSpellDocumentWithTag:completionHandler:"), stringToCheck, range_, checkingTypes, options, tag, completionHandler)
	return rv
}

// Causes the spell checker to update the Spelling panel’s misspelled-word field to reflect word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1524852-updatespellingpanelwithmisspelle?language=objc
func (s_ SpellChecker) UpdateSpellingPanelWithMisspelledWord(word string) {
	objc.Call[objc.Void](s_, objc.Sel("updateSpellingPanelWithMisspelledWord:"), word)
}

// Starts the search for a misspelled word in stringToCheck starting at startingOffset within the string object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1532957-checkspellingofstring?language=objc
func (s_ SpellChecker) CheckSpellingOfStringStartingAt(stringToCheck string, startingOffset int) foundation.Range {
	rv := objc.Call[foundation.Range](s_, objc.Sel("checkSpellingOfString:startingAt:"), stringToCheck, startingOffset)
	return rv
}

// Updates the available panels to account for user changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531045-updatepanels?language=objc
func (s_ SpellChecker) UpdatePanels() {
	objc.Call[objc.Void](s_, objc.Sel("updatePanels"))
}

// Returns whether the specified language is in the Spelling pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1532988-setlanguage?language=objc
func (s_ SpellChecker) SetLanguage(language string) bool {
	rv := objc.Call[bool](s_, objc.Sel("setLanguage:"), language)
	return rv
}

// Initializes the ignored-words document (a dictionary identified by tag with someWords), an array of words to ignore. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535879-setignoredwords?language=objc
func (s_ SpellChecker) SetIgnoredWordsInSpellDocumentWithTag(words []string, tag int) {
	objc.Call[objc.Void](s_, objc.Sel("setIgnoredWords:inSpellDocumentWithTag:"), words, tag)
}

// Instructs the spell checker to ignore all future occurrences of wordToIgnore in the document identified by tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531220-ignoreword?language=objc
func (s_ SpellChecker) IgnoreWordInSpellDocumentWithTag(wordToIgnore string, tag int) {
	objc.Call[objc.Void](s_, objc.Sel("ignoreWord:inSpellDocumentWithTag:"), wordToIgnore, tag)
}

// Specifies a grammar-analysis detail to highlight in the Spelling panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1525300-updatespellingpanelwithgrammarst?language=objc
func (s_ SpellChecker) UpdateSpellingPanelWithGrammarStringDetail(string_ string, detail map[string]objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("updateSpellingPanelWithGrammarString:detail:"), string_, detail)
}

// Notifies the receiver that the user has finished with the tagged document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1534066-closespelldocumentwithtag?language=objc
func (s_ SpellChecker) CloseSpellDocumentWithTag(tag int) {
	objc.Call[objc.Void](s_, objc.Sel("closeSpellDocumentWithTag:"), tag)
}

// Returns the default values for quote replacement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1532171-userquotesarrayforlanguage?language=objc
func (s_ SpellChecker) UserQuotesArrayForLanguage(language string) []string {
	rv := objc.Call[[]string](s_, objc.Sel("userQuotesArrayForLanguage:"), language)
	return rv
}

// Sets the string that appears in the misspelled word field, using the string object aString. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1526688-setwordfieldstringvalue?language=objc
func (s_ SpellChecker) SetWordFieldStringValue(string_ string) {
	objc.Call[objc.Void](s_, objc.Sel("setWordFieldStringValue:"), string_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1792008-preventsautocorrectionbeforestri?language=objc
func (s_ SpellChecker) PreventsAutocorrectionBeforeStringLanguage(string_ string, language string) bool {
	rv := objc.Call[bool](s_, objc.Sel("preventsAutocorrectionBeforeString:language:"), string_, language)
	return rv
}

// Indicates whether the spell checker has learned a given word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1527525-haslearnedword?language=objc
func (s_ SpellChecker) HasLearnedWord(word string) bool {
	rv := objc.Call[bool](s_, objc.Sel("hasLearnedWord:"), word)
	return rv
}

// Returns the array of ignored words for a document identified by tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531488-ignoredwordsinspelldocumentwitht?language=objc
func (s_ SpellChecker) IgnoredWordsInSpellDocumentWithTag(tag int) []string {
	rv := objc.Call[[]string](s_, objc.Sel("ignoredWordsInSpellDocumentWithTag:"), tag)
	return rv
}

// Returns a single proposed correction if a word is mis-spelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531542-correctionforwordrange?language=objc
func (s_ SpellChecker) CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.Range, string_ string, language string, tag int) string {
	rv := objc.Call[string](s_, objc.Sel("correctionForWordRange:inString:language:inSpellDocumentWithTag:"), range_, string_, language, tag)
	return rv
}

// Returns a guaranteed unique tag to use as the spell-document tag for a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1526192-uniquespelldocumenttag?language=objc
func (sc _SpellCheckerClass) UniqueSpellDocumentTag() int {
	rv := objc.Call[int](sc, objc.Sel("uniqueSpellDocumentTag"))
	return rv
}

// Returns a guaranteed unique tag to use as the spell-document tag for a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1526192-uniquespelldocumenttag?language=objc
func SpellChecker_UniqueSpellDocumentTag() int {
	return SpellCheckerClass.UniqueSpellDocumentTag()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2561001-deletesautospacebetweenstring?language=objc
func (s_ SpellChecker) DeletesAutospaceBetweenStringAndStringLanguage(precedingString string, followingString string, language string) bool {
	rv := objc.Call[bool](s_, objc.Sel("deletesAutospaceBetweenString:andString:language:"), precedingString, followingString, language)
	return rv
}

// Returns an array of possible substitutions for the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1527419-guessesforwordrange?language=objc
func (s_ SpellChecker) GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.Range, string_ string, language string, tag int) []string {
	rv := objc.Call[[]string](s_, objc.Sel("guessesForWordRange:inString:language:inSpellDocumentWithTag:"), range_, string_, language, tag)
	return rv
}

// Returns the dictionary used when replacing words. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1524925-userreplacementsdictionary?language=objc
func (s_ SpellChecker) UserReplacementsDictionary() map[string]string {
	rv := objc.Call[map[string]string](s_, objc.Sel("userReplacementsDictionary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869579-automatictextreplacementenabled?language=objc
func (sc _SpellCheckerClass) AutomaticTextReplacementEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticTextReplacementEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869579-automatictextreplacementenabled?language=objc
func SpellChecker_AutomaticTextReplacementEnabled() bool {
	return SpellCheckerClass.AutomaticTextReplacementEnabled()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869581-automaticquotesubstitutionenable?language=objc
func (sc _SpellCheckerClass) AutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticQuoteSubstitutionEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869581-automaticquotesubstitutionenable?language=objc
func SpellChecker_AutomaticQuoteSubstitutionEnabled() bool {
	return SpellCheckerClass.AutomaticQuoteSubstitutionEnabled()
}

// Provides a subset of the available languages to be used for spell checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1525173-userpreferredlanguages?language=objc
func (s_ SpellChecker) UserPreferredLanguages() []string {
	rv := objc.Call[[]string](s_, objc.Sel("userPreferredLanguages"))
	return rv
}

// Returns the substitutions panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1534172-substitutionspanel?language=objc
func (s_ SpellChecker) SubstitutionsPanel() Panel {
	rv := objc.Call[Panel](s_, objc.Sel("substitutionsPanel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869584-automaticperiodsubstitutionenabl?language=objc
func (sc _SpellCheckerClass) AutomaticPeriodSubstitutionEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticPeriodSubstitutionEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869584-automaticperiodsubstitutionenabl?language=objc
func SpellChecker_AutomaticPeriodSubstitutionEnabled() bool {
	return SpellCheckerClass.AutomaticPeriodSubstitutionEnabled()
}

// Returns the spell checker’s panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1532806-spellingpanel?language=objc
func (s_ SpellChecker) SpellingPanel() Panel {
	rv := objc.Call[Panel](s_, objc.Sel("spellingPanel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869580-automatictextcompletionenabled?language=objc
func (sc _SpellCheckerClass) AutomaticTextCompletionEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869580-automatictextcompletionenabled?language=objc
func SpellChecker_AutomaticTextCompletionEnabled() bool {
	return SpellCheckerClass.AutomaticTextCompletionEnabled()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869577-automaticdashsubstitutionenabled?language=objc
func (sc _SpellCheckerClass) AutomaticDashSubstitutionEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticDashSubstitutionEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869577-automaticdashsubstitutionenabled?language=objc
func SpellChecker_AutomaticDashSubstitutionEnabled() bool {
	return SpellCheckerClass.AutomaticDashSubstitutionEnabled()
}

// Returns whether the application’s NSSpellChecker has already been created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535301-sharedspellcheckerexists?language=objc
func (sc _SpellCheckerClass) SharedSpellCheckerExists() bool {
	rv := objc.Call[bool](sc, objc.Sel("sharedSpellCheckerExists"))
	return rv
}

// Returns whether the application’s NSSpellChecker has already been created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1535301-sharedspellcheckerexists?language=objc
func SpellChecker_SharedSpellCheckerExists() bool {
	return SpellCheckerClass.SharedSpellCheckerExists()
}

// Sets the substitutions panel’s accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531645-substitutionspanelaccessoryviewc?language=objc
func (s_ SpellChecker) SubstitutionsPanelAccessoryViewController() ViewController {
	rv := objc.Call[ViewController](s_, objc.Sel("substitutionsPanelAccessoryViewController"))
	return rv
}

// Sets the substitutions panel’s accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1531645-substitutionspanelaccessoryviewc?language=objc
func (s_ SpellChecker) SetSubstitutionsPanelAccessoryViewController(value IViewController) {
	objc.Call[objc.Void](s_, objc.Sel("setSubstitutionsPanelAccessoryViewController:"), objc.Ptr(value))
}

// Returns the NSSpellChecker (one per application). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1529639-sharedspellchecker?language=objc
func (sc _SpellCheckerClass) SharedSpellChecker() SpellChecker {
	rv := objc.Call[SpellChecker](sc, objc.Sel("sharedSpellChecker"))
	return rv
}

// Returns the NSSpellChecker (one per application). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1529639-sharedspellchecker?language=objc
func SpellChecker_SharedSpellChecker() SpellChecker {
	return SpellCheckerClass.SharedSpellChecker()
}

// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1528160-accessoryview?language=objc
func (s_ SpellChecker) AccessoryView() View {
	rv := objc.Call[View](s_, objc.Sel("accessoryView"))
	return rv
}

// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1528160-accessoryview?language=objc
func (s_ SpellChecker) SetAccessoryView(value IView) {
	objc.Call[objc.Void](s_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// Provides a list of all available languages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1530496-availablelanguages?language=objc
func (s_ SpellChecker) AvailableLanguages() []string {
	rv := objc.Call[[]string](s_, objc.Sel("availableLanguages"))
	return rv
}

// Sets whether the spell checker will automatically identify languages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1534335-automaticallyidentifieslanguages?language=objc
func (s_ SpellChecker) AutomaticallyIdentifiesLanguages() bool {
	rv := objc.Call[bool](s_, objc.Sel("automaticallyIdentifiesLanguages"))
	return rv
}

// Sets whether the spell checker will automatically identify languages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/1534335-automaticallyidentifieslanguages?language=objc
func (s_ SpellChecker) SetAutomaticallyIdentifiesLanguages(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAutomaticallyIdentifiesLanguages:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869578-automaticspellingcorrectionenabl?language=objc
func (sc _SpellCheckerClass) AutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticSpellingCorrectionEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869578-automaticspellingcorrectionenabl?language=objc
func SpellChecker_AutomaticSpellingCorrectionEnabled() bool {
	return SpellCheckerClass.AutomaticSpellingCorrectionEnabled()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869583-automaticcapitalizationenabled?language=objc
func (sc _SpellCheckerClass) AutomaticCapitalizationEnabled() bool {
	rv := objc.Call[bool](sc, objc.Sel("automaticCapitalizationEnabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspellchecker/2869583-automaticcapitalizationenabled?language=objc
func SpellChecker_AutomaticCapitalizationEnabled() bool {
	return SpellCheckerClass.AutomaticCapitalizationEnabled()
}
