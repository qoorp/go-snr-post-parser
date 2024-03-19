package snrpost

// StartPost is the first line of every notification file
type StartPost struct {
	SekvNR      string      `snr:"8"`
	UtskriftDat string      `snr:"8"`
	KundNR      string      `snr:"10"`
	UtskriftTid string      `snr:"8"`
	_           interface{} `snr:"32"`
	PostTyp100  string      `snr:"3"`
	_           interface{} `snr:"331"`
}

// FirPost (7xx) ...
type FirPost struct {
	PostBeteck  string      `snr:"6" json:"post_beteck"`
	UtskriftDat string      `snr:"8" json:"utskrift_dat"`
	RegDat      string      `snr:"8" json:"reg_dat"`
	IDNR        string      `snr:"17" json:"id_nr"`
	DinrDag     string      `snr:"8" json:"dinr_dag"`
	DinrDagMapp string      `snr:"4" json:"dinr_dag_mapp"`
	Rattelse    string      `snr:"1" json:"rattelse"`
	Status      string      `snr:"3" json:"status"`
	_           interface{} `snr:"11"`
	PostTyp     string      `snr:"3" json:"post_typ"`
}

// AviserPost ...
type AviserPost struct {
	PostBeteck   string `snr:"6" json:"post_beteck"`
	UtskriftDat  string `snr:"8" json:"utskrift_dat"`
	ErendeRegDat string `snr:"16" json:"erende_reg_dat"`
	ObjTyp       string `snr:"5" json:"obj_typ"`
	Sekel        string `snr:"1" json:"sekel"`
	OrgNR        string `snr:"10" json:"org_nr"`
	LopNR        string `snr:"5" json:"lop_nr"`
	Dinr         string `snr:"7" json:"dinr"`
	DinrAr       string `snr:"4" json:"dinr_ar"`
	ErendeTyp    string `snr:"3" json:"arende_typ"`
	HistMark     string `snr:"1" json:"hist_mark"`
	PostTyp      string `snr:"3" json:"post_typ"`
}

// Post800 ...
type Post800 struct {
	Firma        string `snr:"200" json:"firma"`
	FirmaRegDat  string `snr:"8" json:"firma_reg_dat"`
	Lagerbolag   string `snr:"1" json:"lagerbolag"`
	FirmaMark    string `snr:"1" json:"firma_mark"`
	KonvMark     string `snr:"5" json:"konv_mark"`
	RegLan       string `snr:"2" json:"reg_lan"`
	SkyddslanGRP [21]struct {
		Skyddslan     string `snr:"2" json:"skyddslan"`
		SkyddslanMark string `snr:"1" json:"skyddslan_mark"`
	}
	_ interface{} `snr:"51"`
}

// Post808 ...
type Post808 struct {
	Nybemdat     string      `snr:"8" json:"nybemdat"`
	Nyforebolst  string      `snr:"4" json:"nyforebolst"`
	Nyannatdat   string      `snr:"8" json:"Nyannatdat"`
	Nyemmark     string      `snr:"1" json:"nyemmark"`
	Bemskmark    string      `snr:"1" json:"bemskmark"`
	Skkonvtmark  string      `snr:"1" json:"skkonvtmark"`
	Skoptionmark string      `snr:"1" json:"skoptionmark"`
	Skbemdat     string      `snr:"8" json:"skbemdat"`
	Skforebolst  string      `snr:"4" json:"skforebolst"`
	Skannatdat   string      `snr:"8" json:"skannatdat"`
	Bemregdat    string      `snr:"8" json:"bemregdat"`
	_            interface{} `snr:"279"`
}

