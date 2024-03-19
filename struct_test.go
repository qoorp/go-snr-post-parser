package snrpost

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type exampleData struct {
	post string
	v    interface{}
}

func Test_size(t *testing.T) {
	var start StartPost
	size, err := SnrSize(&start)
	if err != nil {
		t.Fatal(err)
	}
	i, err := SnrSize(start)
	if err != nil {
		t.Fatal(err)
	}
	if i != size {
		t.Fatal("expected", size, "got ", i)
	}
	var av AviserPost
	sizeAV, err := SnrSize(av)
	if err != nil {
		t.Fatal(err)
	}
	if sizeAV != postIDLength {
		t.Fatal("expected", postIDLength, "got ", sizeAV)
	}
}

func TestNilUnmarshal(t *testing.T) {
	_, err := unmarshal(nil, nil)
	if err == nil {
		t.Fatal(err)
	}
}

func TestNonPtrUnmarshal(t *testing.T) {
	var p Post800
	_, err := unmarshal(nil, p)
	if err == nil {
		t.Fatal(err)
	}
}

func Test_content(t *testing.T) {
	posts := []string{
		"800",
		"808",
		"810",
		// "811",
		"812",
		"813",
		// "814",
		"815",
		"816",
		"817",
		"820",
		"830",
		"835",
		"840",
		"880",
		"881",
		"882",
		"883",
		"884",
		"885",
		// "886",
		"887",
		"888",
		"892",
		"893",
		"894",
		"895",
		"896",
		"930",
		"931",
		"970",
	}
	for _, post := range posts {
		err := checkContent(post)
		if err != nil {
			t.Error(err)
		}
	}
}

//
// Internal functions used only in this file
//

func checkContent(post string) error {
	v := postPointer(post)
	size, err := SnrSize(v)
	if err != nil {
		return err
	}
	if size != aviserDataLength {
		return fmt.Errorf("expected %v got %v", aviserDataLength, size)
	}
	p, err := postUnmarshal(postData(post))
	if err != nil {
		return err
	}
	return postCompare(p, v)
}

func exampleLines(file string) ([][]byte, error) {
	result := [][]byte{}
	bs, err := os.ReadFile(file)
	if err != nil {
		return result, err
	}
	// If the ZIP file contents arrive in some other way than as a file...
	b := bytes.NewReader(bs)
	r, err := zip.NewReader(b, b.Size())
	if err != nil {
		return result, err
	}
	f := r.File[0]
	rc, err := f.Open()
	if err != nil {
		return result, err
	}
	br := bufio.NewReader(rc)
	for {
		line, partial, err := br.ReadLine()
		if err == nil {
			if partial {
				return result, errors.New("partial")
			}
			tmp := make([]byte, len(line))
			copy(tmp, line)
			result = append(result, tmp)
		} else {
			if err.Error() == "EOF" {
				break
			}
			return result, err
		}
	}
	return result, nil
}

func exampleTestArguments() []exampleData {
	return []exampleData{
		{"800", &Post800{}},
		{"808", &Post808{}},
		{"810", &Post810{}},
		{"811", &Post811_170211{}},
		{"812", &Post812{}},
		{"813", &Post813{}},
		{"814", &Post814_161008{}},
		{"816", &Post816{}},
		{"820", &Post820{}},
		{"830", &Post830{}},
		{"840", &Post840{}},
		{"880", &Post880{}},
		{"881", &Post881{}},
		{"883", &Post883_161008{}},
		{"884", &Post884{}},
		{"885", &Post885{}},
		{"886", &Post886_230204{}},
		{"888", &Post888{}},
		{"892", &Post892{}},
		{"893", &Post893{}},
		{"894", &Post894{}},
		{"895", &Post895{}},
		{"896", &Post896{}},
		{"930", &Post930{}},
		{"931", &Post931{}},
		{"970", &Post970{}},
	}
}

func p100(sizeAV int, line []byte) error {
	var start StartPost
	p100 := "100"
	post, err := post(sizeAV, line)
	if err != nil {
		return err
	}
	if post != p100 {
		return errors.New(p100 + " not found")
	}
	return postCheck(line, &start, p100, postIDLength+aviserDataLength)
}

func pXXX(t *testing.T, sizeAV int, lines [][]byte, v interface{}, pxxx string) {
	size, err := SnrSize(v)
	if err != nil {
		t.Error(pxxx, err)
		return
	}
	if size != aviserDataLength {
		t.Error(pxxx, "expected", aviserDataLength, "got", size)
		return
	}
	found := false
	for _, line := range lines {
		post, err := post(sizeAV, line)
		if err != nil {
			t.Error(pxxx, err)
			return
		}
		if post == pxxx {
			found = true
			err := postCheck(line, v, pxxx, size)
			if err != nil {
				t.Error(pxxx, err)
			}
		}
	}
	if !found {
		t.Error(pxxx, "not found")
	}
}

func post(sizeAV int, line []byte) (string, error) {
	var av AviserPost
	i, err := UnmarshalAviserPost(line, &av)
	if err != nil {
		return "", err
	}
	if i != sizeAV {
		return "", errors.New("wrong size: " + string(line))
	}
	return av.PostTyp, nil
}

func postCheck(line []byte, v interface{}, p8 string, size int) error {
	i, err := UnmarshalData(line, v)
	if err != nil {
		return err
	}
	if i != size {
		return fmt.Errorf("expected %v got %v", size, i)
	}
	return nil
}

