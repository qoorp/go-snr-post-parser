package snrpost

// PostID ...
type PostID struct {
	PostBeteck  string      `snr:"6"`
	UtskriftDat string      `snr:"8"`
	RegDat      string      `snr:"8"`
	IDNR        string      `snr:"17"`
	DinrDag     string      `snr:"8"`
	DinrDagMapp string      `snr:"4"`
	Rattelse    string      `snr:"1"`
	Status      string      `snr:"3"`
	_           interface{} `snr:"11"`
	PostTyp     string      `snr:"3"`
}

// Post800 ...
type Post800 struct {
	Firma        string `snr:"200"`
	FirmaRegDat  string `snr:"8"`
	LagerBolag   string `snr:"1"`
	FirmaMark    string `snr:"1"`
	KonvMark     string `snr:"5"`
	RegLan       string `snr:"2"`
	SkyddslanGRP [21]struct {
		Skyddslan     string `snr:"2"`
		SkyddslanMark string `snr:"1"`
	}
}
