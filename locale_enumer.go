// Code generated by "enumer -type=Locale -linecomment -json"; DO NOT EDIT.

package azuretexttospeech

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _LocaleName = "af-ZAam-ETar-AEar-BHar-DZar-EGar-IQar-JOar-KWar-LYar-MAar-QAar-SAar-SYar-TNar-YEbg-BGbn-BDbn-INca-EScs-CZcy-GBda-DKde-ATde-CHde-DEel-GRen-AUen-CAen-GBen-HKen-IEen-INen-KEen-NGen-NZen-PHen-SGen-TZen-USen-ZAes-ARes-BOes-CLes-COes-CRes-CUes-DOes-ECes-ESes-GQes-GTes-HNes-MXes-NIes-PAes-PEes-PRes-PYes-SVes-USes-UYes-VEet-EEfa-IRfi-FIfil-PHfr-BEfr-CAfr-CHfr-FRga-IEgl-ESgu-INhe-ILhi-INhr-HRhu-HUid-IDis-ISit-ITja-JPjv-IDkm-KHko-KRlt-LTlv-LVmr-INms-MYmt-MTmy-MMnb-NOnl-BEkk-KZnl-NLpl-PLpt-BRpt-PTro-ROru-RUsk-SKsl-SIso-SOsu-IDsv-SEsw-KEsw-TZta-INta-LKta-SGte-INth-THtr-TRuk-UAur-INur-PKuz-UZvi-VNzh-CNzh-HKzh-TWzu-ZAkn-INlo-LA"

var _LocaleIndex = [...]uint16{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120, 125, 130, 135, 140, 145, 150, 155, 160, 165, 170, 175, 180, 185, 190, 195, 200, 205, 210, 215, 220, 225, 230, 235, 240, 245, 250, 255, 260, 265, 270, 275, 280, 285, 290, 295, 300, 305, 310, 315, 320, 325, 330, 336, 341, 346, 351, 356, 361, 366, 371, 376, 381, 386, 391, 396, 401, 406, 411, 416, 421, 426, 431, 436, 441, 446, 451, 456, 461, 466, 471, 476, 481, 486, 491, 496, 501, 506, 511, 516, 521, 526, 531, 536, 541, 546, 551, 556, 561, 566, 571, 576, 581, 586, 591, 596, 601, 606, 611, 616, 621}

const _LocaleLowerName = "af-zaam-etar-aear-bhar-dzar-egar-iqar-joar-kwar-lyar-maar-qaar-saar-syar-tnar-yebg-bgbn-bdbn-inca-escs-czcy-gbda-dkde-atde-chde-deel-gren-auen-caen-gben-hken-ieen-inen-keen-ngen-nzen-phen-sgen-tzen-usen-zaes-ares-boes-cles-coes-cres-cues-does-eces-eses-gqes-gtes-hnes-mxes-nies-paes-pees-pres-pyes-sves-uses-uyes-veet-eefa-irfi-fifil-phfr-befr-cafr-chfr-frga-iegl-esgu-inhe-ilhi-inhr-hrhu-huid-idis-isit-itja-jpjv-idkm-khko-krlt-ltlv-lvmr-inms-mymt-mtmy-mmnb-nonl-bekk-kznl-nlpl-plpt-brpt-ptro-roru-rusk-sksl-siso-sosu-idsv-sesw-kesw-tzta-inta-lkta-sgte-inth-thtr-truk-uaur-inur-pkuz-uzvi-vnzh-cnzh-hkzh-twzu-zakn-inlo-la"