func postCompare(got, expected interface{}) (result error) {
	switch got.(type) {
	case *Post800:
		g := got.(*Post800)
		e := expected.(*Post800)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post808:
		g := got.(*Post808)
		e := expected.(*Post808)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post810:
		g := got.(*Post810)
		e := expected.(*Post810)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post811_170211:
		g := got.(*Post811_170211)
		e := expected.(*Post811_170211)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post812:
		g := got.(*Post812)
		e := expected.(*Post812)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post813:
		g := got.(*Post813)
		e := expected.(*Post813)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post814_161008:
		g := got.(*Post814_161008)
		e := expected.(*Post814_161008)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post815:
		g := got.(*Post815)
		e := expected.(*Post815)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post816:
		g := got.(*Post816)
		e := expected.(*Post816)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post817:
		g := got.(*Post817)
		e := expected.(*Post817)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post820:
		g := got.(*Post820)
		e := expected.(*Post820)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post830:
		g := got.(*Post830)
		e := expected.(*Post830)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post835:
		g := got.(*Post835)
		e := expected.(*Post835)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post840:
		g := got.(*Post840)
		e := expected.(*Post840)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post880:
		g := got.(*Post880)
		e := expected.(*Post880)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post881:
		g := got.(*Post881)
		e := expected.(*Post881)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post882:
		g := got.(*Post882)
		e := expected.(*Post882)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post883_161008:
		g := got.(*Post883_161008)
		e := expected.(*Post883_161008)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post884:
		g := got.(*Post884)
		e := expected.(*Post884)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post885:
		g := got.(*Post885)
		e := expected.(*Post885)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post886_230204:
		g := got.(*Post886_230204)
		e := expected.(*Post886_230204)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post887:
		g := got.(*Post887)
		e := expected.(*Post887)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post888:
		g := got.(*Post888)
		e := expected.(*Post888)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post892:
		g := got.(*Post892)
		e := expected.(*Post892)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post893:
		g := got.(*Post893)
		e := expected.(*Post893)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post894:
		g := got.(*Post894)
		e := expected.(*Post894)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post895:
		g := got.(*Post895)
		e := expected.(*Post895)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post896:
		g := got.(*Post896)
		e := expected.(*Post896)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post930:
		g := got.(*Post930)
		e := expected.(*Post930)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post931:
		g := got.(*Post931)
		e := expected.(*Post931)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	case *Post970:
		g := got.(*Post970)
		e := expected.(*Post970)
		if *g != *e {
			result = fmt.Errorf("expected %v got %v", e, g)
		}
	default:
		result = fmt.Errorf("unknown type %v", reflect.TypeOf(got))
	}
	return
}

func postData(post string) []byte {
	result := make([]byte, postIDLength+aviserDataLength)
	switch post {
	case "800":
		postData800(result)
	case "808":
		postData808(result)
	case "810":
		postData810(result)
	case "811":
		postData811(result)
	case "812":
		postData812(result)
	case "813":
		copy(result[postIDLength:], postPointer813().Akrad)
	case "814":
		postData814(result)
	case "815":
		copy(result[postIDLength:], postPointer815().Skrad)
	case "816":
		postData816(result)
	case "817":
		copy(result[postIDLength:], postPointer817().Utbnytrad)
	case "820":
		postData820(result)
	case "830":
		postData830(result)
	case "835":
		postData835(result)
	case "840":
		postData840(result)
	case "880":
		copy(result[postIDLength:], postPointer880().Ftrad)
	case "881":
		postData881(result)
	case "882":
		copy(result[postIDLength:], postPointer882().Forad)
	case "883":
		postData883(result)
	case "884":
		copy(result[postIDLength:], postPointer884().Nyrad)
	case "885":
		copy(result[postIDLength:], postPointer885().Porad)
	case "886":
		postData886(result)
	case "887":
		copy(result[postIDLength:], postPointer887().Nerad)
	case "888":
		postData888(result)
	case "892":
		copy(result[postIDLength:], postPointer892().Karad)
	case "893":
		postData893(result)
	case "894":
		copy(result[postIDLength:], postPointer894().Forbrad)
	case "895":
		copy(result[postIDLength:], postPointer895().Exrad)
	case "896":
		copy(result[postIDLength:], postPointer896().Ovrad)
	case "930":
		postData930(result)
	case "931":
		postData931(result)
	case "970":
		postData970(result)
	default:
		// Unknown post, return blank input which will fail test.
		return result
	}
	copy(result[postIDLength-3:], post)
	return result
}

func postData800(result []byte) {
	p := postPointer800()
	copy(result[postIDLength:], p.Firma)
	copy(result[postIDLength+200:], p.FirmaRegDat)
	copy(result[postIDLength+208:], p.Lagerbolag)
	copy(result[postIDLength+209:], p.FirmaMark)
	copy(result[postIDLength+210:], p.KonvMark)
	copy(result[postIDLength+215:], p.RegLan)
	offset := 217
	for i := 0; i < len(p.SkyddslanGRP); i++ {
		copy(result[postIDLength+offset:], p.SkyddslanGRP[i].Skyddslan)
		offset += 2
		copy(result[postIDLength+offset:], p.SkyddslanGRP[i].SkyddslanMark)
		offset++
	}
}

