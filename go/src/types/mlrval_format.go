package types

import (
	"errors"
	"fmt"
)

//----------------------------------------------------------------
//* need int/float
//  llx -> x etc
//  https://golang.org/pkg/fmt/
//
//  pre-stuff
//
//  %
//
//  +-0' space
//
//  ll|l
//  %%
//  bdiouxDOUX fegFEG s
//
//  post-stuff
// ----------------------------------------------------------------

// ----------------------------------------------------------------
//* callsites:
//  o fmtnum($mv, "%d")
//    - numeric only
//  o format($mv, "%s")
//    - make this new DSL function
//  o --ofmt
//    - numeric only
//  o format-values verb
//    - -i, -f, -s
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// pieces:
// * given user-level format string
// * factory-something to make methods/classes which take a mlrval/type & a Go-internal format string
// * methods which take a mlrval/type & a Go-internal format string
// * map from user-level string to instance -- ?
// ----------------------------------------------------------------

// func MlrvalFormatterFactory(userLevelFormatString string) *MlrvalFormatter{
// }

// ----------------------------------------------------------------
type iMlrvalFormatter interface {
	Format(mlrval *Mlrval) string
}

func NewMlrvalFormatter(
	userLevelFormatString string,
) (iMlrvalFormatter, error) {
	return nil, errors.New("TBD") // TODO
}

// ----------------------------------------------------------------
type mlrvalFormatterToFloat struct {
	goFormatString string
}

func newMlrvalFormatterToFloat(userLevelFormatString string) iMlrvalFormatter {
	return &mlrvalFormatterToFloat{
		goFormatString: userLevelFormatString, // TODO temp
	}
}

func (formatter *mlrvalFormatterToFloat) Format(mlrval *Mlrval) string {
	floatValue, isFloat := mlrval.GetFloatValue()
	if isFloat {
		return fmt.Sprintf(formatter.goFormatString, floatValue)
	}
	intValue, isInt := mlrval.GetIntValue()
	if isInt {
		return fmt.Sprintf(formatter.goFormatString, float64(intValue))
	}
	return mlrval.String()
}

// ----------------------------------------------------------------
type mlrvalFormatterToInt struct {
	goFormatString string
}

func newMlrvalFormatterToInt(userLevelFormatString string) iMlrvalFormatter {
	return &mlrvalFormatterToInt{
		goFormatString: userLevelFormatString, // TODO temp
	}
}

func (formatter *mlrvalFormatterToInt) Format(mlrval *Mlrval) string {
	intValue, isInt := mlrval.GetIntValue()
	if isInt {
		return fmt.Sprintf(formatter.goFormatString, intValue)
	}
	floatValue, isFloat := mlrval.GetFloatValue()
	if isFloat {
		return fmt.Sprintf(formatter.goFormatString, int(floatValue))
	}
	return mlrval.String()
}

// ----------------------------------------------------------------
type mlrvalFormatterToString struct {
	goFormatString string
}

func newMlrvalFormatterToString(userLevelFormatString string) iMlrvalFormatter {
	return &mlrvalFormatterToString{
		goFormatString: userLevelFormatString, // TODO temp
	}
}

func (formatter *mlrvalFormatterToString) Format(mlrval *Mlrval) string {
	return fmt.Sprintf(formatter.goFormatString, mlrval.String())
}
