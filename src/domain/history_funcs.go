package domain

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func (h HistoryItemSlice) Filter(filter *Filter, list []string) (out HistoryItemSlice) {
	out = nil
	for i, historyItem := range h {
		if h.FilterEqual(filter, historyItem, list) {
			out = append(out, h[i])
		}
	}
	return out
}

func (h *HistoryItemSlice) FilterEqual(filter *Filter, record *HistoryItem, list []string) (out bool) {
	defer func() {
		if r := recover(); r != nil {
			out = true
		}
	}()
	for _, field := range list {
		query := filter.ByName(field)
		recordString := record.ByName(field)
		if !strings.Contains(recordString, query) {
			return false
		}
	}
	return true
}

func (hi *HistoryItem) ByName(name string) (out string) {
	valueField := reflect.ValueOf(hi).Elem().FieldByName(name)
	if valueField.CanInt() {
		return strconv.Itoa(int(valueField.Int()))
	}
	if valueField.CanFloat() {
		return fmt.Sprintf("%.3f", valueField.Float())
	}
	return valueField.String()
}

func (hi *HistoryItem) Contain(field string, query string) (out bool) {
	defer func() {
		if r := recover(); r != nil {
			out = true
		}
	}()
	value := hi.ByName(field)
	return strings.Contains(value, query)
}
