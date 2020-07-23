package driver

import (
	"fmt"
	"os"
	"testing"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

const (
	deviceResourceBool       = "Bool"
	deviceResourceInt8       = "Int8"
	deviceResourceInt16      = "Int16"
	deviceResourceInt32      = "Int32"
	deviceResourceInt64      = "Int64"
	deviceResourceUint8      = "Uint8"
	deviceResourceUint16     = "Uint16"
	deviceResourceUint32     = "Uint32"
	deviceResourceUint64     = "Uint64"
	deviceResourceFloat32    = "Float32"
	deviceResourceFloat64    = "Float64"
	deviceResourceBinary     = "Binary"
	deviceName               = "Random-Value-Device"
	deviceCommandNameBool    = "Bool"
	deviceCommandNameInt8    = "Int8"
	deviceCommandNameInt16   = "Int16"
	deviceCommandNameInt32   = "Int32"
	deviceCommandNameInt64   = "Int64"
	deviceCommandNameUint8   = "Uint8"
	deviceCommandNameUint16  = "Uint16"
	deviceCommandNameUint32  = "Uint32"
	deviceCommandNameUint64  = "Uint64"
	deviceCommandNameFloat32 = "Float32"
	deviceCommandNameFloat64 = "Float64"
	enableRandomizationTrue  = true
	rounds                   = 10
)

func prepareDB() *db {
	db := getDb()
	if err := db.openDb(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := db.exec(SqlDropTable); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := db.exec(SqlCreateTable); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ds := [][]interface{}{
		{deviceName, deviceCommandNameBool, deviceResourceBool, enableRandomizationTrue, typeBool, "true"},
		{deviceName, deviceCommandNameInt8, deviceResourceInt8, enableRandomizationTrue, typeInt8, "0"},
		{deviceName, deviceCommandNameInt16, deviceResourceInt16, enableRandomizationTrue, typeInt16, "0"},
		{deviceName, deviceCommandNameInt32, deviceResourceInt32, enableRandomizationTrue, typeInt32, "0"},
		{deviceName, deviceCommandNameInt64, deviceResourceInt64, enableRandomizationTrue, typeInt64, "0"},
		{deviceName, deviceCommandNameUint8, deviceResourceUint8, enableRandomizationTrue, typeUint8, "0"},
		{deviceName, deviceCommandNameUint16, deviceResourceUint16, enableRandomizationTrue, typeUint16, "0"},
		{deviceName, deviceCommandNameUint32, deviceResourceUint32, enableRandomizationTrue, typeUint32, "0"},
		{deviceName, deviceCommandNameUint64, deviceResourceUint64, enableRandomizationTrue, typeUint64, "0"},
		{deviceName, deviceCommandNameFloat32, deviceResourceFloat32, enableRandomizationTrue, typeFloat32, "0"},
		{deviceName, deviceCommandNameFloat64, deviceResourceFloat64, enableRandomizationTrue, typeFloat64, "0"},
	}
	for _, d := range ds {
		if err := db.exec(SqlInsert, d...); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return db
}

func TestValue_Bool(t *testing.T) {
	db := prepareDB()
	defer func() {
		if err := db.closeDb(); err != nil {
			t.Fatal(err)
		}
	}()

	vd := newVirtualDevice()
	v1, err := vd.read(deviceName, deviceResourceBool, typeBool, "", "", db)
	if err != nil {
		t.Fatal(err)
	}

	//the return string must be convertible to boolean
	b1, err := v1.BoolValue()
	if err != nil {
		t.Fatal(err)
	}

	//EnableRandomization = true
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, deviceResourceBool, typeBool, "", "", db)
		b2, _ := v2.BoolValue()
		if b1 != b2 {
			break
		}
		if x == rounds {
			t.Fatalf("EnableRandomization is true, but got same read in %d rounds", rounds)
		}
	}

	//EnableRandomization = false
	if err := db.exec(SqlUpdateRandomization, false, deviceName, deviceResourceBool); err != nil {
		t.Fatal(err)
	}

	v1, _ = vd.read(deviceName, deviceResourceBool, typeBool, "", "", db)
	b1, _ = v1.BoolValue()
	for x := 0; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, deviceResourceBool, typeBool, "", "", db)
		b2, _ := v2.BoolValue()
		if b1 != b2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func TestValueIntx(t *testing.T) {
	ValueIntx(t, deviceResourceInt8, typeInt8, "-128", "127")
	ValueIntx(t, deviceResourceInt8, typeInt8, "", "")
	ValueIntx(t, deviceResourceInt16, typeInt16, "-32768", "32767")
	ValueIntx(t, deviceResourceInt16, typeInt16, "", "")
	ValueIntx(t, deviceResourceInt32, typeInt32, "-2147483648", "2147483647")
	ValueIntx(t, deviceResourceInt32, typeInt32, "", "")
	ValueIntx(t, deviceResourceInt64, typeInt64, "-9223372036854775808", "9223372036854775807")
	ValueIntx(t, deviceResourceInt64, typeInt64, "", "")
}

func TestValueUintx(t *testing.T) {
	ValueUintx(t, deviceResourceUint8, typeUint8, "0", "255")
	ValueUintx(t, deviceResourceUint8, typeUint8, "", "")
	ValueUintx(t, deviceResourceUint16, typeUint16, "0", "65535")
	ValueUintx(t, deviceResourceUint16, typeUint16, "", "")
	ValueUintx(t, deviceResourceUint32, typeUint32, "0", "4294967295")
	ValueUintx(t, deviceResourceUint32, typeUint32, "", "")
	ValueUintx(t, deviceResourceUint64, typeUint64, "0", "18446744073709551615")
	ValueUintx(t, deviceResourceUint64, typeUint64, "", "")
}

func TestValueFloatx(t *testing.T) {
	ValueFloatx(t, deviceResourceFloat32, typeFloat32, "-3.40282346638528859811704183484516925440e+38", "3.40282346638528859811704183484516925440e+38")
	ValueFloatx(t, deviceResourceFloat32, typeFloat32, "", "")
	ValueFloatx(t, deviceResourceFloat64, typeFloat64, "-1.797693134862315708145274237317043567981e+308", "1.797693134862315708145274237317043567981e+308")
	ValueFloatx(t, deviceResourceFloat64, typeFloat64, "", "")
}

func TestValueBinary(t *testing.T) {
	vd := newVirtualDevice()
	v1, err := vd.read(deviceName, deviceResourceBinary, typeBinary, "", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	//the return string must be convertible to binary
	_, err = v1.BinaryValue()
	if err != nil {
		t.Fatal(err)
	}
}

func ValueIntx(t *testing.T, dr, typeName, minStr, maxStr string) {
	db := prepareDB()
	defer func() {
		if err := db.closeDb(); err != nil {
			t.Fatal(err)
		}
	}()

	//EnableRandomization = true
	if err := db.exec(SqlUpdateRandomization, true, deviceName, dr); err != nil {
		t.Fatal(err)
	}

	vd := newVirtualDevice()

	min, _ := parseStrToInt(minStr, 64)
	max, _ := parseStrToInt(maxStr, 64)

	var i1 int64
	for x := 1; x <= rounds; x++ {
		vn, err := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		if err != nil {
			t.Fatal(err)
		}
		in := getIntValue(vn)

		if x == 1 {
			i1 = in
		}
		if i1 != in {
			break
		}
		if x == rounds {
			t.Fatalf("EnableRandomization is true, but got same read in %d rounds", rounds)
		}
	}

	for x := 1; x <= rounds; x++ {
		v, err := vd.read(deviceName, dr, typeName, minStr, maxStr, db)

		if err != nil {
			t.Fatal(err)
		}
		i := getIntValue(v)
		if err != nil {
			t.Fatal(err)
		}
		if minStr != "" && maxStr != "" {
			if i < min || i > max {
				t.Fatalf("random read: %d,  out of range: %s ~ %s", i, minStr, maxStr)
			}
		}
	}

	//EnableRandomization = false
	if err := db.exec(SqlUpdateRandomization, false, deviceName, dr); err != nil {
		t.Fatal(err)
	}

	v1, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
	i1 = getIntValue(v1)
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		i2 := getIntValue(v2)
		if i1 != i2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func ValueUintx(t *testing.T, dr, typeName, minStr, maxStr string) {
	db := prepareDB()
	defer func() {
		if err := db.closeDb(); err != nil {
			t.Fatal(err)
		}
	}()

	//EnableRandomization = true
	if err := db.exec(SqlUpdateRandomization, true, deviceName, dr); err != nil {
		t.Fatal(err)
	}

	vd := newVirtualDevice()

	min, _ := parseStrToUint(minStr, 64)
	max, _ := parseStrToUint(maxStr, 64)

	var i1 uint64
	for x := 1; x <= rounds; x++ {
		vn, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		in := getUintValue(vn)

		if x == 1 {
			i1 = in
		}
		if i1 != in {
			break
		}
		if x == rounds {
			t.Fatalf("EnableRandomization is true, but got same read in %d rounds", rounds)
		}
	}

	for x := 1; x <= rounds; x++ {
		v, err := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		if err != nil {
			t.Fatal(err)
		}
		i := getUintValue(v)
		if err != nil {
			t.Fatal(err)
		}
		if minStr != "" && maxStr != "" {
			if i < min || i > max {
				t.Fatalf("random read: %d,  out of range: %d ~ %d", i, min, max)
			}
		}
	}

	//EnableRandomization = false
	if err := db.exec(SqlUpdateRandomization, false, deviceName, dr); err != nil {
		t.Fatal(err)
	}

	v1, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
	i1 = getUintValue(v1)
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		i2 := getUintValue(v2)
		if i1 != i2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func ValueFloatx(t *testing.T, dr, typeName, minStr, maxStr string) {
	db := prepareDB()
	defer func() {
		if err := db.closeDb(); err != nil {
			t.Fatal(err)
		}
	}()

	//EnableRandomization = true
	if err := db.exec(SqlUpdateRandomization, true, deviceName, dr); err != nil {
		t.Fatal(err)
	}

	vd := newVirtualDevice()

	min, _ := parseStrToFloat(minStr, 64)
	max, _ := parseStrToFloat(maxStr, 64)

	var f1 float64
	for x := 1; x <= rounds; x++ {
		vn, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		fn := getFloatValue(vn)
		if x == 1 {
			f1 = fn
		}
		if f1 != fn {
			break
		}
		if x == rounds {
			t.Fatalf("EnableRandomization is true, but got same read in %d rounds", rounds)
		}
	}

	for x := 1; x <= rounds; x++ {
		v, err := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		if err != nil {
			t.Fatal(err)
		}
		f := getFloatValue(v)
		if err != nil {
			t.Fatal(err)
		}
		if minStr != "" && maxStr != "" {
			if f < min || f > max {
				t.Fatalf("random read: %f,  out of range: %f ~ %f", f, min, max)
			}
		}
	}

	//EnableRandomization = false
	if err := db.exec(SqlUpdateRandomization, false, deviceName, dr); err != nil {
		t.Fatal(err)
	}

	v1, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
	f1 = getFloatValue(v1)
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, dr, typeName, minStr, maxStr, db)
		f2 := getFloatValue(v2)
		if f1 != f2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func getIntValue(cv *dsModels.CommandValue) int64 {
	switch cv.Type {
	case dsModels.Int8:
		v, _ := cv.Int8Value()
		return int64(v)
	case dsModels.Int16:
		v, _ := cv.Int16Value()
		return int64(v)
	case dsModels.Int32:
		v, _ := cv.Int32Value()
		return int64(v)
	case dsModels.Int64:
		v, _ := cv.Int64Value()
		return v
	default:
		return 0
	}
}

func getUintValue(cv *dsModels.CommandValue) uint64 {
	switch cv.Type {
	case dsModels.Uint8:
		v, _ := cv.Uint8Value()
		return uint64(v)
	case dsModels.Uint16:
		v, _ := cv.Uint16Value()
		return uint64(v)
	case dsModels.Uint32:
		v, _ := cv.Uint32Value()
		return uint64(v)
	case dsModels.Uint64:
		v, _ := cv.Uint64Value()
		return v
	default:
		return 0
	}
}

func getFloatValue(cv *dsModels.CommandValue) float64 {
	switch cv.Type {
	case dsModels.Float32:
		v, _ := cv.Float32Value()
		return float64(v)
	case dsModels.Float64:
		v, _ := cv.Float64Value()
		return v
	default:
		return 0
	}
}