// Post810 ...
type Post810 struct {
	Co             string      `snr:"50" json:"co"`
	Gata           string      `snr:"32" json:"gata"`
	Postnr         string      `snr:"6" json:"postnr"`
	Padr           string      `snr:"32" json:"padr"`
	Landkod        string      `snr:"5" json:"landkod"`
	Land           string      `snr:"32" json:"land"`
	Lan            string      `snr:"2" json:"lan"`
	Kommun         string      `snr:"2" json:"kommun"`
	Regdat         string      `snr:"8" json:"regdat"`
	Rakperf        string      `snr:"8" json:"rakperf"`
	Rakpert        string      `snr:"8" json:"rakpert"`
	Rakfrom        string      `snr:"4" json:"rakfrom"`
	Raktom         string      `snr:"4" json:"raktom"`
	Adrmark        string      `snr:"1" json:"adrmark"`
	Satemark       string      `snr:"1" json:"satemark"`
	Rakper         string      `snr:"1" json:"rakper"`
	Priv           string      `snr:"1" json:"priv"`
	Publ           string      `snr:"1" json:"publ"`
	Skade          string      `snr:"1" json:"skade"`
	Liv            string      `snr:"1" json:"liv"`
	Bildatdat      string      `snr:"8" json:"bildatdat"`
	Ordenstadgregl string      `snr:"8" json:"ordenstadgregl"`
	Stadfast       string      `snr:"8" json:"stadfast"`
	Stadfav        string      `snr:"2" json:"stadfav"`
	Oktrojater     string      `snr:"8" json:"oktrojater"`
	Forloktroj     string      `snr:"8" json:"forloktroj"`
	Koncforv       string      `snr:"8" json:"koncforv"`
	Sateort        string      `snr:"25" json:"sateort"`
	Tillstdat      string      `snr:"8" json:"tillstdat"`
	Tillsdat1      string      `snr:"8" json:"tillsdat1"`
	Tillstyp1      string      `snr:"5" json:"tillstyp1"`
	Tillsaterk1    string      `snr:"1" json:"tillsaterk1"`
	Tillsdat2      string      `snr:"8" json:"tillsdat2"`
	Tillstyp2      string      `snr:"5" json:"tillstyp2"`
	Tillsaterk2    string      `snr:"1" json:"tillsaterk2"`
	Tillsdat3      string      `snr:"8" json:"tillsdat3"`
	Tillstyp3      string      `snr:"5" json:"tillstyp3"`
	Tillsaterk3    string      `snr:"1" json:"tillsaterk3"`
	_              interface{} `snr:"6"`
}

// Post811_170211 ... Version 1.0
// As of 2023-10-14 this is still valid
type Post811_170211 struct {
	Akap      string      `snr:"12,6" json:"akap"`
	Valuta1   string      `snr:"3" json:"valuta1"`
	Akapl     string      `snr:"12,6" json:"akapl"`
	Valuta2   string      `snr:"3" json:"valuta2"`
	Akaph     string      `snr:"12,6" json:"akaph"`
	Valuta3   string      `snr:"3" json:"valuta3"`
	_         interface{} `snr:"18"`
	Totantakt string      `snr:"18" json:"totantakt"`
	Antaktl   string      `snr:"18" json:"antaktl"`
	Antakth   string      `snr:"18" json:"antakth"`
	_         interface{} `snr:"196"`
}

// Post811_161008 ... Version 4.9
type Post811_161008 struct {
	Akap      string      `snr:"12,6" json:"akap"`
	Valuta1   string      `snr:"3" json:"valuta1"`
	Akapl     string      `snr:"12,6" json:"akapl"`
	Valuta2   string      `snr:"3" json:"valuta2"`
	Akaph     string      `snr:"12,6" json:"akaph"`
	Valuta3   string      `snr:"3" json:"valuta3"`
	_         interface{} `snr:"18"`
	Totantakt string      `snr:"18" json:"totantakt"`
	Antaktl   string      `snr:"18" json:"antaktl"`
	Antakth   string      `snr:"18" json:"antakth"`
	_         interface{} `snr:"194"`
}

// Post811_151205 ... Version 4.7
// This is valid from 151205, but could also work for earlier versions, double check the specs
type Post811_151205 struct {
	Akap      string      `snr:"12,6" json:"akap"`
	Valuta1   string      `snr:"3" json:"valuta1"`
	Akapl     string      `snr:"16" json:"akapl"`
	Valuta2   string      `snr:"3" json:"valuta2"`
	Akaph     string      `snr:"16" json:"akaph"`
	Valuta3   string      `snr:"3" json:"valuta3"`
	_         interface{} `snr:"18"`
	Totantakt string      `snr:"18" json:"totantakt"`
	Antaktl   string      `snr:"18" json:"antaktl"`
	Antakth   string      `snr:"18" json:"antakth"`
	_         interface{} `snr:"200"`
}

// Post812 ...
type Post812 struct {
	Aktieslag string      `snr:"30" json:"aktieslag"`
	Antslag   string      `snr:"18" json:"antslag"`
	Rostvarde string      `snr:"10" json:"rostvarde"`
	Antaktla  string      `snr:"18" json:"antaktla"`
	Antaktho  string      `snr:"18" json:"antaktho"`
	_         interface{} `snr:"237"`
}

