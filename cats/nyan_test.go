package cats

import "testing"

func TestSimpleMeowify(t *testing.T) {
	nyan := Nyan{}
	if nyan.Meowify("cats") != "nyan" {
		t.Error(`nyan.Meowify("cats") != "nyan"`)
	}
}

func TestSimpleMeowifyLineUrl(t *testing.T) {
	nyan := Nyan{}
	if nyan.MeowifyLine("http://9gag.com") != "nyan://nyan.nya" {
		t.Error(`nyan.MeowifyLine("http://9gag.com") != "nyan://nyan.nya"`)
	}
}

func TestNyanifyLineUrl(t *testing.T) {
	nyan := Nyan{}
	if nyan.MeowifyLine("http://www.example.com/cat") != "nyan://nya.nyannya.nya/nya" {
		t.Error(`nyan.MeowifyLine("http://www.example.com/cat") != "nyan://nya.nyannya.nya/nya"`)
	}
}
