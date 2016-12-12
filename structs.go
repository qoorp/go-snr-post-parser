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
}
