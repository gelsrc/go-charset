// Copyright (c) AO Gelicon Consulting 2018. All rights reserved.
// Исключительные права на Файл принадлежат АО Геликон Консалтинг.

package charset

// Byte to rune mapping range from 0x80 to 0xff (128 to 255).
//
// [Ђ] [Ѓ] [‚] [ѓ] [„] […] [†] [‡]
// [€] [‰] [Љ] [‹] [Њ] [Ќ] [Ћ] [Џ]
// [ђ] [‘] [’] [“] [”] [•] [–] [—]
// --- [™] [љ] [›] [њ] [ќ] [ћ] [џ]
// [ ] [Ў] [ў] [Ј] [¤] [Ґ] [¦] [§]
// [Ё] [©] [Є] [«] [¬] [­] [®] [Ї]
// [°] [±] [І] [і] [ґ] [µ] [¶] [·]
// [ё] [№] [є] [»] [ј] [Ѕ] [ѕ] [ї]
// [А] [Б] [В] [Г] [Д] [Е] [Ж] [З]
// [И] [Й] [К] [Л] [М] [Н] [О] [П]
// [Р] [С] [Т] [У] [Ф] [Х] [Ц] [Ч]
// [Ш] [Щ] [Ъ] [Ы] [Ь] [Э] [Ю] [Я]
// [а] [б] [в] [г] [д] [е] [ж] [з]
// [и] [й] [к] [л] [м] [н] [о] [п]
// [р] [с] [т] [у] [ф] [х] [ц] [ч]
// [ш] [щ] [ъ] [ы] [ь] [э] [ю] [я]
var cp1251Runes = [128]rune{
	0x0402, 0x0403, 0x201a, 0x0453, 0x201e, 0x2026, 0x2020, 0x2021,
	0x20ac, 0x2030, 0x0409, 0x2039, 0x040a, 0x040c, 0x040b, 0x040f,
	0x0452, 0x2018, 0x2019, 0x201c, 0x201d, 0x2022, 0x2013, 0x2014,
	0xfffd, 0x2122, 0x0459, 0x203a, 0x045a, 0x045c, 0x045b, 0x045f,
	0x00a0, 0x040e, 0x045e, 0x0408, 0x00a4, 0x0490, 0x00a6, 0x00a7,
	0x0401, 0x00a9, 0x0404, 0x00ab, 0x00ac, 0x00ad, 0x00ae, 0x0407,
	0x00b0, 0x00b1, 0x0406, 0x0456, 0x0491, 0x00b5, 0x00b6, 0x00b7,
	0x0451, 0x2116, 0x0454, 0x00bb, 0x0458, 0x0405, 0x0455, 0x0457,
	0x0410, 0x0411, 0x0412, 0x0413, 0x0414, 0x0415, 0x0416, 0x0417,
	0x0418, 0x0419, 0x041a, 0x041b, 0x041c, 0x041d, 0x041e, 0x041f,
	0x0420, 0x0421, 0x0422, 0x0423, 0x0424, 0x0425, 0x0426, 0x0427,
	0x0428, 0x0429, 0x042a, 0x042b, 0x042c, 0x042d, 0x042e, 0x042f,
	0x0430, 0x0431, 0x0432, 0x0433, 0x0434, 0x0435, 0x0436, 0x0437,
	0x0438, 0x0439, 0x043a, 0x043b, 0x043c, 0x043d, 0x043e, 0x043f,
	0x0440, 0x0441, 0x0442, 0x0443, 0x0444, 0x0445, 0x0446, 0x0447,
	0x0448, 0x0449, 0x044a, 0x044b, 0x044c, 0x044d, 0x044e, 0x044f,
}

// Rune to byte mapping range from 0xa0 to 0xbb (160 to 187).
//
// [ ] --- --- --- [¤] --- [¦] [§]
// --- [©] --- [«] [¬] [­] [®] ---
// [°] [±] --- --- --- [µ] [¶] [·]
// --- --- --- [»]
var cp1251RangeFrom0xa0To0xbb = [28]byte{
	0xa0, 0x98, 0x98, 0x98, 0xa4, 0x98, 0xa6, 0xa7,
	0x98, 0xa9, 0x98, 0xab, 0xac, 0xad, 0xae, 0x98,
	0xb0, 0xb1, 0x98, 0x98, 0x98, 0xb5, 0xb6, 0xb7,
	0x98, 0x98, 0x98, 0xbb,
}

