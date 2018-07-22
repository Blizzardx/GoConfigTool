package common

import (
	"errors"
	"strconv"
)

func Parser_int32(content string, attachmentValue *int32) error {
	result, err := strconv.ParseInt(content, 10, 32)
	if err != nil {
		return err
	}
	*attachmentValue = int32(result)
	return nil
}
func Parser_int64(content string, attachmentValue *int64) error {
	result, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		return err
	}
	*attachmentValue = result
	return nil
}
func Parser_float32(content string, attachmentValue *float32) error {
	result, err := strconv.ParseFloat(content, 32)
	if err != nil {
		return err
	}
	*attachmentValue = float32(result)
	return nil
}
func Parser_float64(content string, attachmentValue *float64) error {
	result, err := strconv.ParseFloat(content, 64)
	if err != nil {
		return err
	}
	*attachmentValue = result
	return nil
}
func Parser_bool(content string, attachmentValue *bool) error {
	if content == "true" {
		*attachmentValue = true
	} else if content == "false" {
		*attachmentValue = true
	} else {
		return errors.New("can't parser bool " + content)
	}
	return nil
}
func Parser_string(content string, attachmentValue *string) error {

	*attachmentValue = content
	return nil
}