func (i Locale) String() string {
	if i < 0 || i >= Locale(len(_LocaleIndex)-1) {
		return fmt.Sprintf("Locale(%d)", i)
	}
	return _LocaleName[_LocaleIndex[i]:_LocaleIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _LocaleNoOp() {
	var x [1]struct{}
	_ = x[LocaleafZA-(0)]
	_ = x[LocaleamET-(1)]
	_ = x[LocalearAE-(2)]
	_ = x[LocalearBH-(3)]
	_ = x[LocalearDZ-(4)]
	_ = x[LocalearEG-(5)]
	_ = x[LocalearIQ-(6)]
	_ = x[LocalearJO-(7)]
	_ = x[LocalearKW-(8)]
	_ = x[LocalearLY-(9)]
	_ = x[LocalearMA-(10)]
	_ = x[LocalearQA-(11)]
	_ = x[LocalearSA-(12)]
	_ = x[LocalearSY-(13)]
	_ = x[LocalearTN-(14)]
	_ = x[LocalearYE-(15)]
	_ = x[LocalebgBG-(16)]
	_ = x[LocalebnBD-(17)]
	_ = x[LocalebnIN-(18)]
	_ = x[LocalecaES-(19)]
	_ = x[LocalecsCZ-(20)]
	_ = x[LocalecyGB-(21)]
	_ = x[LocaledaDK-(22)]
	_ = x[LocaledeAT-(23)]
	_ = x[LocaledeCH-(24)]
	_ = x[LocaledeDE-(25)]
	_ = x[LocaleelGR-(26)]
	_ = x[LocaleenAU-(27)]
	_ = x[LocaleenCA-(28)]
	_ = x[LocaleenGB-(29)]
	_ = x[LocaleenHK-(30)]
	_ = x[LocaleenIE-(31)]
	_ = x[LocaleenIN-(32)]
	_ = x[LocaleenKE-(33)]
	_ = x[LocaleenNG-(34)]
	_ = x[LocaleenNZ-(35)]
	_ = x[LocaleenPH-(36)]
	_ = x[LocaleenSG-(37)]
	_ = x[LocaleenTZ-(38)]
	_ = x[LocaleenUS-(39)]
	_ = x[LocaleenZA-(40)]
	_ = x[LocaleesAR-(41)]
	_ = x[LocaleesBO-(42)]
	_ = x[LocaleesCL-(43)]
	_ = x[LocaleesCO-(44)]
	_ = x[LocaleesCR-(45)]
	_ = x[LocaleesCU-(46)]
	_ = x[LocaleesDO-(47)]
	_ = x[LocaleesEC-(48)]
	_ = x[LocaleesES-(49)]
	_ = x[LocaleesGQ-(50)]
	_ = x[LocaleesGT-(51)]
	_ = x[LocaleesHN-(52)]
	_ = x[LocaleesMX-(53)]
	_ = x[LocaleesNI-(54)]
	_ = x[LocaleesPA-(55)]
	_ = x[LocaleesPE-(56)]
	_ = x[LocaleesPR-(57)]
	_ = x[LocaleesPY-(58)]
	_ = x[LocaleesSV-(59)]
	_ = x[LocaleesUS-(60)]
	_ = x[LocaleesUY-(61)]
	_ = x[LocaleesVE-(62)]
	_ = x[LocaleetEE-(63)]
	_ = x[LocalefaIR-(64)]
	_ = x[LocalefiFI-(65)]
	_ = x[LocalefilPH-(66)]
	_ = x[LocalefrBE-(67)]
	_ = x[LocalefrCA-(68)]
	_ = x[LocalefrCH-(69)]
	_ = x[LocalefrFR-(70)]
	_ = x[LocalegaIE-(71)]
	_ = x[LocaleglES-(72)]
	_ = x[LocaleguIN-(73)]
	_ = x[LocaleheIL-(74)]
	_ = x[LocalehiIN-(75)]
	_ = x[LocalehrHR-(76)]
	_ = x[LocalehuHU-(77)]
	_ = x[LocaleidID-(78)]
	_ = x[LocaleisIS-(79)]
	_ = x[LocaleitIT-(80)]
	_ = x[LocalejaJP-(81)]
	_ = x[LocalejvID-(82)]
	_ = x[LocalekmKH-(83)]
	_ = x[LocalekoKR-(84)]
	_ = x[LocaleltLT-(85)]
	_ = x[LocalelvLV-(86)]
	_ = x[LocalemrIN-(87)]
	_ = x[LocalemsMY-(88)]
	_ = x[LocalemtMT-(89)]
	_ = x[LocalemyMM-(90)]
	_ = x[LocalenbNO-(91)]
	_ = x[LocalenlBE-(92)]
	_ = x[LocalekkKZ-(93)]
	_ = x[LocalenlNL-(94)]
	_ = x[LocaleplPL-(95)]
	_ = x[LocaleptBR-(96)]
	_ = x[LocaleptPT-(97)]
	_ = x[LocaleroRO-(98)]
	_ = x[LocaleruRU-(99)]
	_ = x[LocaleskSK-(100)]
	_ = x[LocaleslSI-(101)]
	_ = x[LocalesoSO-(102)]
	_ = x[LocalesuID-(103)]
	_ = x[LocalesvSE-(104)]
	_ = x[LocaleswKE-(105)]
	_ = x[LocaleswTZ-(106)]
	_ = x[LocaletaIN-(107)]
	_ = x[LocaletaLK-(108)]
	_ = x[LocaletaSG-(109)]
	_ = x[LocaleteIN-(110)]
	_ = x[LocalethTH-(111)]
	_ = x[LocaletrTR-(112)]
	_ = x[LocaleukUA-(113)]
	_ = x[LocaleurIN-(114)]
	_ = x[LocaleurPK-(115)]
	_ = x[LocaleuzUZ-(116)]
	_ = x[LocaleviVN-(117)]
	_ = x[LocalezhCN-(118)]
	_ = x[LocalezhHK-(119)]
	_ = x[LocalezhTW-(120)]
	_ = x[LocalezuZA-(121)]
	_ = x[LocaleknIN-(122)]
	_ = x[LocaleloLA-(123)]
}

var _LocaleValues = []Locale{LocaleafZA, LocaleamET, LocalearAE, LocalearBH, LocalearDZ, LocalearEG, LocalearIQ, LocalearJO, LocalearKW, LocalearLY, LocalearMA, LocalearQA, LocalearSA, LocalearSY, LocalearTN, LocalearYE, LocalebgBG, LocalebnBD, LocalebnIN, LocalecaES, LocalecsCZ, LocalecyGB, LocaledaDK, LocaledeAT, LocaledeCH, LocaledeDE, LocaleelGR, LocaleenAU, LocaleenCA, LocaleenGB, LocaleenHK, LocaleenIE, LocaleenIN, LocaleenKE, LocaleenNG, LocaleenNZ, LocaleenPH, LocaleenSG, LocaleenTZ, LocaleenUS, LocaleenZA, LocaleesAR, LocaleesBO, LocaleesCL, LocaleesCO, LocaleesCR, LocaleesCU, LocaleesDO, LocaleesEC, LocaleesES, LocaleesGQ, LocaleesGT, LocaleesHN, LocaleesMX, LocaleesNI, LocaleesPA, LocaleesPE, LocaleesPR, LocaleesPY, LocaleesSV, LocaleesUS, LocaleesUY, LocaleesVE, LocaleetEE, LocalefaIR, LocalefiFI, LocalefilPH, LocalefrBE, LocalefrCA, LocalefrCH, LocalefrFR, LocalegaIE, LocaleglES, LocaleguIN, LocaleheIL, LocalehiIN, LocalehrHR, LocalehuHU, LocaleidID, LocaleisIS, LocaleitIT, LocalejaJP, LocalejvID, LocalekmKH, LocalekoKR, LocaleltLT, LocalelvLV, LocalemrIN, LocalemsMY, LocalemtMT, LocalemyMM, LocalenbNO, LocalenlBE, LocalekkKZ, LocalenlNL, LocaleplPL, LocaleptBR, LocaleptPT, LocaleroRO, LocaleruRU, LocaleskSK, LocaleslSI, LocalesoSO, LocalesuID, LocalesvSE, LocaleswKE, LocaleswTZ, LocaletaIN, LocaletaLK, LocaletaSG, LocaleteIN, LocalethTH, LocaletrTR, LocaleukUA, LocaleurIN, LocaleurPK, LocaleuzUZ, LocaleviVN, LocalezhCN, LocalezhHK, LocalezhTW, LocalezuZA, LocaleknIN, LocaleloLA}

var _LocaleNameToValueMap = map[string]Locale{
	_LocaleName[0:5]:          LocaleafZA,
	_LocaleLowerName[0:5]:     LocaleafZA,
	_LocaleName[5:10]:         LocaleamET,
	_LocaleLowerName[5:10]:    LocaleamET,
	_LocaleName[10:15]:        LocalearAE,
	_LocaleLowerName[10:15]:   LocalearAE,
	_LocaleName[15:20]:        LocalearBH,
	_LocaleLowerName[15:20]:   LocalearBH,
	_LocaleName[20:25]:        LocalearDZ,
	_LocaleLowerName[20:25]:   LocalearDZ,
	_LocaleName[25:30]:        LocalearEG,
	_LocaleLowerName[25:30]:   LocalearEG,
	_LocaleName[30:35]:        LocalearIQ,
	_LocaleLowerName[30:35]:   LocalearIQ,
	_LocaleName[35:40]:        LocalearJO,
	_LocaleLowerName[35:40]:   LocalearJO,
	_LocaleName[40:45]:        LocalearKW,
	_LocaleLowerName[40:45]:   LocalearKW,
	_LocaleName[45:50]:        LocalearLY,
	_LocaleLowerName[45:50]:   LocalearLY,
	_LocaleName[50:55]:        LocalearMA,
	_LocaleLowerName[50:55]:   LocalearMA,
	_LocaleName[55:60]:        LocalearQA,
	_LocaleLowerName[55:60]:   LocalearQA,
	_LocaleName[60:65]:        LocalearSA,
	_LocaleLowerName[60:65]:   LocalearSA,
	_LocaleName[65:70]:        LocalearSY,
	_LocaleLowerName[65:70]:   LocalearSY,
	_LocaleName[70:75]:        LocalearTN,
	_LocaleLowerName[70:75]:   LocalearTN,
	_LocaleName[75:80]:        LocalearYE,
	_LocaleLowerName[75:80]:   LocalearYE,
	_LocaleName[80:85]:        LocalebgBG,
	_LocaleLowerName[80:85]:   LocalebgBG,
	_LocaleName[85:90]:        LocalebnBD,
	_LocaleLowerName[85:90]:   LocalebnBD,
	_LocaleName[90:95]:        LocalebnIN,
	_LocaleLowerName[90:95]:   LocalebnIN,
	_LocaleName[95:100]:       LocalecaES,
	_LocaleLowerName[95:100]:  LocalecaES,
	_LocaleName[100:105]:      LocalecsCZ,
	_LocaleLowerName[100:105]: LocalecsCZ,
	_LocaleName[105:110]:      LocalecyGB,
	_LocaleLowerName[105:110]: LocalecyGB,
	_LocaleName[110:115]:      LocaledaDK,
	_LocaleLowerName[110:115]: LocaledaDK,
	_LocaleName[115:120]:      LocaledeAT,
	_LocaleLowerName[115:120]: LocaledeAT,
	_LocaleName[120:125]:      LocaledeCH,
	_LocaleLowerName[120:125]: LocaledeCH,
	_LocaleName[125:130]:      LocaledeDE,
	_LocaleLowerName[125:130]: LocaledeDE,
	_LocaleName[130:135]:      LocaleelGR,
	_LocaleLowerName[130:135]: LocaleelGR,
	_LocaleName[135:140]:      LocaleenAU,
	_LocaleLowerName[135:140]: LocaleenAU,
	_LocaleName[140:145]:      LocaleenCA,
	_LocaleLowerName[140:145]: LocaleenCA,
	_LocaleName[145:150]:      LocaleenGB,
	_LocaleLowerName[145:150]: LocaleenGB,
	_LocaleName[150:155]:      LocaleenHK,
	_LocaleLowerName[150:155]: LocaleenHK,
	_LocaleName[155:160]:      LocaleenIE,
	_LocaleLowerName[155:160]: LocaleenIE,
	_LocaleName[160:165]:      LocaleenIN,
	_LocaleLowerName[160:165]: LocaleenIN,
	_LocaleName[165:170]:      LocaleenKE,
	_LocaleLowerName[165:170]: LocaleenKE,
	_LocaleName[170:175]:      LocaleenNG,
	_LocaleLowerName[170:175]: LocaleenNG,
	_LocaleName[175:180]:      LocaleenNZ,
	_LocaleLowerName[175:180]: LocaleenNZ,
	_LocaleName[180:185]:      LocaleenPH,
	_LocaleLowerName[180:185]: LocaleenPH,
	_LocaleName[185:190]:      LocaleenSG,
	_LocaleLowerName[185:190]: LocaleenSG,
	_LocaleName[190:195]:      LocaleenTZ,
	_LocaleLowerName[190:195]: LocaleenTZ,
	_LocaleName[195:200]:      LocaleenUS,
	_LocaleLowerName[195:200]: LocaleenUS,
	_LocaleName[200:205]:      LocaleenZA,
	_LocaleLowerName[200:205]: LocaleenZA,
	_LocaleName[205:210]:      LocaleesAR,
	_LocaleLowerName[205:210]: LocaleesAR,
	_LocaleName[210:215]:      LocaleesBO,
	_LocaleLowerName[210:215]: LocaleesBO,
	_LocaleName[215:220]:      LocaleesCL,
	_LocaleLowerName[215:220]: LocaleesCL,
	_LocaleName[220:225]:      LocaleesCO,
	_LocaleLowerName[220:225]: LocaleesCO,
	_LocaleName[225:230]:      LocaleesCR,
	_LocaleLowerName[225:230]: LocaleesCR,
	_LocaleName[230:235]:      LocaleesCU,
	_LocaleLowerName[230:235]: LocaleesCU,
	_LocaleName[235:240]:      LocaleesDO,
	_LocaleLowerName[235:240]: LocaleesDO,
	_LocaleName[240:245]:      LocaleesEC,
	_LocaleLowerName[240:245]: LocaleesEC,
	_LocaleName[245:250]:      LocaleesES,
	_LocaleLowerName[245:250]: LocaleesES,
	_LocaleName[250:255]:      LocaleesGQ,
	_LocaleLowerName[250:255]: LocaleesGQ,
	_LocaleName[255:260]:      LocaleesGT,
	_LocaleLowerName[255:260]: LocaleesGT,
	_LocaleName[260:265]:      LocaleesHN,
	_LocaleLowerName[260:265]: LocaleesHN,
	_LocaleName[265:270]:      LocaleesMX,
	_LocaleLowerName[265:270]: LocaleesMX,
	_LocaleName[270:275]:      LocaleesNI,
	_LocaleLowerName[270:275]: LocaleesNI,
	_LocaleName[275:280]:      LocaleesPA,
	_LocaleLowerName[275:280]: LocaleesPA,
	_LocaleName[280:285]:      LocaleesPE,
	_LocaleLowerName[280:285]: LocaleesPE,
	_LocaleName[285:290]:      LocaleesPR,
	_LocaleLowerName[285:290]: LocaleesPR,
	_LocaleName[290:295]:      LocaleesPY,
	_LocaleLowerName[290:295]: LocaleesPY,
	_LocaleName[295:300]:      LocaleesSV,
	_LocaleLowerName[295:300]: LocaleesSV,
	_LocaleName[300:305]:      LocaleesUS,
	_LocaleLowerName[300:305]: LocaleesUS,
	_LocaleName[305:310]:      LocaleesUY,
	_LocaleLowerName[305:310]: LocaleesUY,
	_LocaleName[310:315]:      LocaleesVE,
	_LocaleLowerName[310:315]: LocaleesVE,
	_LocaleName[315:320]:      LocaleetEE,
	_LocaleLowerName[315:320]: LocaleetEE,
	_LocaleName[320:325]:      LocalefaIR,
	_LocaleLowerName[320:325]: LocalefaIR,
	_LocaleName[325:330]:      LocalefiFI,
	_LocaleLowerName[325:330]: LocalefiFI,
	_LocaleName[330:336]:      LocalefilPH,
	_LocaleLowerName[330:336]: LocalefilPH,
	_LocaleName[336:341]:      LocalefrBE,
	_LocaleLowerName[336:341]: LocalefrBE,
	_LocaleName[341:346]:      LocalefrCA,
	_LocaleLowerName[341:346]: LocalefrCA,
	_LocaleName[346:351]:      LocalefrCH,
	_LocaleLowerName[346:351]: LocalefrCH,
	_LocaleName[351:356]:      LocalefrFR,
	_LocaleLowerName[351:356]: LocalefrFR,
	_LocaleName[356:361]:      LocalegaIE,
	_LocaleLowerName[356:361]: LocalegaIE,
	_LocaleName[361:366]:      LocaleglES,
	_LocaleLowerName[361:366]: LocaleglES,
	_LocaleName[366:371]:      LocaleguIN,
	_LocaleLowerName[366:371]: LocaleguIN,
	_LocaleName[371:376]:      LocaleheIL,
	_LocaleLowerName[371:376]: LocaleheIL,
	_LocaleName[376:381]:      LocalehiIN,
	_LocaleLowerName[376:381]: LocalehiIN,
	_LocaleName[381:386]:      LocalehrHR,
	_LocaleLowerName[381:386]: LocalehrHR,
	_LocaleName[386:391]:      LocalehuHU,
	_LocaleLowerName[386:391]: LocalehuHU,
	_LocaleName[391:396]:      LocaleidID,
	_LocaleLowerName[391:396]: LocaleidID,
	_LocaleName[396:401]:      LocaleisIS,
	_LocaleLowerName[396:401]: LocaleisIS,
	_LocaleName[401:406]:      LocaleitIT,
	_LocaleLowerName[401:406]: LocaleitIT,
	_LocaleName[406:411]:      LocalejaJP,
	_LocaleLowerName[406:411]: LocalejaJP,
	_LocaleName[411:416]:      LocalejvID,
	_LocaleLowerName[411:416]: LocalejvID,
	_LocaleName[416:421]:      LocalekmKH,
	_LocaleLowerName[416:421]: LocalekmKH,
	_LocaleName[421:426]:      LocalekoKR,
	_LocaleLowerName[421:426]: LocalekoKR,
	_LocaleName[426:431]:      LocaleltLT,
	_LocaleLowerName[426:431]: LocaleltLT,
	_LocaleName[431:436]:      LocalelvLV,
	_LocaleLowerName[431:436]: LocalelvLV,
	_LocaleName[436:441]:      LocalemrIN,
	_LocaleLowerName[436:441]: LocalemrIN,
	_LocaleName[441:446]:      LocalemsMY,
	_LocaleLowerName[441:446]: LocalemsMY,
	_LocaleName[446:451]:      LocalemtMT,
	_LocaleLowerName[446:451]: LocalemtMT,
	_LocaleName[451:456]:      LocalemyMM,
	_LocaleLowerName[451:456]: LocalemyMM,
	_LocaleName[456:461]:      LocalenbNO,
	_LocaleLowerName[456:461]: LocalenbNO,
	_LocaleName[461:466]:      LocalenlBE,
	_LocaleLowerName[461:466]: LocalenlBE,
	_LocaleName[466:471]:      LocalekkKZ,
	_LocaleLowerName[466:471]: LocalekkKZ,
	_LocaleName[471:476]:      LocalenlNL,
	_LocaleLowerName[471:476]: LocalenlNL,
	_LocaleName[476:481]:      LocaleplPL,
	_LocaleLowerName[476:481]: LocaleplPL,
	_LocaleName[481:486]:      LocaleptBR,
	_LocaleLowerName[481:486]: LocaleptBR,
	_LocaleName[486:491]:      LocaleptPT,
	_LocaleLowerName[486:491]: LocaleptPT,
	_LocaleName[491:496]:      LocaleroRO,
	_LocaleLowerName[491:496]: LocaleroRO,
	_LocaleName[496:501]:      LocaleruRU,
	_LocaleLowerName[496:501]: LocaleruRU,
	_LocaleName[501:506]:      LocaleskSK,
	_LocaleLowerName[501:506]: LocaleskSK,
	_LocaleName[506:511]:      LocaleslSI,
	_LocaleLowerName[506:511]: LocaleslSI,
	_LocaleName[511:516]:      LocalesoSO,
	_LocaleLowerName[511:516]: LocalesoSO,
	_LocaleName[516:521]:      LocalesuID,
	_LocaleLowerName[516:521]: LocalesuID,
	_LocaleName[521:526]:      LocalesvSE,
	_LocaleLowerName[521:526]: LocalesvSE,
	_LocaleName[526:531]:      LocaleswKE,
	_LocaleLowerName[526:531]: LocaleswKE,
	_LocaleName[531:536]:      LocaleswTZ,
	_LocaleLowerName[531:536]: LocaleswTZ,
	_LocaleName[536:541]:      LocaletaIN,
	_LocaleLowerName[536:541]: LocaletaIN,
	_LocaleName[541:546]:      LocaletaLK,
	_LocaleLowerName[541:546]: LocaletaLK,
	_LocaleName[546:551]:      LocaletaSG,
	_LocaleLowerName[546:551]: LocaletaSG,
	_LocaleName[551:556]:      LocaleteIN,
	_LocaleLowerName[551:556]: LocaleteIN,
	_LocaleName[556:561]:      LocalethTH,
	_LocaleLowerName[556:561]: LocalethTH,
	_LocaleName[561:566]:      LocaletrTR,
	_LocaleLowerName[561:566]: LocaletrTR,
	_LocaleName[566:571]:      LocaleukUA,
	_LocaleLowerName[566:571]: LocaleukUA,
	_LocaleName[571:576]:      LocaleurIN,
	_LocaleLowerName[571:576]: LocaleurIN,
	_LocaleName[576:581]:      LocaleurPK,
	_LocaleLowerName[576:581]: LocaleurPK,
	_LocaleName[581:586]:      LocaleuzUZ,
	_LocaleLowerName[581:586]: LocaleuzUZ,
	_LocaleName[586:591]:      LocaleviVN,
	_LocaleLowerName[586:591]: LocaleviVN,
	_LocaleName[591:596]:      LocalezhCN,
	_LocaleLowerName[591:596]: LocalezhCN,
	_LocaleName[596:601]:      LocalezhHK,
	_LocaleLowerName[596:601]: LocalezhHK,
	_LocaleName[601:606]:      LocalezhTW,
	_LocaleLowerName[601:606]: LocalezhTW,
	_LocaleName[606:611]:      LocalezuZA,
	_LocaleLowerName[606:611]: LocalezuZA,
	_LocaleName[611:616]:      LocaleknIN,
	_LocaleLowerName[611:616]: LocaleknIN,
	_LocaleName[616:621]:      LocaleloLA,
	_LocaleLowerName[616:621]: LocaleloLA,
}

var _LocaleNames = []string{
	_LocaleName[0:5],
	_LocaleName[5:10],
	_LocaleName[10:15],
	_LocaleName[15:20],
	_LocaleName[20:25],
	_LocaleName[25:30],
	_LocaleName[30:35],
	_LocaleName[35:40],
	_LocaleName[40:45],
	_LocaleName[45:50],
	_LocaleName[50:55],
	_LocaleName[55:60],
	_LocaleName[60:65],
	_LocaleName[65:70],
	_LocaleName[70:75],
	_LocaleName[75:80],
	_LocaleName[80:85],
	_LocaleName[85:90],
	_LocaleName[90:95],
	_LocaleName[95:100],
	_LocaleName[100:105],
	_LocaleName[105:110],
	_LocaleName[110:115],
	_LocaleName[115:120],
	_LocaleName[120:125],
	_LocaleName[125:130],
	_LocaleName[130:135],
	_LocaleName[135:140],
	_LocaleName[140:145],
	_LocaleName[145:150],
	_LocaleName[150:155],
	_LocaleName[155:160],
	_LocaleName[160:165],
	_LocaleName[165:170],
	_LocaleName[170:175],
	_LocaleName[175:180],
	_LocaleName[180:185],
	_LocaleName[185:190],
	_LocaleName[190:195],
	_LocaleName[195:200],
	_LocaleName[200:205],
	_LocaleName[205:210],
	_LocaleName[210:215],
	_LocaleName[215:220],
	_LocaleName[220:225],
	_LocaleName[225:230],
	_LocaleName[230:235],
	_LocaleName[235:240],
	_LocaleName[240:245],
	_LocaleName[245:250],
	_LocaleName[250:255],
	_LocaleName[255:260],
	_LocaleName[260:265],
	_LocaleName[265:270],
	_LocaleName[270:275],
	_LocaleName[275:280],
	_LocaleName[280:285],
	_LocaleName[285:290],
	_LocaleName[290:295],
	_LocaleName[295:300],
	_LocaleName[300:305],
	_LocaleName[305:310],
	_LocaleName[310:315],
	_LocaleName[315:320],
	_LocaleName[320:325],
	_LocaleName[325:330],
	_LocaleName[330:336],
	_LocaleName[336:341],
	_LocaleName[341:346],
	_LocaleName[346:351],
	_LocaleName[351:356],
	_LocaleName[356:361],
	_LocaleName[361:366],
	_LocaleName[366:371],
	_LocaleName[371:376],
	_LocaleName[376:381],
	_LocaleName[381:386],
	_LocaleName[386:391],
	_LocaleName[391:396],
	_LocaleName[396:401],
	_LocaleName[401:406],
	_LocaleName[406:411],
	_LocaleName[411:416],
	_LocaleName[416:421],
	_LocaleName[421:426],
	_LocaleName[426:431],
	_LocaleName[431:436],
	_LocaleName[436:441],
	_LocaleName[441:446],
	_LocaleName[446:451],
	_LocaleName[451:456],
	_LocaleName[456:461],
	_LocaleName[461:466],
	_LocaleName[466:471],
	_LocaleName[471:476],
	_LocaleName[476:481],
	_LocaleName[481:486],
	_LocaleName[486:491],
	_LocaleName[491:496],
	_LocaleName[496:501],
	_LocaleName[501:506],
	_LocaleName[506:511],
	_LocaleName[511:516],
	_LocaleName[516:521],
	_LocaleName[521:526],
	_LocaleName[526:531],
	_LocaleName[531:536],
	_LocaleName[536:541],
	_LocaleName[541:546],
	_LocaleName[546:551],
	_LocaleName[551:556],
	_LocaleName[556:561],
	_LocaleName[561:566],
	_LocaleName[566:571],
	_LocaleName[571:576],
	_LocaleName[576:581],
	_LocaleName[581:586],
	_LocaleName[586:591],
	_LocaleName[591:596],
	_LocaleName[596:601],
	_LocaleName[601:606],
	_LocaleName[606:611],
	_LocaleName[611:616],
	_LocaleName[616:621],
}

// LocaleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func LocaleString(s string) (Locale, error) {
	if val, ok := _LocaleNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _LocaleNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Locale values", s)
}

// LocaleValues returns all values of the enum
func LocaleValues() []Locale {
	return _LocaleValues
}

// LocaleStrings returns a slice of all String values of the enum
func LocaleStrings() []string {
	strs := make([]string, len(_LocaleNames))
	copy(strs, _LocaleNames)
	return strs
}

// IsALocale returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Locale) IsALocale() bool {
	for _, v := range _LocaleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Locale
func (i Locale) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Locale
func (i *Locale) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Locale should be a string, got %s", data)
	}

	var err error
	*i, err = LocaleString(s)
	return err
}