// Post813 ...
type Post813 struct {
	Akrad string      `snr:"320" json:"akrad"`
	_     interface{} `snr:"11"`
}

// Post814_161008 ...Version 4.9
// As of 2023-10-14 this is still valid
type Post814_161008 struct {
	Skvaluta    string      `snr:"3" json:"skvaluta"`
	Skbesldat   string      `snr:"8" json:"skbesldat"`
	Skmark      string      `snr:"1" json:"skmark"`
	Skbeslbel   string      `snr:"12,6" json:"skbeslbel"`
	Sklagst     string      `snr:"12,6" json:"sklagst"`
	Skhogst     string      `snr:"12,6" json:"skhogst"`
	Skteknatbel string      `snr:"12,6" json:"skteknatbel"`
	Skkonvbel   string      `snr:"12,6" json:"skkonvbel"`
	Sktidutb1   string      `snr:"8" json:"sktidutb1"`
	Sktidutb2   string      `snr:"8" json:"sktidutb2"`
	Skoptionbel string      `snr:"12,6" json:"skoptionbel"`
	Sktidnyt1   string      `snr:"8" json:"sktidnyt1"`
	Sktidnyt2   string      `snr:"8" json:"sktidnyt2"`
	Skapapokn   string      `snr:"12,6" json:"skapapokn"`
	Skaktieslag string      `snr:"40" json:"skaktieslag"`
	Toptant     string      `snr:"12" json:"toptant"`
	Anttoptl    string      `snr:"16" json:"anttoptl"`
	Anttopth    string      `snr:"16" json:"anttopth"`
	Toptfrom    string      `snr:"8" json:"toptfrom"`
	Topttom     string      `snr:"8" json:"topttom"`
	Toptml      string      `snr:"1" json:"toptml"`
	Toptul      string      `snr:"1" json:"toptul"`
	Sktyp       string      `snr:"1" json:"sktyp"`
	Skdelreg    string      `snr:"1" json:"skdelreg"`
	_           interface{} `snr:"57"`
}

// Post814_151205 ... Version 4.7
// This is valid from 151205, but could also work for earlier versions, double check the specs
type Post814_151205 struct {
	Skvaluta    string      `snr:"3" json:"skvaluta"`
	Skbesldat   string      `snr:"8" json:"skbesldat"`
	Skmark      string      `snr:"1" json:"skmark"`
	Skbeslbel   string      `snr:"12,6" json:"skbeslbel"`
	Sklagst     string      `snr:"16" json:"sklagst"`
	Skhogst     string      `snr:"16" json:"skhogst"`
	Skteknatbel string      `snr:"12,6" json:"skteknatbel"`
	Skkonvbel   string      `snr:"12,6" json:"skkonvbel"`
	Sktidutb1   string      `snr:"8" json:"sktidutb1"`
	Sktidutb2   string      `snr:"8" json:"sktidutb2"`
	Skoptionbel string      `snr:"12,6" json:"skoptionbel"`
	Sktidnyt1   string      `snr:"8" json:"sktidnyt1"`
	Sktidnyt2   string      `snr:"8" json:"sktidnyt2"`
	Skapapokn   string      `snr:"12,6" json:"skapapokn"`
	Skaktieslag string      `snr:"40" json:"skaktieslag"`
	Toptant     string      `snr:"12" json:"toptant"`
	Anttoptl    string      `snr:"16" json:"anttoptl"`
	Anttopth    string      `snr:"16" json:"anttopth"`
	Toptfrom    string      `snr:"8" json:"toptfrom"`
	Topttom     string      `snr:"8" json:"topttom"`
	Toptml      string      `snr:"1" json:"toptml"`
	Toptul      string      `snr:"1" json:"toptul"`
	Sktyp       string      `snr:"1" json:"sktyp"`
	Skdelreg    string      `snr:"1" json:"skdelreg"`
	_           interface{} `snr:"61"`
}

// Post815 ...
type Post815 struct {
	Skrad string      `snr:"320" json:"skrad"`
	_     interface{} `snr:"11"`
}

// Post816 ...
type Post816 struct {
	Nytutbvaluta          string      `snr:"3" json:"nytutbvaluta"`
	Nytutbmark            string      `snr:"1" json:"nytutbmark"`
	Nytutbbel             string      `snr:"12" json:"nytutbbel"`
	NytutbbelPrecision    string      `snr:"6" json:"nytutbbel_precision"`
	Nytutbnombel          string      `snr:"12" json:"nytutbnombel"`
	NytutbnombelPrecision string      `snr:"6" json:"nytutbnombel_precision"`
	Nytutbantakt          string      `snr:"18" json:"nytutbantakt"`
	_                     interface{} `snr:"273"`
}