func postData808(result []byte) {
	p := postPointer808()
	copy(result[postIDLength:], p.Nybemdat)
	copy(result[postIDLength+8:], p.Nyforebolst)
	copy(result[postIDLength+12:], p.Nyannatdat)
	copy(result[postIDLength+20:], p.Nyemmark)
	copy(result[postIDLength+21:], p.Bemskmark)
	copy(result[postIDLength+22:], p.Skkonvtmark)
	copy(result[postIDLength+23:], p.Skoptionmark)
	copy(result[postIDLength+24:], p.Skbemdat)
	copy(result[postIDLength+32:], p.Skforebolst)
	copy(result[postIDLength+36:], p.Skannatdat)
	copy(result[postIDLength+44:], p.Bemregdat)
}

func postData810(result []byte) {
	p := postPointer810()
	copy(result[postIDLength+0:], p.Co)
	copy(result[postIDLength+50:], p.Gata)
	copy(result[postIDLength+82:], p.Postnr)
	copy(result[postIDLength+88:], p.Padr)
	copy(result[postIDLength+120:], p.Landkod)
	copy(result[postIDLength+125:], p.Land)
	copy(result[postIDLength+157:], p.Lan)
	copy(result[postIDLength+159:], p.Kommun)
	copy(result[postIDLength+161:], p.Regdat)
	copy(result[postIDLength+169:], p.Rakperf)
	copy(result[postIDLength+177:], p.Rakpert)
	copy(result[postIDLength+185:], p.Rakfrom)
	copy(result[postIDLength+189:], p.Raktom)
	copy(result[postIDLength+193:], p.Adrmark)
	copy(result[postIDLength+194:], p.Satemark)
	copy(result[postIDLength+195:], p.Rakper)
	copy(result[postIDLength+196:], p.Priv)
	copy(result[postIDLength+197:], p.Publ)
	copy(result[postIDLength+198:], p.Skade)
	copy(result[postIDLength+199:], p.Liv)
	copy(result[postIDLength+200:], p.Bildatdat)
	copy(result[postIDLength+208:], p.Ordenstadgregl)
	copy(result[postIDLength+216:], p.Stadfast)
	copy(result[postIDLength+224:], p.Stadfav)
	copy(result[postIDLength+226:], p.Oktrojater)
	copy(result[postIDLength+234:], p.Forloktroj)
	copy(result[postIDLength+242:], p.Koncforv)
	copy(result[postIDLength+250:], p.Sateort)
	copy(result[postIDLength+275:], p.Tillstdat)
	copy(result[postIDLength+283:], p.Tillsdat1)
	copy(result[postIDLength+291:], p.Tillstyp1)
	copy(result[postIDLength+296:], p.Tillsaterk1)
	copy(result[postIDLength+297:], p.Tillsdat2)
	copy(result[postIDLength+305:], p.Tillstyp2)
	copy(result[postIDLength+310:], p.Tillsaterk2)
	copy(result[postIDLength+311:], p.Tillsdat3)
	copy(result[postIDLength+319:], p.Tillstyp3)
	copy(result[postIDLength+324:], p.Tillsaterk3)
}

func postData811(result []byte) {
	p := postPointer811()
	copy(result[postIDLength:], p.Akap)
	copy(result[postIDLength+18:], p.Valuta1)
	copy(result[postIDLength+21:], p.Akapl)
	copy(result[postIDLength+39:], p.Valuta2)
	copy(result[postIDLength+42:], p.Akaph)
	copy(result[postIDLength+60:], p.Valuta3)
	copy(result[postIDLength+63:], p.Nombel)
	copy(result[postIDLength+81:], p.Totantakt)
	copy(result[postIDLength+99:], p.Antaktl)
	copy(result[postIDLength+117:], p.Antakth)
}

func postData812(result []byte) {
	p := postPointer812()
	copy(result[postIDLength:], p.Aktieslag)
	copy(result[postIDLength+30:], p.Antslag)
	copy(result[postIDLength+48:], p.Rostvarde)
	copy(result[postIDLength+58:], p.Antaktla)
	copy(result[postIDLength+76:], p.Antaktho)
}

func postData814(result []byte) {
	p := postPointer814()
	copy(result[postIDLength:], p.Skvaluta)
	copy(result[postIDLength+3:], p.Skbesldat)
	copy(result[postIDLength+11:], p.Skmark)
	copy(result[postIDLength+12:], p.Skbeslbel)
	copy(result[postIDLength+30:], p.Sklagst)
	copy(result[postIDLength+48:], p.Skhogst)
	copy(result[postIDLength+66:], p.Skteknatbel)
	copy(result[postIDLength+84:], p.Skkonvbel)
	copy(result[postIDLength+102:], p.Sktidutb1)
	copy(result[postIDLength+110:], p.Sktidutb2)
	copy(result[postIDLength+118:], p.Skoptionbel)
	copy(result[postIDLength+136:], p.Sktidnyt1)
	copy(result[postIDLength+144:], p.Sktidnyt2)
	copy(result[postIDLength+152:], p.Skapapokn)
	copy(result[postIDLength+170:], p.Skaktieslag)
	copy(result[postIDLength+210:], p.Toptant)
	copy(result[postIDLength+222:], p.Anttoptl)
	copy(result[postIDLength+238:], p.Anttopth)
	copy(result[postIDLength+254:], p.Toptfrom)
	copy(result[postIDLength+262:], p.Topttom)
	copy(result[postIDLength+270:], p.Toptml)
	copy(result[postIDLength+271:], p.Toptul)
	copy(result[postIDLength+272:], p.Sktyp)
	copy(result[postIDLength+273:], p.Skdelreg)
}

