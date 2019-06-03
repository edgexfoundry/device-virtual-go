package driver

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

const (
	deviceName               = "Random-Value-Device"
	deviceCommandNameBool    = "RandomValue_Bool"
	deviceCommandNameInt8    = "RandomValue_Int8"
	deviceCommandNameInt16   = "RandomValue_Int16"
	deviceCommandNameInt32   = "RandomValue_Int32"
	deviceCommandNameInt64   = "RandomValue_Int64"
	deviceCommandNameUint8   = "RandomValue_Uint8"
	deviceCommandNameUint16  = "RandomValue_Uint16"
	deviceCommandNameUint32  = "RandomValue_Uint32"
	deviceCommandNameUint64  = "RandomValue_Uint64"
	deviceCommandNameFloat32 = "RandomValue_Float32"
	deviceCommandNameFloat64 = "RandomValue_Float64"
	enableRandomizationTrue  = "true"
)

func init() {
	if _, err := os.Stat(qlDatabaseDir); os.IsNotExist(err) {
		if err := os.Mkdir(qlDatabaseDir, os.ModeDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	db := getDb()
	if err := db.openDb(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		if err := db.closeDb(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	if err := db.exec(SqlDropTable); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := db.exec(SqlCreateTable); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ds := [][]string{
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
		b, _ := strconv.ParseBool(d[3])
		if err := db.exec(SqlInsert, d[0], d[1], d[2], b, d[4], d[5]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func TestValue_Bool(t *testing.T) {
	db := getDb()
	if err := db.openDb(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.closeDb(); err != nil {
			t.Fatal(err)
		}
	}()

	vd := newVirtualDevice()
	v1, err := vd.read(deviceName, deviceResourceBool, "", "", db)
	if err != nil {
		t.Fatal(err)
	}

	//the return string must be convertible to boolean
	b1, err := v1.BoolValue()
	if err != nil {
		t.Fatal(err)
	}

	rounds := 20
	//EnableRandomization = true
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, deviceResourceBool, "", "", db)
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

	v1, _ = vd.read(deviceName, deviceResourceBool, "", "", db)
	b1, _ = v1.BoolValue()
	for x := 0; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, deviceResourceBool, "", "", db)
		b2, _ := v2.BoolValue()
		if b1 != b2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func TestValueIntx(t *testing.T) {
	ValueIntx(t, deviceResourceInt8, "-128", "127")
	ValueIntx(t, deviceResourceInt8, "", "")
	ValueIntx(t, deviceResourceInt16, "-32768", "32767")
	ValueIntx(t, deviceResourceInt16, "", "")
	ValueIntx(t, deviceResourceInt32, "-2147483648", "2147483647")
	ValueIntx(t, deviceResourceInt32, "", "")
	ValueIntx(t, deviceResourceInt64, "-9223372036854775808", "9223372036854775807")
	ValueIntx(t, deviceResourceInt64, "", "")
}

func TestValueUintx(t *testing.T) {
	ValueUintx(t, deviceResourceUint8, "0", "255")
	ValueUintx(t, deviceResourceUint8, "", "")
	ValueUintx(t, deviceResourceUint16, "0", "65535")
	ValueUintx(t, deviceResourceUint16, "", "")
	ValueUintx(t, deviceResourceUint32, "0", "4294967295")
	ValueUintx(t, deviceResourceUint32, "", "")
	ValueUintx(t, deviceResourceUint64, "0", "18446744073709551615")
	ValueUintx(t, deviceResourceUint64, "", "")
}

func TestValueFloatx(t *testing.T) {
	ValueFloatx(t, deviceResourceFloat32, "-3.40282346638528859811704183484516925440e+38", "3.40282346638528859811704183484516925440e+38")
	ValueFloatx(t, deviceResourceFloat32, "", "")
	ValueFloatx(t, deviceResourceFloat64, "-1.797693134862315708145274237317043567981e+308", "1.797693134862315708145274237317043567981e+308")
	ValueFloatx(t, deviceResourceFloat64, "", "")
}

func ValueIntx(t *testing.T, dr, minStr, maxStr string) {
	db := getDb()
	if err := db.openDb(); err != nil {
		t.Fatal(err)
	}
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

	rounds := 100

	min, _ := parseStrToInt(minStr, 64)
	max, _ := parseStrToInt(maxStr, 64)

	var i1 int64
	for x := 1; x <= rounds; x++ {
		vn, err := vd.read(deviceName, dr, minStr, maxStr, db)
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

	//generate read 100 times
	for x := 1; x <= rounds; x++ {
		v, err := vd.read(deviceName, dr, minStr, maxStr, db)

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

	v1, _ := vd.read(deviceName, dr, minStr, maxStr, db)
	i1 = getIntValue(v1)
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, dr, minStr, maxStr, db)
		i2 := getIntValue(v2)
		if i1 != i2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func ValueUintx(t *testing.T, dr, minStr, maxStr string) {
	db := getDb()
	if err := db.openDb(); err != nil {
		t.Fatal(err)
	}
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

	rounds := 100

	min, _ := parseStrToUint(minStr, 64)
	max, _ := parseStrToUint(maxStr, 64)

	var i1 uint64
	for x := 1; x <= rounds; x++ {
		vn, _ := vd.read(deviceName, dr, minStr, maxStr, db)
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

	//generate read 100 times
	for x := 1; x <= rounds; x++ {
		v, err := vd.read(deviceName, dr, minStr, maxStr, db)
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

	v1, _ := vd.read(deviceName, dr, minStr, maxStr, db)
	i1 = getUintValue(v1)
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, dr, minStr, maxStr, db)
		i2 := getUintValue(v2)
		if i1 != i2 {
			t.Fatalf("EnableRandomization is false, but got different read")
		}
	}
}

func ValueFloatx(t *testing.T, dr, minStr, maxStr string) {
	db := getDb()
	if err := db.openDb(); err != nil {
		t.Fatal(err)
	}
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

	rounds := 100

	min, _ := parseStrToFloat(minStr, 64)
	max, _ := parseStrToFloat(maxStr, 64)

	var f1 float64
	for x := 1; x <= rounds; x++ {
		vn, _ := vd.read(deviceName, dr, minStr, maxStr, db)
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

	//generate read 100 times
	for x := 1; x <= rounds; x++ {
		v, err := vd.read(deviceName, dr, minStr, maxStr, db)
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

	v1, _ := vd.read(deviceName, dr, minStr, maxStr, db)
	f1 = getFloatValue(v1)
	for x := 1; x <= rounds; x++ {
		v2, _ := vd.read(deviceName, dr, minStr, maxStr, db)
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
