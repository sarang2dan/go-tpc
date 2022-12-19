package util

import (
	"errors"
	"fmt"
	"strings"
)

var DriverList = []string{
	"mysql",
	"postgres",
	"tidb",
	"greenplum",
	"starrocks",
	"singlestore",
}

const (
	DrvIdxMySQL = iota + 0
	DrvIdxPostgres
	DrvIdxTiDB
	DrvIdxGreenPlum
	DrvIdxStarRocks
	DrvIdxSingleStore
	// add here for new type
	DrvIdxInvalid
)

type ProtocolType int

const (
	PtTypeMySQL ProtocolType = iota + 0
	PtTypePostgres
)

var protocolType = []ProtocolType{
	PtTypeMySQL,    // "mysql",
	PtTypePostgres, // "postgres",
	PtTypeMySQL,    // "tidb",
	PtTypePostgres, // "greenplum",
	PtTypeMySQL,    // "starrocks",
	PtTypeMySQL,    //"singlestore",
}

type DriverMeta string

func (d *DriverMeta) String() string {
	return string(*d)
}

func (d *DriverMeta) SetDriver(driver string) error {
	driver = strings.TrimSpace(strings.ToLower(driver))

	switch driver {
	case "mysql":
	case "tidb":
	case "starrocks":
	case "singlestore":
	case "postgres":
	case "greenplum":
	default:
		errStr := fmt.Sprintf("Driver is not valid (%s)", driver)
		return errors.New(errStr)
	}

	*d = DriverMeta(driver)
	return nil
}

func (d *DriverMeta) GetDriverIdx() int {
	switch *d {
	case "mysql":
		return DrvIdxMySQL
	case "tidb":
		return DrvIdxTiDB
	case "starrocks":
		return DrvIdxStarRocks
	case "singlestore":
		return DrvIdxSingleStore
	case "postgres":
		return DrvIdxPostgres
	case "greenplum":
		return DrvIdxGreenPlum
	default:
		return DrvIdxInvalid
	}

	//return DrvIdxInvalid
}

func (d *DriverMeta) GetProtocolType() ProtocolType {
	switch *d {
	case "mysql":
		fallthrough
	case "tidb":
		fallthrough
	case "starrocks":
		fallthrough
	case "singlestore":
		return PtTypeMySQL

	case "postgres":
		fallthrough
	case "greenplum":
		return PtTypePostgres
	}

	panic(*d)
}

func (d *DriverMeta) GetProtocolTypeString() string {
	switch d.GetProtocolType() {
	case PtTypeMySQL:
		return "mysql"
	case PtTypePostgres:
		return "postgres"
	}
	panic(*d)
}
