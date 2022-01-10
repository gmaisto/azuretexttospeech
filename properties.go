package azuretexttospeech

// AudioOutput types represent the supported audio encoding formats for the text-to-speech endpoint.
// This type is required when requesting to azuretexttospeech.Synthesize text-to-speed request.
// Each incorporates a bitrate and encoding type. The Speech service supports 24 kHz, 16 kHz, and 8 kHz audio outputs.
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#audio-outputs

type AudioOutput int

const (
	RAW16khz16bitMonoPCM AudioOutput = iota
	RAW24khz16bitMonoPCM
	RAW48khz16bitMonoPCM
	RAW8khz8bitMonoMulaw
	RAW8khz8bitMonoAlaw
	AUDIO16khz32kbitrateMonoMP3
	AUDIO16khz128kbitrateMonoMP3
	AUDIO24khz96kbitrateMonoMP3
	AUDIO48khz96kbitrateMonoMP3
	RAW16khz16bitMonoTruesilk
	WEBM16khz16bitMonoOpus
	OGG16khz16bitMonoOpus
	OGG48khz16bitMonoOpus
	RIFF16khz16bitMonoPCM
	RIFF24khz16bitMonoPCM
	RIFF48khz16bitMonoPCM
	RIFF8khz8bitMonoMulaw
	RIFF8khz8bitMonoAlaw
	AUDIO16khz64kbitrateMonoMP3
	AUDIO24khz48kbitrateMonoMP3
	AUDIO24khz160kbitrateMonoMP3
	AUDIO48khz192kbitrateMonoMP3
	RAW24khz16bitMonoTruesilk
	WEBM24khz16bitMonoOpus
	OGG24khz16bitMonoOpus
)

var MapAudioFileExtensions = map[string]string{
	"RAW16khz16bitMonoPCM":         "raw",
	"RAW24khz16bitMonoPCM":         "raw",
	"RAW48khz16bitMonoPCM":         "raw",
	"RAW8khz8bitMonoMulaw":         "raw",
	"RAW8khz8bitMonoAlaw":          "raw",
	"AUDIO16khz32kbitrateMonoMP3":  "mp3",
	"AUDIO16khz128kbitrateMonoMP3": "mp3",
	"AUDIO24khz96kbitrateMonoMP3":  "mp3",
	"AUDIO48khz96kbitrateMonoMP3":  "mp3",
	"RAW16khz16bitMonoTruesilk":    "raw",
	"WEBM16khz16bitMonoOpus":       "opus",
	"OGG16khz16bitMonoOpus":        "opus",
	"OGG48khz16bitMonoOpus":        "opus",
	"RIFF16khz16bitMonoPCM":        "wav",
	"RIFF24khz16bitMonoPCM":        "wav",
	"RIFF48khz16bitMonoPCM":        "wav",
	"RIFF8khz8bitMonoMulaw":        "wav",
	"RIFF8khz8bitMonoAlaw":         "wav",
	"AUDIO16khz64kbitrateMonoMP3":  "mp3",
	"AUDIO24khz48kbitrateMonoMP3":  "mp3",
	"AUDIO24khz160kbitrateMonoMP3": "mp3",
	"AUDIO48khz192kbitrateMonoMP3": "mp3",
	"RAW24khz16bitMonoTruesilk":    "raw",
	"WEBM24khz16bitMonoOpus":       "webm",
	"OGG24khz16bitMonoOpus":        "ogg",
}

var MapAudioToFormatid = map[string]AudioOutput{
	"RAW16khz16bitMonoPCM":         RAW16khz16bitMonoPCM,
	"RAW24khz16bitMonoPCM":         RAW24khz16bitMonoPCM,
	"RAW48khz16bitMonoPCM":         RAW48khz16bitMonoPCM,
	"RAW8khz8bitMonoMulaw":         RAW8khz8bitMonoMulaw,
	"RAW8khz8bitMonoAlaw":          RAW8khz8bitMonoAlaw,
	"AUDIO16khz32kbitrateMonoMP3":  AUDIO16khz32kbitrateMonoMP3,
	"AUDIO16khz128kbitrateMonoMP3": AUDIO16khz128kbitrateMonoMP3,
	"AUDIO24khz96kbitrateMonoMP3":  AUDIO24khz96kbitrateMonoMP3,
	"AUDIO48khz96kbitrateMonoMP3":  AUDIO48khz96kbitrateMonoMP3,
	"RAW16khz16bitMonoTruesilk":    RAW16khz16bitMonoTruesilk,
	"WEBM16khz16bitMonoOpus":       WEBM16khz16bitMonoOpus,
	"OGG16khz16bitMonoOpus":        OGG16khz16bitMonoOpus,
	"OGG48khz16bitMonoOpus":        OGG48khz16bitMonoOpus,
	"RIFF16khz16bitMonoPCM":        RIFF16khz16bitMonoPCM,
	"RIFF24khz16bitMonoPCM":        RIFF24khz16bitMonoPCM,
	"RIFF48khz16bitMonoPCM":        RIFF48khz16bitMonoPCM,
	"RIFF8khz8bitMonoMulaw":        RIFF8khz8bitMonoMulaw,
	"RIFF8khz8bitMonoAlaw":         RIFF8khz8bitMonoAlaw,
	"AUDIO16khz64kbitrateMonoMP3":  AUDIO16khz64kbitrateMonoMP3,
	"AUDIO24khz48kbitrateMonoMP3":  AUDIO24khz48kbitrateMonoMP3,
	"AUDIO24khz160kbitrateMonoMP3": AUDIO24khz160kbitrateMonoMP3,
	"AUDIO48khz192kbitrateMonoMP3": AUDIO48khz192kbitrateMonoMP3,
	"RAW24khz16bitMonoTruesilk":    RAW24khz16bitMonoTruesilk,
	"WEBM24khz16bitMonoOpus":       WEBM24khz16bitMonoOpus,
	"OGG24khz16bitMonoOpus":        OGG24khz16bitMonoOpus,
}

