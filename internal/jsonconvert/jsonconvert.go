package jsonconvert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type DataList struct {
	Data []Data `json:"data"`
}

type Data struct {
	RowCount int `json:"row_count"`
	Rows     []Row `json:"rows"`
}

type Row struct {
	CompanyID        int    `json:"company_id"`
	FundID           int    `json:"fund_id"`
	BranchID         int    `json:"branch_id"`
	CapitalAccountNo string `json:"capital_account_no"`
	BusinessTypeList string `json:"business_type_list"`
}

type RDataList struct {
	Data []RData `json:"data"`
}


type RData struct {
	RowCount int `json:"row_count"`
	Rows     []RRow `json:"rows"`
}

type RRow struct {
	Collection       string `json:"数据集名称"`
	CompanyID        int    `json:"company_id"`
	FundID           int    `json:"fund_id"`
	BranchID         int    `json:"branch_id"`
	CapitalAccountNo string `json:"capital_account_no"`
	BusinessTypeList string `json:"business_type_list"`
}

func Convert(filename string) {
	log.Println(filename)
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var datalist DataList
	json.Unmarshal(byteValue, &datalist)
	if len(datalist.Data) <= 0 {
		return
	}

	data := datalist.Data[0]
	count := data.RowCount

	// var rdatalist RDataList
	var rdata RData
	var row RRow

	if count <= 0 {
		return
	}

	rdata.RowCount = count
	for i := 0; i < count; i++ {
		row.Collection = fmt.Sprintf("数据集-%d", i)
		row.CompanyID = data.Rows[i].CompanyID
		row.FundID = data.Rows[i].FundID
		row.BranchID = data.Rows[i].BranchID
		row.BusinessTypeList = data.Rows[i].BusinessTypeList
		row.CapitalAccountNo = data.Rows[i].CapitalAccountNo

		rdata.Rows = append(rdata.Rows, row)
	}

	// rdatalist.Data = append(rdatalist.Data, rdata)

	res, err := json.Marshal(rdata.Rows)
	if err != nil {
		log.Fatal(err.Error())
	}

	destfilename := filepath.Dir(filename) + "\\jsonc.json"
	log.Println(destfilename)
	err = ioutil.WriteFile(destfilename, res, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