func postData816(result []byte) {
	p := postPointer816()
	copy(result[postIDLength:], p.Nytutbvaluta)
	copy(result[postIDLength+3:], p.Nytutbmark)
	copy(result[postIDLength+4:], p.Nytutbbel)
	copy(result[postIDLength+16:], p.NytutbbelPrecision)
	copy(result[postIDLength+22:], p.Nytutbnombel)
	copy(result[postIDLength+34:], p.NytutbnombelPrecision)
	copy(result[postIDLength+40:], p.Nytutbantakt)
}

func postData820(result []byte) {
	p := postPointer820()
	copy(result[postIDLength:], p.Ledl)
	copy(result[postIDLength+5:], p.Ledh)
	copy(result[postIDLength+10:], p.Suppl)
	copy(result[postIDLength+15:], p.Supph)
	copy(result[postIDLength+20:], p.Revl)
	copy(result[postIDLength+25:], p.Revh)
	copy(result[postIDLength+30:], p.Bolmvak)
	copy(result[postIDLength+38:], p.Bolmsakn)
	copy(result[postIDLength+46:], p.Forestsakn)
	copy(result[postIDLength+54:], p.Kompsakn)
	copy(result[postIDLength+62:], p.Styrsakn)
	copy(result[postIDLength+70:], p.Anm)
	copy(result[postIDLength+78:], p.Ejfult)
	copy(result[postIDLength+86:], p.Ejbehor)
	copy(result[postIDLength+94:], p.Vdsakn)
	copy(result[postIDLength+102:], p.Likvsakn)
	copy(result[postIDLength+110:], p.Delgmsakn)
	copy(result[postIDLength+118:], p.Revsakn)
	copy(result[postIDLength+126:], p.Senregdat)
	copy(result[postIDLength+134:], p.Vakansfinns)
	copy(result[postIDLength+135:], p.Ledvald)
	copy(result[postIDLength+138:], p.Supplvald)
	copy(result[postIDLength+141:], p.Styandrat)
}

func postData830(result []byte) {
	p := postPointer830()
	copy(result[postIDLength:], p.Sekf)
	copy(result[postIDLength+1:], p.Pnr)
	copy(result[postIDLength+11:], p.Funk1)
	copy(result[postIDLength+16:], p.Funk2)
	copy(result[postIDLength+21:], p.Funk3)
	copy(result[postIDLength+26:], p.Funk4)
	copy(result[postIDLength+31:], p.Typa)
	copy(result[postIDLength+32:], p.Type)
	copy(result[postIDLength+33:], p.Typu)
	copy(result[postIDLength+34:], p.Namn1)
	copy(result[postIDLength+84:], p.Namn2)
	copy(result[postIDLength+134:], p.Fco)
	copy(result[postIDLength+184:], p.Fgata)
	copy(result[postIDLength+216:], p.Fpostnr)
	copy(result[postIDLength+222:], p.Fpadr)
	copy(result[postIDLength+254:], p.Flandkod)
	copy(result[postIDLength+259:], p.Land)
	copy(result[postIDLength+291:], p.Utltillst)
	copy(result[postIDLength+299:], p.Perforsekl)
	copy(result[postIDLength+300:], p.Perforpnr)
	copy(result[postIDLength+310:], p.Insats)
	copy(result[postIDLength+326:], p.InsatsPrecision)
	copy(result[postIDLength+328:], p.Kval)
}

func postData835(result []byte) {
	p := postPointer835()
	copy(result[postIDLength:], p.Astsekel)
	copy(result[postIDLength+1:], p.Astpnr)
	copy(result[postIDLength+11:], p.Astftyp)
	copy(result[postIDLength+12:], p.Astrad)
}

func postData840(result []byte) {
	p := postPointer840()
	copy(result[postIDLength:], p.NamnTyp)
	copy(result[postIDLength+2:], p.Namn)
	copy(result[postIDLength+202:], p.NamnRegDat)
	copy(result[postIDLength+210:], p.BeslStamma)
	copy(result[postIDLength+211:], p.BeslStadg)
	offset := 212
	for i := 0; i < len(p.SkyddslanGRP); i++ {
		copy(result[postIDLength+offset:], p.SkyddslanGRP[i].Skyddslan)
		offset += 2
		copy(result[postIDLength+offset:], p.SkyddslanGRP[i].SkyddslanMark)
		offset++
	}
}

func postData881(result []byte) {
	p := postPointer881()
	copy(result[postIDLength:], p.Fovaluta)
	copy(result[postIDLength+3:], p.Fobesldat)
	copy(result[postIDLength+11:], p.Genom1)
	copy(result[postIDLength+13:], p.Genom2)
	copy(result[postIDLength+15:], p.Genom3)
	copy(result[postIDLength+17:], p.Genom4)
	copy(result[postIDLength+19:], p.Genom5)
	copy(result[postIDLength+21:], p.Genom6)
	copy(result[postIDLength+23:], p.Fondembel)
	copy(result[postIDLength+35:], p.FondembelPrecision)
	copy(result[postIDLength+41:], p.Fondnombel)
	copy(result[postIDLength+53:], p.FondnombelPrecision)
	copy(result[postIDLength+59:], p.Foantaktier)
	copy(result[postIDLength+77:], p.Genom7)
}