func (a AudioOutput) String() string {
	return []string{
		"raw-16khz-16bit-mono-pcm",
		"raw-24khz-16bit-mono-pcm",
		"raw-48khz-16bit-mono-pcm",
		"raw-8khz-8bit-mono-mulaw",
		"raw-8khz-8bit-mono-alaw",
		"audio-16khz-32kbitrate-mono-mp3",
		"audio-16khz-128kbitrate-mono-mp3",
		"audio-24khz-96kbitrate-mono-mp3",
		"audio-48khz-96kbitrate-mono-mp3",
		"raw-16khz-16bit-mono-truesilk",
		"webm-16khz-16bit-mono-opus",
		"ogg-16khz-16bit-mono-opus",
		"ogg-48khz-16bit-mono-opus",
		"riff-16khz-16bit-mono-pcm",
		"riff-24khz-16bit-mono-pcm",
		"riff-48khz-16bit-mono-pcm",
		"riff-8khz-8bit-mono-mulaw",
		"riff-8khz-8bit-mono-alaw",
		"audio-16khz-64kbitrate-mono-mp3",
		"audio-24khz-48kbitrate-mono-mp3",
		"audio-24khz-160kbitrate-mono-mp3",
		"audio-48khz-192kbitrate-mono-mp3",
		"raw-24khz-16bit-mono-truesilk",
		"webm-24khz-16bit-mono-opus",
		"ogg-24khz-16bit-mono-opus",
	}[a]
}

// Gender type for the digitized language
//go:generate enumer -type=Gender -linecomment -json
type Gender int

const (
	// GenderMale , GenderFemale are the static Gender constants for digitized voices.
	// See Gender in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices for breakdown
	GenderMale   Gender = iota // Male
	GenderFemale               // Female
)

// Locale references the language or locale for text-to-speech.
// See "locale" in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices
//go:generate enumer -type=Locale -linecomment -json
type Locale int