// Post817 ...
type Post817 struct {
	Utbnytrad string      `snr:"320" json:"utbnytrad"`
	_         interface{} `snr:"11"`
}

// Post820 ...
type Post820 struct {
	Ledl        string      `snr:"5" json:"ledl"`
	Ledh        string      `snr:"5" json:"ledh"`
	Suppl       string      `snr:"5" json:"suppl"`
	Supph       string      `snr:"5" json:"supph"`
	Revl        string      `snr:"5" json:"revl"`
	Revh        string      `snr:"5" json:"revh"`
	Bolmvak     string      `snr:"8" json:"bolmvak"`
	Bolmsakn    string      `snr:"8" json:"bolmsakn"`
	Forestsakn  string      `snr:"8" json:"forestsakn"`
	Kompsakn    string      `snr:"8" json:"kompsakn"`
	Styrsakn    string      `snr:"8" json:"styrsakn"`
	Anm         string      `snr:"8" json:"anm"`
	Ejfult      string      `snr:"8" json:"ejfult"`
	Ejbehor     string      `snr:"8" json:"ejbehor"`
	Vdsakn      string      `snr:"8" json:"vdsakn"`
	Likvsakn    string      `snr:"8" json:"likvsakn"`
	Delgmsakn   string      `snr:"8" json:"delgmsakn"`
	Revsakn     string      `snr:"8" json:"revsakn"`
	Senregdat   string      `snr:"8" json:"senregdat"`
	Vakansfinns string      `snr:"1" json:"vakansfinns"`
	Ledvald     string      `snr:"3" json:"ledvald"`
	Supplvald   string      `snr:"3" json:"supplvald"`
	Styandrat   string      `snr:"8" json:"styandrat"`
	_           interface{} `snr:"182"`
}

// Post830 ...
type Post830 struct {
	Sekf            string      `snr:"1" json:"sekf"`
	Pnr             string      `snr:"10" json:"pnr"`
	Funk1           string      `snr:"5" json:"funk1"`
	Funk2           string      `snr:"5" json:"funk2"`
	Funk3           string      `snr:"5" json:"funk3"`
	Funk4           string      `snr:"5" json:"funk4"`
	Typa            string      `snr:"1" json:"typa"`
	Type            string      `snr:"1" json:"type"`
	Typu            string      `snr:"1" json:"typu"`
	Namn1           string      `snr:"50" json:"namn1"`
	Namn2           string      `snr:"50" json:"namn2"`
	Fco             string      `snr:"50" json:"fco"`
	Fgata           string      `snr:"32" json:"fgata"`
	Fpostnr         string      `snr:"6" json:"fpostnr"`
	Fpadr           string      `snr:"32" json:"fpadr"`
	Flandkod        string      `snr:"5" json:"flandkod"`
	Land            string      `snr:"32" json:"land"`
	Utltillst       string      `snr:"8" json:"utltillst"`
	Perforsekl      string      `snr:"1" json:"perforsekl"`
	Perforpnr       string      `snr:"10" json:"perforpnr"`
	Insats          string      `snr:"16" json:"insats"`
	InsatsPrecision string      `snr:"2" json:"insats_precision"`
	Kval            string      `snr:"1" json:"kval"`
	_               interface{} `snr:"2"`
}

// Post835 ...
type Post835 struct {
	Astsekel string      `snr:"1" json:"astsekel"`
	Astpnr   string      `snr:"10" json:"astpnr"`
	Astftyp  string      `snr:"1" json:"astftyp"`
	Astrad   string      `snr:"50" json:"astrad"`
	_        interface{} `snr:"269"`
}

// Post840 ...
type Post840 struct {
	NamnTyp      string `snr:"2"`
	Namn         string `snr:"200"`
	NamnRegDat   string `snr:"8"`
	BeslStamma   string `snr:"1"`
	BeslStadg    string `snr:"1"`
	SkyddslanGRP [21]struct {
		Skyddslan     string `snr:"2" json:"skyddslan"`
		SkyddslanMark string `snr:"1" json:"skyddslan_mark"`
	}
	_ interface{} `snr:"56"`
}