func postData883(result []byte) {
	p := postPointer883()
	copy(result[postIDLength:], p.Nyvaluta)
	copy(result[postIDLength+3:], p.Nybesldat)
	copy(result[postIDLength+11:], p.Nybeslav)
	copy(result[postIDLength+12:], p.Nyejfullbet)
	copy(result[postIDLength+13:], p.Nyembel)
	copy(result[postIDLength+25:], p.NyembelPrecision)
	copy(result[postIDLength+31:], p.Nygrans1)
	copy(result[postIDLength+43:], p.Nygrans1Precision)
	copy(result[postIDLength+49:], p.Nygrans2)
	copy(result[postIDLength+61:], p.Nygrans2Precision)
	copy(result[postIDLength+85:], p.Nyantaktier)
	copy(result[postIDLength+103:], p.Nyokurs)
	copy(result[postIDLength+115:], p.NyokursPrecision)
	copy(result[postIDLength+121:], p.Nydelantal)
	copy(result[postIDLength+139:], p.Nydelbel)
	copy(result[postIDLength+151:], p.NydelbelPrecision)
	copy(result[postIDLength+157:], p.Nyfullantal)
	copy(result[postIDLength+175:], p.Nyfullbet)
	copy(result[postIDLength+187:], p.NyfullbetPrecision)
	copy(result[postIDLength+193:], p.Nyejenomfbel)
	copy(result[postIDLength+205:], p.NyejenomfbelPrecision)
	copy(result[postIDLength+211:], p.Nyejenomfant)
	copy(result[postIDLength+229:], p.Nybetkon)
	copy(result[postIDLength+231:], p.Nybetapp)
	copy(result[postIDLength+233:], p.Nybetkvi)
	copy(result[postIDLength+235:], p.Nydelreg)
}

func postData886(result []byte) {
	p := postPointer886()
	copy(result[postIDLength:], p.Nevaluta)
	copy(result[postIDLength+3:], p.Nebesldat)
	copy(result[postIDLength+11:], p.Ngenom1)
	copy(result[postIDLength+13:], p.Ngenom2)
	copy(result[postIDLength+15:], p.Ngenom3)
	copy(result[postIDLength+17:], p.Ngenom4)
	copy(result[postIDLength+19:], p.Ngenom5)
	copy(result[postIDLength+21:], p.Ngenom6)
	copy(result[postIDLength+23:], p.Nandamal1)
	copy(result[postIDLength+25:], p.Nandamal2)
	copy(result[postIDLength+27:], p.Nandamal3)
	copy(result[postIDLength+29:], p.Nandamal4)
	copy(result[postIDLength+31:], p.Nandamal5)
	copy(result[postIDLength+33:], p.Nandamal6)
	copy(result[postIDLength+35:], p.Nembel)
	copy(result[postIDLength+53:], p.Netotbeslbet)
	copy(result[postIDLength+71:], p.Nenombel)
	copy(result[postIDLength+89:], p.Neejvarkst)
	copy(result[postIDLength+90:], p.Neantaktier)
	copy(result[postIDLength+108:], p.Netilstbesl)
	copy(result[postIDLength+116:], p.Netilstav)
	copy(result[postIDLength+120:], p.Netilstdat)
	copy(result[postIDLength+128:], p.Netillsbel)
	copy(result[postIDLength+146:], p.Neforfdat)
	copy(result[postIDLength+154:], p.Neforfav)
	copy(result[postIDLength+158:], p.Neforfbel)
	copy(result[postIDLength+176:], p.Neforfanta)
	copy(result[postIDLength+194:], p.Ngenom7)
	copy(result[postIDLength+196:], p.Ngenom8)
	copy(result[postIDLength+198:], p.Nlagst)
	copy(result[postIDLength+216:], p.Nhogst)
}

func postData888(result []byte) {
	p := postPointer888()
	copy(result[postIDLength:], p.Kalltyp)
	copy(result[postIDLength+10:], p.Kalldinr)
	copy(result[postIDLength+17:], p.Kalldinrar)
	copy(result[postIDLength+21:], p.Kalldatum)
	copy(result[postIDLength+29:], p.Kallkungdat)
	copy(result[postIDLength+37:], p.Kallbestrdat)
	copy(result[postIDLength+45:], p.Kallotrdinr)
	copy(result[postIDLength+52:], p.Kallotrdinrar)
	copy(result[postIDLength+56:], p.Kallotrdat)
	copy(result[postIDLength+64:], p.Kallotrdom)
}

