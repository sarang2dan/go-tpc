package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	TblLineItem = iota + 0
	TblPartSupp
	TblSupplier
	TbTbllPart
	TblOrders
	TblCustomer
	TblRegion
	TblNation
)

type TableDDL struct {
	Lineitem []string `json:"lineitem"`
	PartSupp []string `json:"partsupp"`
	Supplier []string `json:"supplier"`
	Part     []string `json:"part"`
	Orders   []string `json:"orders"`
	Customer []string `json:"customer"`
	Region   []string `json:"region"`
	Nation   []string `json:"nation"`
	nnn      []string `json:"dummy"`
}

// "./conf/ddl/{driver}/{bmtType}.json"
// "./conf/ddl/{bmtType}/{driver}.json"
// func GetDDLQueries(drv DriverMeta, bmtType BmtType) (*TableDDL, error) {
func GetDDLQueries(drv string, bmtType string) (*TableDDL, error) {
	cwd, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}

	ddlJsonPath := fmt.Sprint(
		cwd,
		string(os.PathSeparator),
		"conf",
		string(os.PathSeparator),
		"ddl",
		string(os.PathSeparator),
		bmtType,
		string(os.PathSeparator),
		drv,
		".json",
	)

	ddlJsonFile, err := os.Open(ddlJsonPath)
	if err != nil {
		errStr := fmt.Sprintf("%s(%s)", err.Error(), ddlJsonPath)
		return nil, errors.New(errStr)
	}
	defer ddlJsonFile.Close()

	decoder := json.NewDecoder(ddlJsonFile)

	var tblddl = &TableDDL{}
	err = decoder.Decode(&tblddl)
	if err != nil {
		return nil, err
	}

	return tblddl, nil
}
