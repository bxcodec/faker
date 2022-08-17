package interfaces

// LangRuneBoundary is for language boundary
type LangRuneBoundary struct {
	Start   rune
	End     rune
	Exclude []rune
}

// Language rune boundaries here
var (
	// LangENG is for english language
	LangENG = LangRuneBoundary{65, 122, []rune{91, 92, 93, 94, 95, 96}}
	// LangCHI is for chinese language
	LangCHI = LangRuneBoundary{19968, 40869, nil}
	// LangRUS is for russian language
	LangRUS = LangRuneBoundary{1025, 1105, nil}
	// LangJPN is for japanese Hiragana Katakana language
	LangJPN = LangRuneBoundary{12353, 12534, []rune{12436, 12437, 12438, 12439, 12440, 12441, 12442, 12443, 12444, 12445, 12446, 12447, 12448}}
	// LangKOR is for korean Hangul language
	LangKOR = LangRuneBoundary{44032, 55203, nil}
	// EmotEMJ is for emoticons
	EmotEMJ = LangRuneBoundary{126976, 129535, nil}
)