// Post880 ...
type Post880 struct {
	Ftrad string      `snr:"320" json:"ftrad"`
	_     interface{} `snr:"11"`
}

// Post881 ...
type Post881 struct {
	Fovaluta            string      `snr:"3" json:"fovaluta"`
	Fobesldat           string      `snr:"8" json:"fobesldat"`
	Genom1              string      `snr:"2" json:"genom1"`
	Genom2              string      `snr:"2" json:"genom2"`
	Genom3              string      `snr:"2" json:"genom3"`
	Genom4              string      `snr:"2" json:"genom4"`
	Genom5              string      `snr:"2" json:"genom5"`
	Genom6              string      `snr:"2" json:"genom6"`
	Fondembel           string      `snr:"12" json:"fondembel"`
	FondembelPrecision  string      `snr:"6" json:"fondmbel_precision"`
	Fondnombel          string      `snr:"12" json:"fondnombel"`
	FondnombelPrecision string      `snr:"6" json:"fondnombel_precision"`
	Foantaktier         string      `snr:"18" json:"foantaktiert"`
	Genom7              string      `snr:"2" json:"genom7"`
	_                   interface{} `snr:"252"`
}

// Post882 ...
type Post882 struct {
	Forad string      `snr:"320" json:"forad"`
	_     interface{} `snr:"11"`
}

// Post883_161008 ... Version 4.9
// As of 2023-10-14 this is still valid
type Post883_161008 struct {
	Nyvaluta              string      `snr:"3" json:"nyvaluta"`
	Nybesldat             string      `snr:"8" json:"nybesldat"`
	Nybeslav              string      `snr:"1" json:"nybeslav"`
	Nyejfullbet           string      `snr:"1" json:"nyejfullbet"`
	Nyembel               string      `snr:"12" json:"nyembel"`
	NyembelPrecision      string      `snr:"6" json:"nyembel_precision"`
	Nygrans1              string      `snr:"12" json:"nygrans1"`
	Nygrans1Precision     string      `snr:"6" json:"nygrans1_precision"`
	Nygrans2              string      `snr:"12" json:"nygrans2"`
	Nygrans2Precision     string      `snr:"6" json:"nygrans2_precision"`
	_                     interface{} `snr:"18"` // NYNOMBEL, not used for AB, BAB, FAB
	Nyantaktier           string      `snr:"18" json:"nyantaktier"`
	Nyokurs               string      `snr:"12" json:"nyokurs"`
	NyokursPrecision      string      `snr:"6" json:"nyokurs_precision"`
	Nydelantal            string      `snr:"18" json:"nydelantal"`
	Nydelbel              string      `snr:"12" json:"nydelbel"`
	NydelbelPrecision     string      `snr:"6" json:"nydelbel_precision"`
	Nyfullantal           string      `snr:"18" json:"nyfullantal"`
	Nyfullbet             string      `snr:"12" json:"nyfullbet"`
	NyfullbetPrecision    string      `snr:"6" json:"nyfullbet_precision"`
	Nyejenomfbel          string      `snr:"12" json:"nyejenomfbel"`
	NyejenomfbelPrecision string      `snr:"6" json:"nyejenomfbel_precision"`
	Nyejenomfant          string      `snr:"18" json:"nyejenomfant"`
	Nybetkon              string      `snr:"2" json:"nybetkon"`
	Nybetapp              string      `snr:"2" json:"nybetapp"`
	Nybetkvi              string      `snr:"2" json:"nybetkvi"`
	Nydelreg              string      `snr:"1" json:"nydelreg"`
	_                     interface{} `snr:"95"`
}