func postData893(result []byte) {
	p := postPointer893()
	copy(result[postIDLength:], p.Markforb)
	copy(result[postIDLength+1:], p.Markomtyp1)
	copy(result[postIDLength+2:], p.Markomtyp2)
	copy(result[postIDLength+3:], p.Markomtyp3)
	copy(result[postIDLength+4:], p.Markomtyp4)
	copy(result[postIDLength+5:], p.Markomtyp5)
	copy(result[postIDLength+6:], p.Markomtyp6)
	copy(result[postIDLength+7:], p.Forbfom)
	copy(result[postIDLength+15:], p.Markomtyp7)
	copy(result[postIDLength+16:], p.Markomtyp8)
	copy(result[postIDLength+17:], p.Markomtyp9)
	copy(result[postIDLength+18:], p.Ejrevbsldat)
	copy(result[postIDLength+26:], p.Omvandlingsforbehall)
	copy(result[postIDLength+27:], p.Inlosenforbehall)
	copy(result[postIDLength+28:], p.Likvidationsbestammelse)
	copy(result[postIDLength+29:], p.Uppkopserbjudande)
}

func postData930(result []byte) {
	p := postPointer930()
	copy(result[postIDLength:], p.Forsenenper)
	copy(result[postIDLength+8:], p.Forsenregdat)
	copy(result[postIDLength+16:], p.Forsennr)
	copy(result[postIDLength+18:], p.Forbel)
	copy(result[postIDLength+31:], p.ForbelPrecision)
	copy(result[postIDLength+33:], p.Lagakraft)
	copy(result[postIDLength+41:], p.Avskrivet)
	copy(result[postIDLength+49:], p.Avskranel)
}

func postData931(result []byte) {
	p := postPointer931()
	copy(result[postIDLength:], p.Forflomper)
	copy(result[postIDLength+8:], p.Forftomper)
	copy(result[postIDLength+16:], p.Forldat)
	copy(result[postIDLength+24:], p.Forlregdat)
}

func postData970(result []byte) {
	p := postPointer970()
	copy(result[postIDLength:], p.Isekl)
	copy(result[postIDLength+1:], p.Iorgnummer)
	copy(result[postIDLength+11:], p.Ilopnr)
	copy(result[postIDLength+14:], p.Idinrar)
	copy(result[postIDLength+18:], p.Idinr)
	copy(result[postIDLength+25:], p.Inkdat)
	copy(result[postIDLength+33:], p.Iobjtyp)
	offset := 38
	for i := 0; i < len(p.RubricGRP); i++ {
		copy(result[postIDLength+offset:], p.RubricGRP[i].Rubric)
		offset += 10
	}
}

func postPointer(post string) (result interface{}) {
	switch post {
	case "800":
		result = postPointer800()
	case "808":
		result = postPointer808()
	case "810":
		result = postPointer810()
	case "811":
		result = postPointer811()
	case "812":
		result = postPointer812()
	case "813":
		result = postPointer813()
	case "814":
		result = postPointer814()
	case "815":
		result = postPointer815()
	case "816":
		result = postPointer816()
	case "817":
		result = postPointer817()
	case "820":
		result = postPointer820()
	case "830":
		result = postPointer830()
	case "835":
		result = postPointer835()
	case "840":
		result = postPointer840()
	case "880":
		result = postPointer880()
	case "881":
		result = postPointer881()
	case "882":
		result = postPointer882()
	case "883":
		result = postPointer883()
	case "884":
		result = postPointer884()
	case "885":
		result = postPointer885()
	case "886":
		result = postPointer886()
	case "887":
		result = postPointer887()
	case "888":
		result = postPointer888()
	case "892":
		result = postPointer892()
	case "893":
		result = postPointer893()
	case "894":
		result = postPointer894()
	case "895":
		result = postPointer895()
	case "896":
		result = postPointer896()
	case "930":
		result = postPointer930()
	case "931":
		result = postPointer931()
	case "970":
		result = postPointer970()
	}
	return result
}

func postPointer800() *Post800 {
	p := &Post800{
		Firma:       "0",
		FirmaRegDat: "1",
		Lagerbolag:  "2",
		FirmaMark:   "3",
		KonvMark:    "4",
		RegLan:      "5",
	}
	for i := 0; i < len(p.SkyddslanGRP); i++ {
		p.SkyddslanGRP[i].Skyddslan = fmt.Sprintf("%v", i)
		// 97 is 'a''
		p.SkyddslanGRP[i].SkyddslanMark = fmt.Sprintf("%c", i+97)
	}
	return p
}

func postPointer808() *Post808 {
	return &Post808{
		Nybemdat:     "0",
		Nyforebolst:  "1",
		Nyannatdat:   "2",
		Nyemmark:     "3",
		Bemskmark:    "4",
		Skkonvtmark:  "5",
		Skoptionmark: "6",
		Skbemdat:     "7",
		Skforebolst:  "8",
		Skannatdat:   "9",
		Bemregdat:    "10",
	}
}

func postPointer810() *Post810 {
	return &Post810{
		Co:             "0",
		Gata:           "1",
		Postnr:         "2",
		Padr:           "3",
		Landkod:        "4",
		Land:           "5",
		Lan:            "6",
		Kommun:         "7",
		Regdat:         "8",
		Rakperf:        "9",
		Rakpert:        "10",
		Rakfrom:        "11",
		Raktom:         "12",
		Adrmark:        "a",
		Satemark:       "b",
		Rakper:         "c",
		Priv:           "d",
		Publ:           "e",
		Skade:          "f",
		Liv:            "g",
		Bildatdat:      "20",
		Ordenstadgregl: "21",
		Stadfast:       "22",
		Stadfav:        "23",
		Oktrojater:     "24",
		Forloktroj:     "25",
		Koncforv:       "26",
		Sateort:        "27",
		Tillstdat:      "28",
		Tillsdat1:      "29",
		Tillstyp1:      "30",
		Tillsaterk1:    "h",
		Tillsdat2:      "32",
		Tillstyp2:      "33",
		Tillsaterk2:    "i",
		Tillsdat3:      "35",
		Tillstyp3:      "36",
		Tillsaterk3:    "j",
	}
}

