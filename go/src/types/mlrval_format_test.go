// ================================================================
// Most Miller tests (thousands of them) are command-line-driven via
// mlr regtest. Here are some cases needing special focus.
// ================================================================

// Invoke as:
// * cd go/src/types
// * go test
// Or:
// * cd go
// * go test miller/src/types/...

package types

import (
	"testing"
)

func TestFormatterToString(t *testing.T) {

	mv := MlrvalPointerFromString("hello")
	formatter := newMlrvalFormatterToString("%s")
	if formatter.Format(mv) != "hello" {
		t.Fatal()
	}

	mv = MLRVAL_NULL
	formatter = newMlrvalFormatterToString("%s")
	if formatter.Format(mv) != "null" {
		t.Fatal()
	}

	mv = MLRVAL_TRUE
	formatter = newMlrvalFormatterToString("%s")
	if formatter.Format(mv) != "true" {
		t.Fatal()
	}

	mv = MLRVAL_FALSE
	formatter = newMlrvalFormatterToString("%s")
	if formatter.Format(mv) != "false" {
		t.Fatal()
	}

	mv = MlrvalPointerFromInt(10)
	formatter = newMlrvalFormatterToString("%s")
	if formatter.Format(mv) != "10" {
		t.Fatal()
	}

	mv = MlrvalPointerFromString("hello")
	formatter = newMlrvalFormatterToString("[[%s]]")
	if formatter.Format(mv) != "[[hello]]" {
		t.Fatal()
	}

}

func TestFormatterToInt(t *testing.T) {

	mv := MlrvalPointerFromString("hello")
	formatter := newMlrvalFormatterToInt("%d")
	if formatter.Format(mv) != "hello" {
		t.Fatal()
	}

	mv = MLRVAL_NULL
	formatter = newMlrvalFormatterToInt("%d")
	if formatter.Format(mv) != "null" {
		t.Fatal()
	}

	mv = MLRVAL_TRUE
	formatter = newMlrvalFormatterToInt("%d")
	if formatter.Format(mv) != "true" {
		t.Fatal()
	}

	mv = MLRVAL_FALSE
	formatter = newMlrvalFormatterToInt("%d")
	if formatter.Format(mv) != "false" {
		t.Fatal()
	}

	mv = MlrvalPointerFromInt(10)
	formatter = newMlrvalFormatterToInt("%d")
	if formatter.Format(mv) != "10" {
		t.Fatal()
	}

	mv = MlrvalPointerFromInt(10)
	formatter = newMlrvalFormatterToInt("[[0x%x]]")
	if formatter.Format(mv) != "[[0xa]]" {
		t.Fatal()
	}

	mv = MlrvalPointerFromFloat64(10.1)
	formatter = newMlrvalFormatterToInt("%d")
	if formatter.Format(mv) != "10" {
		t.Fatal()
	}

}

func TestFormatterToFloat(t *testing.T) {

	mv := MlrvalPointerFromString("hello")
	formatter := newMlrvalFormatterToFloat("%.3f")
	if formatter.Format(mv) != "hello" {
		t.Fatal()
	}

	mv = MLRVAL_NULL
	formatter = newMlrvalFormatterToFloat("%.3f")
	if formatter.Format(mv) != "null" {
		t.Fatal()
	}

	mv = MLRVAL_TRUE
	formatter = newMlrvalFormatterToFloat("%.3f")
	if formatter.Format(mv) != "true" {
		t.Fatal()
	}

	mv = MLRVAL_FALSE
	formatter = newMlrvalFormatterToFloat("%.3f")
	if formatter.Format(mv) != "false" {
		t.Fatal()
	}

	mv = MlrvalPointerFromInt(10)
	formatter = newMlrvalFormatterToFloat("%.3f")
	if formatter.Format(mv) != "10.000" {
		t.Fatal()
	}

	mv = MlrvalPointerFromFloat64(1.234)
	formatter = newMlrvalFormatterToFloat("%.3f")
	if formatter.Format(mv) != "1.234" {
		t.Fatal()
	}

}
