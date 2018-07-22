package main

import (
	"errors"
	"github.com/Blizzardx/GoConfigTool/common"
	"strconv"
)

type itemLine struct {
	Id   int32
	Name string
}
type itemConfigTable struct {
	Content []*itemLine
}

func ParserConfig_itemConfigTable(decoder common.ConfigDecoder, configContent [][]string) ([]byte, error) {
	table := &itemConfigTable{}
	for line, lineContent := range configContent {
		lineElem, err := ParserLine_itemConfigTable(lineContent)
		if nil != err {
			return nil, errors.New("error on load config itemConfigTable at line: " + strconv.Itoa(line+1) + " " + err.Error())
		}
		table.Content = append(table.Content, lineElem)
	}
	content, err := decoder.Encode(table)
	if nil != err {
		return nil, err
	}
	return content.([]byte), nil
}
func ParserLine_itemConfigTable(lineContent []string) (*itemLine, error) {
	line := &itemLine{}
	var err error = nil
	columnIndex := 0

	err = common.Parser_int32(lineContent[columnIndex], &line.Id)
	if nil != err {
		return nil, errors.New(" column " + strconv.Itoa(columnIndex) + " named Id " + err.Error())
	}
	columnIndex++

	return line, nil
}