// Post883_151205...
// This is valid from 151205, but could also work for earlier versions, double check the specs
type Post883_151205 struct {
	Nyvaluta              string      `snr:"3" json:"nyvaluta"`
	Nybesldat             string      `snr:"8" json:"nybesldat"`
	Nybeslav              string      `snr:"1" json:"nybeslav"`
	Nyejfullbet           string      `snr:"1" json:"nyejfullbet"`
	Nyembel               string      `snr:"12" json:"nyembel"`
	NyembelPrecision      string      `snr:"6" json:"nyembel_precision"`
	Nygrans1              string      `snr:"16" json:"nygrans1"`
	Nygrans2              string      `snr:"16" json:"nygrans2"`
	_                     interface{} `snr:"18"` // NYNOMBEL, not used for AB, BAB, FAB
	Nyantaktier           string      `snr:"18" json:"nyantaktier"`
	Nyokurs               string      `snr:"12" json:"nyokurs"`
	NyokursPrecision      string      `snr:"6" json:"nyokurs_precision"`
	Nydelantal            string      `snr:"18" json:"nydelantal"`
	Nydelbel              string      `snr:"12" json:"nydelbel"`
	NydelbelPrecision     string      `snr:"6" json:"nydelbel_precision"`
	Nyfullantal           string      `snr:"18" json:"nyfullantal"`
	Nyfullbet             string      `snr:"12" json:"nyfullbet"`
	NyfullbetPrecision    string      `snr:"6" json:"nyfullbet_precision"`
	Nyejenomfbel          string      `snr:"12" json:"nyejenomfbel"`
	NyejenomfbelPrecision string      `snr:"6" json:"nyejenomfbel_precision"`
	Nyejenomfant          string      `snr:"18" json:"nyejenomfant"`
	Nybetkon              string      `snr:"2" json:"nybetkon"`
	Nybetapp              string      `snr:"2" json:"nybetapp"`
	Nybetkvi              string      `snr:"2" json:"nybetkvi"`
	Nydelreg              string      `snr:"1" json:"nydelreg"`
	_                     interface{} `snr:"99"`
}

// Post884 ...
type Post884 struct {
	Nyrad string      `snr:"320" json:"forad"`
	_     interface{} `snr:"11"`
}

// Post885 ...
type Post885 struct {
	Porad string      `snr:"320" json:"porad"`
	_     interface{} `snr:"11"`
}

// Post886_230204 ... Version 1.7
// As of 2023-10-14 this is still valid
type Post886_230204 struct {
	Nevaluta     string      `snr:"3" json:"nevaluta"`
	Nebesldat    string      `snr:"8" json:"nebesldat"`
	Ngenom1      string      `snr:"2" json:"ngenom1"`
	Ngenom2      string      `snr:"2" json:"ngenom2"`
	Ngenom3      string      `snr:"2" json:"ngenom3"`
	Ngenom4      string      `snr:"2" json:"ngenom4"`
	Ngenom5      string      `snr:"2" json:"ngenom5"`
	Ngenom6      string      `snr:"2" json:"ngenom6"`
	Nandamal1    string      `snr:"2" json:"nandamal1"`
	Nandamal2    string      `snr:"2" json:"nandamal2"`
	Nandamal3    string      `snr:"2" json:"nandamal3"`
	Nandamal4    string      `snr:"2" json:"nandamal4"`
	Nandamal5    string      `snr:"2" json:"nandamal5"`
	Nandamal6    string      `snr:"2" json:"nandamal6"`
	Nembel       string      `snr:"12,6" json:"nembel"`
	Netotbeslbet string      `snr:"12,6" json:"netotbeslbet"`
	_            interface{} `snr:"18" `
	Neejvarkst   string      `snr:"1" json:"neejvarkst"`
	Neantaktier  string      `snr:"18" json:"neantaktier"`
	Netilstbesl  string      `snr:"8" json:"netilstbesl"`
	Netilstav    string      `snr:"4" json:"netilstav"`
	Netilstdat   string      `snr:"8" json:"netilstdat"`
	Netillsbel   string      `snr:"12,6" json:"netillsbel"`
	Neforfdat    string      `snr:"8" json:"neforfdat"`
	Neforfav     string      `snr:"4" json:"neforfav"`
	Neforfbel    string      `snr:"12,6" json:"neforfbel"`
	Neforfanta   string      `snr:"18" json:"neforfanta"`
	Ngenom7      string      `snr:"2" json:"ngenom7"`
	Ngenom8      string      `snr:"2" json:"ngenom8"`
	Nlagst       string      `snr:"12,6" json:"nlagst"`
	Nhogst       string      `snr:"12,6" json:"nhogst"`
	Beslav       string      `snr:"1" json:"beslav"`
	_            interface{} `snr:"96"`
}

