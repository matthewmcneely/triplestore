package triplestore

import (
	"fmt"
	"strings"
)

type XsdType string

var (
	XsdString   = XsdType("xs:string")
	XsdBoolean  = XsdType("xs:boolean")
	XsdDateTime = XsdType("xs:dateTime")

	// 64-bit floating point numbers
	XsdDouble = XsdType("xs:double")
	// 32-bit floating point numbers
	XsdFloat = XsdType("xs:float")

	// signed 32 or 64 bit
	XsdInteger = XsdType("xs:integer")
	// signed (8 bit)
	XsdByte = XsdType("xs:integer")
	// signed (16 bit)
	XsdShort = XsdType("xs:integer")

	// unsigned 32 or 64 bit
	XsdUinteger = XsdType("xs:integer")
	// unsigned 8 bit
	XsdUnsignedByte = XsdType("xs:integer")
	// unsigned 16 bit
	XsdUnsignedShort = XsdType("xs:integer")
)

const XMLSchemaNamespace = "http://www.w3.org/2001/XMLSchema"

func (x XsdType) NTriplesNamespaced() string {
	splits := strings.Split(string(x), ":")
	if len(splits) != 2 {
		return string(x)
	}

	return fmt.Sprintf("%s#%s", XMLSchemaNamespace, splits[1])
}
