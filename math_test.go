package sprigmath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tpl := `{{ 3 | add 1 2}}`
	if err := runt(tpl, `6`); err != nil {
		t.Error(err)
	}

	tpl = `{{ 3 | add 1 2.5}}`
	if err := runt(tpl, `6.5`); err != nil {
		t.Error(err)
	}
}

func TestBiggest(t *testing.T) {
	tpl := `{{ biggest 1 2 3 345 5 6 7}}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}

	tpl = `{{ max 345}}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}

	tpl = `{{ biggest 1 2 3.0 345.7 5 6 7}}`
	if err := runt(tpl, `345.7`); err != nil {
		t.Error(err)
	}
}
func TestMin(t *testing.T) {
	tpl := `{{ min 1 2 3 345 5 6 7}}`
	if err := runt(tpl, `1`); err != nil {
		t.Error(err)
	}

	tpl = `{{ min 345}}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}

	tpl = `{{ min 1.2 2 3 345 5 6 7}}`
	if err := runt(tpl, `1.2`); err != nil {
		t.Error(err)
	}

	tpl = `{{ min 1.2 2 3 345.3 5 6 7}}`
	if err := runt(tpl, `1.2`); err != nil {
		t.Error(err)
	}
}

func TestMisc(t *testing.T) {
	tpl := `{{ pi | cos }}`
	if err := runt(tpl, `-1`); err != nil {
		t.Error(err)
	}
}

func TestMul(t *testing.T) {
	tpl := `{{ mul 1 2 3 }}`
	if err := runt(tpl, "6"); err != nil {
		t.Error(err)
	}

	tpl = `{{ mul 1 2.5 3 }}`
	if err := runt(tpl, "7.5"); err != nil {
		t.Error(err)
	}

	tpl = `{{ mul 1.0 2.5 3.0 }}`
	if err := runt(tpl, "7.5"); err != nil {
		t.Error(err)
	}

	tpl = `{{ mul 5.4 }}`
	if err := runt(tpl, "5.4"); err != nil {
		t.Error(err)
	}

	tpl = `{{ mul 5 }}`
	if err := runt(tpl, "5"); err != nil {
		t.Error(err)
	}

	tpl = `{{ mul "bob" }}`
	if err := runerr(tpl, "mul[arg0]: bob is not a float64 or int64"); err != nil {
		t.Error(err)
	}
}
