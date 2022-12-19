package util

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type BmtType int

const (
	BmtTypeTpcc BmtType = iota
	BmtTypeTpch
	BmtTypeTpcch
	BmtTypeRawsql
	// Add here for new type
	BmtTypeMAX
)

var bmtTypeStr = []string{
	"tpcc",
	"tpch",
	"tpcch",
	"rawsql",
}

func (b BmtType) String() string {
	if b >= BmtTypeMAX {
		log.Fatalf("Does not support type of test:(%d)", int(b))
	}
	return bmtTypeStr[b]
}

func (b *BmtType) Set(tp string) error {
	switch tp {
	case "tpcc":
		*b = BmtTypeTpcc
	case "tpch":
		*b = BmtTypeTpch
	case "tpcch":
		*b = BmtTypeTpcch
	case "rawsql":
		*b = BmtTypeRawsql
	default:
		errStr := fmt.Sprintf("Does not support type of test:(%d)", int(*b))
		return errors.New(errStr)
	}

	return nil
}