func postPointer811() *Post811_170211 {
	return &Post811_170211{
		Akap:      "000000000000100000",
		Valuta1:   "1",
		Akapl:     "200000000000100000",
		Valuta2:   "3",
		Akaph:     "400000000000100000",
		Valuta3:   "5",
		Nombel:    "600000000000100000",
		Totantakt: "7",
		Antaktl:   "8",
		Antakth:   "9",
	}
}

func postPointer812() *Post812 {
	return &Post812{
		Aktieslag: "0",
		Antslag:   "1",
		Rostvarde: "2",
		Antaktla:  "3",
		Antaktho:  "4",
	}
}

func postPointer813() *Post813 {
	return &Post813{Akrad: "Akrad"}
}

func postPointer814() *Post814_161008 {
	return &Post814_161008{
		Skvaluta:    "0",
		Skbesldat:   "0",
		Skmark:      "0",
		Skbeslbel:   "0",
		Sklagst:     "0",
		Skhogst:     "0",
		Skteknatbel: "0",
		Skkonvbel:   "0",
		Sktidutb1:   "0",
		Sktidutb2:   "0",
		Skoptionbel: "0",
		Sktidnyt1:   "0",
		Sktidnyt2:   "0",
		Skapapokn:   "0",
		Skaktieslag: "0",
		Toptant:     "0",
		Anttoptl:    "0",
		Anttopth:    "0",
		Toptfrom:    "0",
		Topttom:     "0",
		Toptml:      "0",
		Toptul:      "0",
		Sktyp:       "0",
		Skdelreg:    "0",
	}
}

func postPointer815() *Post815 {
	return &Post815{Skrad: "Skrad"}
}

func postPointer816() *Post816 {
	return &Post816{
		Nytutbvaluta:          "0",
		Nytutbmark:            "1",
		Nytutbbel:             "2",
		NytutbbelPrecision:    "3",
		Nytutbnombel:          "4",
		NytutbnombelPrecision: "5",
		Nytutbantakt:          "6",
	}
}

func postPointer817() *Post817 {
	return &Post817{Utbnytrad: "Utbnytrad"}
}

func postPointer820() *Post820 {
	return &Post820{
		Ledl:        "0",
		Ledh:        "1",
		Suppl:       "2",
		Supph:       "3",
		Revl:        "4",
		Revh:        "5",
		Bolmvak:     "6",
		Bolmsakn:    "7",
		Forestsakn:  "8",
		Kompsakn:    "9",
		Styrsakn:    "10",
		Anm:         "11",
		Ejfult:      "12",
		Ejbehor:     "13",
		Vdsakn:      "14",
		Likvsakn:    "15",
		Delgmsakn:   "16",
		Revsakn:     "17",
		Senregdat:   "18",
		Vakansfinns: "a",
		Ledvald:     "20",
		Supplvald:   "21",
		Styandrat:   "22",
	}
}

func postPointer830() *Post830 {
	return &Post830{
		Sekf:            "0",
		Pnr:             "1",
		Funk1:           "2",
		Funk2:           "3",
		Funk3:           "4",
		Funk4:           "5",
		Typa:            "6",
		Type:            "7",
		Typu:            "8",
		Namn1:           "9",
		Namn2:           "10",
		Fco:             "11",
		Fgata:           "12",
		Fpostnr:         "13",
		Fpadr:           "14",
		Flandkod:        "15",
		Land:            "16",
		Utltillst:       "17",
		Perforsekl:      "a",
		Perforpnr:       "19",
		Insats:          "20",
		InsatsPrecision: "21",
		Kval:            "b",
	}
}

func postPointer835() *Post835 {
	return &Post835{
		Astsekel: "1",
		Astpnr:   "Astpnr",
		Astftyp:  "2",
		Astrad:   "Astrad",
	}
}

func postPointer840() *Post840 {
	p := &Post840{
		NamnTyp:    "0",
		Namn:       "1",
		NamnRegDat: "2",
		BeslStamma: "3",
		BeslStadg:  "4",
	}
	for i := 0; i < len(p.SkyddslanGRP); i++ {
		p.SkyddslanGRP[i].Skyddslan = fmt.Sprintf("%v", i)
		// 97 is 'a''
		p.SkyddslanGRP[i].SkyddslanMark = fmt.Sprintf("%c", i+97)
	}
	return p
}

func postPointer880() *Post880 {
	return &Post880{Ftrad: "Ftrad"}
}

func postPointer881() *Post881 {
	return &Post881{
		Fovaluta:            "0",
		Fobesldat:           "1",
		Genom1:              "2",
		Genom2:              "3",
		Genom3:              "4",
		Genom4:              "5",
		Genom5:              "6",
		Genom6:              "7",
		Fondembel:           "8",
		FondembelPrecision:  "9",
		Fondnombel:          "10",
		FondnombelPrecision: "11",
		Foantaktier:         "12",
		Genom7:              "13",
	}
}

func postPointer882() *Post882 {
	return &Post882{Forad: "Forad"}
}