// Post886_170213 ...
type Post886_170213 struct {
	Nevaluta     string      `snr:"3" json:"nevaluta"`
	Nebesldat    string      `snr:"8" json:"nebesldat"`
	Ngenom1      string      `snr:"2" json:"ngenom1"`
	Ngenom2      string      `snr:"2" json:"ngenom2"`
	Ngenom3      string      `snr:"2" json:"ngenom3"`
	Ngenom4      string      `snr:"2" json:"ngenom4"`
	Ngenom5      string      `snr:"2" json:"ngenom5"`
	Ngenom6      string      `snr:"2" json:"ngenom6"`
	Nandamal1    string      `snr:"2" json:"nandamal1"`
	Nandamal2    string      `snr:"2" json:"nandamal2"`
	Nandamal3    string      `snr:"2" json:"nandamal3"`
	Nandamal4    string      `snr:"2" json:"nandamal4"`
	Nandamal5    string      `snr:"2" json:"nandamal5"`
	Nandamal6    string      `snr:"2" json:"nandamal6"`
	Nembel       string      `snr:"12,6" json:"nembel"`
	Netotbeslbet string      `snr:"12,6" json:"netotbeslbet"`
	_            interface{} `snr:"18" `
	Neejvarkst   string      `snr:"1" json:"neejvarkst"`
	Neantaktier  string      `snr:"18" json:"neantaktier"`
	Netilstbesl  string      `snr:"8" json:"netilstbesl"`
	Netilstav    string      `snr:"4" json:"netilstav"`
	Netilstdat   string      `snr:"8" json:"netilstdat"`
	Netillsbel   string      `snr:"12,6" json:"netillsbel"`
	Neforfdat    string      `snr:"8" json:"neforfdat"`
	Neforfav     string      `snr:"4" json:"neforfav"`
	Neforfbel    string      `snr:"12,6" json:"neforfbel"`
	Neforfanta   string      `snr:"18" json:"neforfanta"`
	Ngenom7      string      `snr:"2" json:"ngenom7"`
	Ngenom8      string      `snr:"2" json:"ngenom8"`
	Nlagst       string      `snr:"12,6" json:"nlagst"`
	Nhogst       string      `snr:"12,6" json:"nhogst"`
	_            interface{} `snr:"97"`
}

// Post886_151205...
// Bolagsverket changed the structure sometime in 2017
// This is support for the old format
type Post886_151205 struct {
	Nevaluta     string      `snr:"3" json:"nevaluta"`
	Nebesldat    string      `snr:"8" json:"nebesldat"`
	Ngenom1      string      `snr:"2" json:"ngenom1"`
	Ngenom2      string      `snr:"2" json:"ngenom2"`
	Ngenom3      string      `snr:"2" json:"ngenom3"`
	Ngenom4      string      `snr:"2" json:"ngenom4"`
	Ngenom5      string      `snr:"2" json:"ngenom5"`
	Ngenom6      string      `snr:"2" json:"ngenom6"`
	Nandamal1    string      `snr:"2" json:"nandamal1"`
	Nandamal2    string      `snr:"2" json:"nandamal2"`
	Nandamal3    string      `snr:"2" json:"nandamal3"`
	Nandamal4    string      `snr:"2" json:"nandamal4"`
	Nandamal5    string      `snr:"2" json:"nandamal5"`
	Nandamal6    string      `snr:"2" json:"nandamal6"`
	Nembel       string      `snr:"12,6" json:"nembel"`
	Netotbeslbet string      `snr:"12,6" json:"netotbeslbet"`
	_            interface{} `snr:"18" `
	Neejvarkst   string      `snr:"1" json:"neejvarkst"`
	Neantaktier  string      `snr:"18" json:"neantaktier"`
	Netilstbesl  string      `snr:"8" json:"netilstbesl"`
	Netilstav    string      `snr:"4" json:"netilstav"`
	Netilstdat   string      `snr:"8" json:"netilstdat"`
	Netillsbel   string      `snr:"12,6" json:"netillsbel"`
	Neforfdat    string      `snr:"8" json:"neforfdat"`
	Neforfav     string      `snr:"4" json:"neforfav"`
	Neforfbel    string      `snr:"12,6" json:"neforfbel"`
	Neforfanta   string      `snr:"18" json:"neforfanta"`
	Ngenom7      string      `snr:"2" json:"ngenom7"`
	Ngenom8      string      `snr:"2" json:"ngenom8"`
	Nlagst       string      `snr:"16" json:"nlagst"`
	Nhogst       string      `snr:"16" json:"nhogst"`
	_            interface{} `snr:"101"`
}

// Post887 ...
type Post887 struct {
	Nerad string      `snr:"320" json:"nerad"`
	_     interface{} `snr:"11"`
}

