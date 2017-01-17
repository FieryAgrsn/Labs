package main

import (
	"testing"
	"strings"
	"strconv"
)

func getTab(tables *TablesInMemory) *Table {
	if (*tables) != nil {
		gtm.Unlock()
		return (*tables)
	}
	table := DecodeJSON()
	if table != nil {
		*tables = table
	}
	return table
}

func set(tables *TablesInMemory, query_split []string) string {
	if len(query_split) == 5 {
		table := getTable(tables)
		if table == nil{
			table = NewTable(make(map[string][3]int))
			*tables = table
		}
		valuel, _ := strconv.Atoi(query_split[2])
		valuew, _ := strconv.Atoi(query_split[3])
		valuel_b, _ := strconv.Atoi(query_split[4])
		values := [3]int{  valuel, valuew ,  valuel_b,}
		table.data[query_split[1]] = values
		return "OK"
	} else {
		return "Error"
	}
}



func TestSet(t *testing.T) {
	var tables TablesInMemory
	query := "set Warsaw 11 28 0" 
	query_split := strings.Fields(query)
	if set(&tables, query_split) == "OK"{
		t.Log("Test: Set operation to existing file/nStatus: passed successful.")
	} else {
		t.Error("FAILED: Set operation in existing file.")
	}

	
}