func postPointer883() *Post883_161008 {
	return &Post883_161008{
		Nyvaluta:              "0",
		Nybesldat:             "1",
		Nybeslav:              "2",
		Nyejfullbet:           "3",
		Nyembel:               "4",
		NyembelPrecision:      "5",
		Nygrans1:              "6",
		Nygrans1Precision:     "7",
		Nygrans2:              "8",
		Nygrans2Precision:     "9",
		Nyantaktier:           "12",
		Nyokurs:               "13",
		NyokursPrecision:      "14",
		Nydelantal:            "15",
		Nydelbel:              "16",
		NydelbelPrecision:     "17",
		Nyfullantal:           "18",
		Nyfullbet:             "19",
		NyfullbetPrecision:    "20",
		Nyejenomfbel:          "21",
		NyejenomfbelPrecision: "22",
		Nyejenomfant:          "23",
		Nybetkon:              "24",
		Nybetapp:              "25",
		Nybetkvi:              "26",
		Nydelreg:              "a",
	}
}

func postPointer884() *Post884 {
	return &Post884{Nyrad: "Nyrad"}
}

func postPointer885() *Post885 {
	return &Post885{Porad: "Porad"}
}

func postPointer886() *Post886_230204 {
	return &Post886_230204{
		Nevaluta:     "0",
		Nebesldat:    "1",
		Ngenom1:      "2",
		Ngenom2:      "3",
		Ngenom3:      "4",
		Ngenom4:      "5",
		Ngenom5:      "6",
		Ngenom6:      "7",
		Nandamal1:    "8",
		Nandamal2:    "9",
		Nandamal3:    "10",
		Nandamal4:    "11",
		Nandamal5:    "12",
		Nandamal6:    "13",
		Nembel:       "14",
		Netotbeslbet: "16",
		Nenombel:     "18",
		Neejvarkst:   "a",
		Neantaktier:  "21",
		Netilstbesl:  "22",
		Netilstav:    "23",
		Netilstdat:   "24",
		Netillsbel:   "25",
		Neforfdat:    "27",
		Neforfav:     "28",
		Neforfbel:    "29",
		Neforfanta:   "31",
		Ngenom7:      "32",
		Ngenom8:      "33",
		Nlagst:       "34",
		Nhogst:       "36",
	}
}

func postPointer887() *Post887 {
	return &Post887{Nerad: "Nerad"}
}

func postPointer888() *Post888 {
	return &Post888{
		Kalltyp:       "0",
		Kalldinr:      "1",
		Kalldinrar:    "2",
		Kalldatum:     "3",
		Kallkungdat:   "4",
		Kallbestrdat:  "5",
		Kallotrdinr:   "6",
		Kallotrdinrar: "7",
		Kallotrdat:    "8",
		Kallotrdom:    "9",
	}
}

func postPointer892() *Post892 {
	return &Post892{Karad: "Karad"}
}

func postPointer893() *Post893 {
	return &Post893{
		Markforb:                "0",
		Markomtyp1:              "1",
		Markomtyp2:              "2",
		Markomtyp3:              "3",
		Markomtyp4:              "4",
		Markomtyp5:              "5",
		Markomtyp6:              "6",
		Forbfom:                 "7",
		Markomtyp7:              "8",
		Markomtyp8:              "9",
		Markomtyp9:              "a",
		Ejrevbsldat:             "11",
		Omvandlingsforbehall:    "a",
		Inlosenforbehall:        "b",
		Likvidationsbestammelse: "c",
		Uppkopserbjudande:       "d",
	}
}

func postPointer894() *Post894 {
	return &Post894{Forbrad: "Forbrad"}
}

func postPointer895() *Post895 {
	return &Post895{Exrad: "Exrad"}
}

func postPointer896() *Post896 {
	return &Post896{Ovrad: "Ovrad"}
}

func postPointer930() *Post930 {
	return &Post930{
		Forsenenper:     "0",
		Forsenregdat:    "1",
		Forsennr:        "2",
		Forbel:          "3",
		ForbelPrecision: "4",
		Lagakraft:       "5",
		Avskrivet:       "6",
		Avskranel:       "7",
	}
}

func postPointer931() *Post931 {
	return &Post931{
		Forflomper: "0",
		Forftomper: "1",
		Forldat:    "2",
		Forlregdat: "3",
	}
}

func postPointer970() *Post970 {
	p := &Post970{
		Isekl:      "0",
		Iorgnummer: "1",
		Ilopnr:     "2",
		Idinrar:    "3",
		Idinr:      "4",
		Inkdat:     "5",
		Iobjtyp:    "6",
	}
	for i := 0; i < len(p.RubricGRP); i++ {
		p.RubricGRP[i].Rubric = fmt.Sprintf("%v", i)
	}
	return p
}

func postUnmarshal(line []byte) (interface{}, error) {
	var av AviserPost
	i, err := UnmarshalAviserPost(line, &av)
	if err != nil {
		return nil, err
	}
	if i != postIDLength {
		return nil, errors.New("wrong Avisers ize: " + string(line))
	}
	p := postPointer(av.PostTyp)
	i, err = UnmarshalData(line, p)
	if err != nil {
		return nil, err
	}
	if i != aviserDataLength {
		return nil, errors.New("wrong Data size: " + string(line))
	}
	return p, nil
}