// Post888 ...
type Post888 struct {
	Kalltyp       string      `snr:"10" json:"kalltyp"`
	Kalldinr      string      `snr:"7" json:"kalldinr"`
	Kalldinrar    string      `snr:"4" json:"kalldinrar"`
	Kalldatum     string      `snr:"8" json:"kalldatum"`
	Kallkungdat   string      `snr:"8" json:"kallkungdat"`
	Kallbestrdat  string      `snr:"8" json:"Kallbestrdat"`
	Kallotrdinr   string      `snr:"7" json:"kallotrdinr"`
	Kallotrdinrar string      `snr:"4" json:"kallotrdinrar"`
	Kallotrdat    string      `snr:"8" json:"kallotrdat"`
	Kallotrdom    string      `snr:"3" json:"kallotrdom"`
	_             interface{} `snr:"264"`
}

// Post890 ...
type Post890 struct {
	Vsrad string      `snr:"320" json:"vsrad"`
	_     interface{} `snr:"11"`
}

// Post892 ...
type Post892 struct {
	Karad string      `snr:"320" json:"karad"`
	_     interface{} `snr:"11"`
}

// Post893 ...
type Post893 struct {
	Markforb                string      `snr:"1" json:"markforb"`
	Markomtyp1              string      `snr:"1" json:"markomtyp1"`
	Markomtyp2              string      `snr:"1" json:"markomtyp2"`
	Markomtyp3              string      `snr:"1" json:"markomtyp3"`
	Markomtyp4              string      `snr:"1" json:"markomtyp4"`
	Markomtyp5              string      `snr:"1" json:"markomtyp5"`
	Markomtyp6              string      `snr:"1" json:"markomtyp6"`
	Forbfom                 string      `snr:"8" json:"forbfom"`
	Markomtyp7              string      `snr:"1" json:"markomtyp7"`
	Markomtyp8              string      `snr:"1" json:"markomtyp8"`
	Markomtyp9              string      `snr:"1" json:"markomtyp9"`
	Ejrevbsldat             string      `snr:"8" json:"ejrevbsldat"`
	Omvandlingsforbehall    string      `snr:"1" json:"omvandlingsforbehall"`
	Inlosenforbehall        string      `snr:"1" json:"inlosenforbehall"`
	Likvidationsbestammelse string      `snr:"1" json:"likvidationsbestammelse"`
	Uppkopserbjudande       string      `snr:"1" json:"uppkopserbjudande"`
	_                       interface{} `snr:"301"`
}

// Post894 ...
type Post894 struct {
	Forbrad string      `snr:"320" json:"forbrad"`
	_       interface{} `snr:"11"`
}

// Post895 ...
type Post895 struct {
	Exrad string      `snr:"320" json:"exrad"`
	_     interface{} `snr:"11"`
}

// Post896 ...
type Post896 struct {
	Ovrad string      `snr:"320" json:"ovrad"`
	_     interface{} `snr:"11"`
}

// Post930 ...
type Post930 struct {
	Forsenenper     string      `snr:"8" json:"forsenenper"`
	Forsenregdat    string      `snr:"8" json:"forsenregdat"`
	Forsennr        string      `snr:"2" json:"forsennr"`
	Forbel          string      `snr:"13" json:"forbel"`
	ForbelPrecision string      `snr:"2" json:"forbel_precision"`
	Lagakraft       string      `snr:"8" json:"lagakraft"`
	Avskrivet       string      `snr:"8" json:"avskrivet"`
	Avskranel       string      `snr:"50" json:"avskranel"`
	_               interface{} `snr:"232"`
}

// Post931 ...
type Post931 struct {
	Forflomper string      `snr:"8" json:"forflomper"`
	Forftomper string      `snr:"8" json:"forftomper"`
	Forldat    string      `snr:"8" json:"forldat"`
	Forlregdat string      `snr:"8" json:"forlregdat"`
	_          interface{} `snr:"299"`
}

// Post970 ...
type Post970 struct {
	Isekl      string `snr:"1" json:"isekl"`
	Iorgnummer string `snr:"10" json:"iorgnummer"`
	Ilopnr     string `snr:"3" json:"ilopnr"`
	Idinrar    string `snr:"4" json:"idinrar"`
	Idinr      string `snr:"7" json:"idinr"`
	Inkdat     string `snr:"8" json:"inkdat"`
	Iobjtyp    string `snr:"5" json:"iobjtyp"`
	RubricGRP  [20]struct {
		Rubric string `snr:"10" json:"rubric"`
	}
	_ interface{} `snr:"93"`
}
