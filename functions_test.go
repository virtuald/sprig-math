package sprigmath

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var EPSILON = 0.00000001

func floatEquals(a, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

func testError(errstr string, err error) error {
	if err == nil {
		if errstr != "" {
			return fmt.Errorf("Expected '%s', error did not occur", errstr)
		}
	} else {
		if errstr == "" {
			return fmt.Errorf("Did not expect error, got '%s'", err)
		} else if !strings.HasSuffix(fmt.Sprintf("%s", err), errstr) {
			return fmt.Errorf("Expected error '%s', got '%s'", errstr, err)
		}
	}

	return nil
}

// runt runs a template and checks that the output exactly matches the expected string.
func runt(tpl, expect string) error {
	return runtv(tpl, expect, map[string]string{})
}

func runerr(tpl, errstr string) error {
	err := runtv(tpl, "", map[string]string{})
	return testError(errstr, err)
}

// runtv takes a template, and expected return, and values for substitution.
//
// It runs the template and verifies that the output is an exact match.
func runtv(tpl, expect string, vars interface{}) error {
	fmap := GenericFuncMap()
	t := template.Must(template.New("test").Funcs(fmap).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return err
	}
	if expect != b.String() {
		return fmt.Errorf("Expected '%s', got '%s'", expect, b.String())
	}
	return nil
}

// runRaw runs a template with the given variables and returns the result.
func runRaw(tpl string, vars interface{}) (string, error) {
	fmap := GenericFuncMap()
	t := template.Must(template.New("test").Funcs(fmap).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
