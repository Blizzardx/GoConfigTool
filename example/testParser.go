package main

import (
	"errors"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/common"
	"strconv"
	"strings"
)

type itemLine struct {
	Id   int32
	Name string
	Ids  []int32
}
type itemConfigTable struct {
	Content map[int32]*itemLine
}

func ParserConfig_itemConfigTable(decoder common.ConfigDecoder, configContent [][]string) ([]byte, error) {
	table := &itemConfigTable{}
	table.Content = map[int32]*itemLine{}

	for line, lineContent := range configContent {
		lineElem, err := ParserLine_itemConfigTable(lineContent)
		if nil != err {
			str := fmt.Sprintf("error on load config itemConfigTable at line: " + strconv.Itoa(line+1) + " " + err.Error())
			return nil, errors.New(str)
		}
		if v, ok := table.Content[lineElem.Id]; ok {
			str := fmt.Sprintf("error on load config itemConfigTable at line: "+strconv.Itoa(line+1)+" key Id already in table ", v)
			return nil, errors.New(str)
		}
		table.Content[lineElem.Id] = lineElem
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
	err = common.CheckValueLimit_int32(line.Id, "min", "max")
	if err != nil {
		return nil, errors.New(" column " + strconv.Itoa(columnIndex) + " named Id error on check min and max" + err.Error())
	}
	columnIndex++

	for _, lineElem := range strings.Split(lineContent[columnIndex], "|") {
		var elem int32
		err = common.Parser_int32(lineElem, &elem)
		if nil != err {
			return nil, errors.New(" column " + strconv.Itoa(columnIndex) + err.Error())
		}

		err = common.CheckValueLimit_int32(elem, "min", "max")
		if err != nil {
			return nil, errors.New(" column " + strconv.Itoa(columnIndex) + " named Id error on check min and max" + err.Error())
		}
		line.Ids = append(line.Ids, elem)
	}
	columnIndex++

	return line, nil
}