const (
	LocaleafZA  Locale = iota //af-ZA
	LocaleamET                //am-ET
	LocalearAE                //ar-AE
	LocalearBH                //ar-BH
	LocalearDZ                //ar-DZ
	LocalearEG                //ar-EG
	LocalearIQ                //ar-IQ
	LocalearJO                //ar-JO
	LocalearKW                //ar-KW
	LocalearLY                //ar-LY
	LocalearMA                //ar-MA
	LocalearQA                //ar-QA
	LocalearSA                //ar-SA
	LocalearSY                //ar-SY
	LocalearTN                //ar-TN
	LocalearYE                //ar-YE
	LocalebgBG                //bg-BG
	LocalebnBD                //bn-BD
	LocalebnIN                //bn-IN
	LocalecaES                //ca-ES
	LocalecsCZ                //cs-CZ
	LocalecyGB                //cy-GB
	LocaledaDK                //da-DK
	LocaledeAT                //de-AT
	LocaledeCH                //de-CH
	LocaledeDE                //de-DE
	LocaleelGR                //el-GR
	LocaleenAU                //en-AU
	LocaleenCA                //en-CA
	LocaleenGB                //en-GB
	LocaleenHK                //en-HK
	LocaleenIE                //en-IE
	LocaleenIN                //en-IN
	LocaleenKE                //en-KE
	LocaleenNG                //en-NG
	LocaleenNZ                //en-NZ
	LocaleenPH                //en-PH
	LocaleenSG                //en-SG
	LocaleenTZ                //en-TZ
	LocaleenUS                //en-US
	LocaleenZA                //en-ZA
	LocaleesAR                //es-AR
	LocaleesBO                //es-BO
	LocaleesCL                //es-CL
	LocaleesCO                //es-CO
	LocaleesCR                //es-CR
	LocaleesCU                //es-CU
	LocaleesDO                //es-DO
	LocaleesEC                //es-EC
	LocaleesES                //es-ES
	LocaleesGQ                //es-GQ
	LocaleesGT                //es-GT
	LocaleesHN                //es-HN
	LocaleesMX                //es-MX
	LocaleesNI                //es-NI
	LocaleesPA                //es-PA
	LocaleesPE                //es-PE
	LocaleesPR                //es-PR
	LocaleesPY                //es-PY
	LocaleesSV                //es-SV
	LocaleesUS                //es-US
	LocaleesUY                //es-UY
	LocaleesVE                //es-VE
	LocaleetEE                //et-EE
	LocalefaIR                //fa-IR
	LocalefiFI                //fi-FI
	LocalefilPH               // fil-PH
	LocalefrBE                //fr-BE
	LocalefrCA                //fr-CA
	LocalefrCH                //fr-CH
	LocalefrFR                //fr-FR
	LocalegaIE                //ga-IE
	LocaleglES                //gl-ES
	LocaleguIN                //gu-IN
	LocaleheIL                //he-IL
	LocalehiIN                //hi-IN
	LocalehrHR                //hr-HR
	LocalehuHU                //hu-HU
	LocaleidID                //id-ID
	LocaleisIS                //is-IS
	LocaleitIT                //it-IT
	LocalejaJP                //ja-JP
	LocalejvID                //jv-ID
	LocalekkKZ                //kk-KZ
	LocalekmKH                //km-KH
	LocaleknIN                //kn-IN
	LocalekoKR                //ko-KR
	LocaleloLA                //lo-LA
	LocaleltLT                //lt-LT
	LocalelvLV                //lv-LV
	LocalemkMK                //mk-MK
	LocalemlIN                //ml-IN
	LocalemrIN                //mr-IN
	LocalemsMY                //ms-MY
	LocalemtMT                //mt-MT
	LocalemyMM                //my-MM
	LocalenbNO                //nb-NO
	LocalenlBE                //nl-BE
	LocalenlNL                //nl-NL
	LocaleplPL                //pl-PL
	LocalepsAF                //ps-AF
	LocaleptBR                //pt-BR
	LocaleptPT                //pt-PT
	LocaleroRO                //ro-RO
	LocaleruRU                //ru-RU
	LocalesiLK                //si-LK
	LocaleskSK                //sk-SK
	LocaleslSI                //sl-SI
	LocalesoSO                //so-SO
	LocalesrRS                //sr-RS
	LocalesuID                //su-ID
	LocalesvSE                //sv-SE
	LocaleswKE                //sw-KE
	LocaleswTZ                //sw-TZ
	LocaletaIN                //ta-IN
	LocaletaLK                //ta-LK
	LocaletaSG                //ta-SG
	LocaleteIN                //te-IN
	LocalethTH                //th-TH
	LocaletrTR                //tr-TR
	LocaleukUA                //uk-UA
	LocaleurIN                //ur-IN
	LocaleurPK                //ur-PK
	LocaleuzUZ                //uz-UZ
	LocaleviVN                //vi-VN
	LocalezhCN                //zh-CN
	LocalezhHK                //zh-HK
	LocalezhTW                //zh-TW
	LocalezuZA                //zu-ZA
)

// Region references the locations of the availability of standard voices.
// See https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/regions#standard-voices
type Region int

const (
	// Azure regions and their endpoints that support the Text To Speech service.
	RegionAustraliaEast Region = iota
	RegionBrazilSouth
	RegionCanadaCentral
	RegionCentralUS
	RegionEastAsia
	RegionEastUS
	RegionEastUS2
	RegionFranceCentral
	RegionIndiaCentral
	RegionJapanEast
	RegionJapanWest
	RegionKoreaCentral
	RegionNorthCentralUS
	RegionNorthEurope
	RegionSouthCentralUS
	RegionSoutheastAsia
	RegionUKSouth
	RegionWestEurope
	RegionWestUS
	RegionWestUS2
)

func (t Region) String() string {
	return [...]string{
		"australiaeast",
		"brazilsouth",
		"canadacentral",
		"centralus",
		"eastasia",
		"eastus",
		"eastus2",
		"francecentral",
		"indiacentral",
		"japaneast",
		"japanwest",
		"koreacentral",
		"northcentralus",
		"northeurope",
		"southcentralus",
		"southeastasia",
		"uksouth",
		"westeurope",
		"westus",
		"westus2",
	}[t]

}