// Rune to byte mapping range from 0x0401 to 0x045f (1025 to 1119).
//
// [Ё] [Ђ] [Ѓ] [Є] [Ѕ] [І] [Ї] [Ј]
// [Љ] [Њ] [Ћ] [Ќ] --- [Ў] [Џ] [А]
// [Б] [В] [Г] [Д] [Е] [Ж] [З] [И]
// [Й] [К] [Л] [М] [Н] [О] [П] [Р]
// [С] [Т] [У] [Ф] [Х] [Ц] [Ч] [Ш]
// [Щ] [Ъ] [Ы] [Ь] [Э] [Ю] [Я] [а]
// [б] [в] [г] [д] [е] [ж] [з] [и]
// [й] [к] [л] [м] [н] [о] [п] [р]
// [с] [т] [у] [ф] [х] [ц] [ч] [ш]
// [щ] [ъ] [ы] [ь] [э] [ю] [я] ---
// [ё] [ђ] [ѓ] [є] [ѕ] [і] [ї] [ј]
// [љ] [њ] [ћ] [ќ] --- [ў] [џ]
var cp1251RangeFrom0x0401To0x045f = [95]byte{
	0xa8, 0x80, 0x81, 0xaa, 0xbd, 0xb2, 0xaf, 0xa3,
	0x8a, 0x8c, 0x8e, 0x8d, 0x98, 0xa1, 0x8f, 0xc0,
	0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8,
	0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf, 0xd0,
	0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8,
	0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xe0,
	0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7, 0xe8,
	0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef, 0xf0,
	0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8,
	0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff, 0x98,
	0xb8, 0x90, 0x83, 0xba, 0xbe, 0xb3, 0xbf, 0xbc,
	0x9a, 0x9c, 0x9e, 0x9d, 0x98, 0xa2, 0x9f,
}

// Rune to byte mapping range from 0x2013 to 0x203a (8211 to 8250).
//
// [–] [—] --- --- --- [‘] [’] [‚]
// --- [“] [”] [„] --- [†] [‡] [•]
// --- --- --- […] --- --- --- ---
// --- --- --- --- --- [‰] --- ---
// --- --- --- --- --- --- [‹] [›]
var cp1251RangeFrom0x2013To0x203a = [40]byte{
	0x96, 0x97, 0x98, 0x98, 0x98, 0x91, 0x92, 0x82,
	0x98, 0x93, 0x94, 0x84, 0x98, 0x86, 0x87, 0x95,
	0x98, 0x98, 0x98, 0x85, 0x98, 0x98, 0x98, 0x98,
	0x98, 0x98, 0x98, 0x98, 0x98, 0x89, 0x98, 0x98,
	0x98, 0x98, 0x98, 0x98, 0x98, 0x98, 0x8b, 0x9b,
}

func Cp1251BytesToRunes(in []byte) []rune {

	out := make([]rune, len(in))

	for k, v := range in {

		var tmp rune

		if v < 0x80 {
			tmp = rune(v)
		} else {
			tmp = cp1251Runes[v-0x80]
		}

		out[k] = tmp
	}

	return out
}

func Cp1251RunesToBytes(in []rune) []byte {

	out := make([]byte, len(in))

	for k, v := range in {

		var tmp byte

		if v < 0x80 {
			tmp = byte(v)
		} else if i := int(v) - 0x0401; i >= 0 && i < len(cp1251RangeFrom0x0401To0x045f) {
			tmp = cp1251RangeFrom0x0401To0x045f[i]
		} else if i := int(v) - 0x2013; i >= 0 && i < len(cp1251RangeFrom0x2013To0x203a) {
			tmp = cp1251RangeFrom0x2013To0x203a[i]
		} else if i := int(v) - 0x00a0; i >= 0 && i < len(cp1251RangeFrom0xa0To0xbb) {
			tmp = cp1251RangeFrom0xa0To0xbb[i]
		} else if v == 0x0490 {
			tmp = 0xa5 // [Ґ]
		} else if v == 0x0491 {
			tmp = 0xb4 // [ґ]
		} else if v == 0x20ac {
			tmp = 0x88 // [€]
		} else if v == 0x2116 {
			tmp = 0xb9 // [№]
		} else if v == 0x2122 {
			tmp = 0x99 // [™]
		} else {
			tmp = 0x98
		}

		out[k] = tmp
	}

	return out
}
