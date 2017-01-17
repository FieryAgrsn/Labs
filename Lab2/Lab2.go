package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"math"
	//"reflect"
)

var (
	gtm sync.Mutex
)

type TablesInMemory *Table

type objectsTable struct{
	l int `json:"l"`
	w int `json:"w"`
	l_b int `json:"l_b"`
}

type Table struct {
	data map[string][3]int
	m sync.RWMutex
}

func NewTable(data map[string][3]int) *Table {
	return &Table{data: data,}
}

func DecodeJSON() *Table{
	//var tmp map[string][3]int
	key_val, err := ioutil.ReadFile("db/testdb")
	if err != nil {
		return nil 
	}
	var f map[string][3]int
    	err = json.Unmarshal(key_val, &f)
	fmt.Println(f)
    	if err != nil {
		return nil 
	}
	/*for k, v := range f {
		fmt.Println(v) 
		var m [3]int
		err = json.Unmarshal([]byte(v), &m)
		checkErr(err)
		fmt.Println(m) 
		tmp[k] = m
		//vv := m//.(map[string]interface{})
   		
   		//tmp[k] = objectsTable{l: int(vv["l"].(float64)), w: int(vv["w"].(float64)), l_b : int(vv["l_b"].(float64)),}
	}*/
    	t := NewTable(f)
	return t
}

func EncodeJSON(tablechan <-chan Table) {
	for {
		table := <-tablechan
		m := make(map[string][3]int, len(table.data))
		for k, v := range table.data {
			//tmp := objectsTable{l : v.l, w : v.w, l_b : v.l_b}
			//data, _ := json.Marshal(tmp)

			m[k] = [3]int{ v[0],  v[1],  v[2]}
			
		}
		jsonData, err := json.Marshal(m)
		checkErr(err)
		f, err := os.Create("db/testdb")
		checkErr(err)
		defer f.Close()
		_, err = f.Write(jsonData)
		checkErr(err)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getTable(tables *TablesInMemory) *Table {
	gtm.Lock()
	if (*tables) != nil {
		gtm.Unlock()
		return (*tables)
	}
	table := DecodeJSON()
	if table != nil {
		*tables = table
	}

	gtm.Unlock()
	return table
}

func exit(c net.Conn) {
	c.Write([]byte(string("Bye\n")))
	c.Close()
}

func help(c net.Conn) {
	c.Write([]byte(string("get [key]\n")))
	c.Write([]byte(string("set [key] [value]\n")))
	c.Write([]byte(string("del [key]\n")))
	c.Write([]byte(string("exit\n")))
}

func getKeys(c net.Conn, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 1  {
		table := getTable(tables)
		table.m.RLock()
		keys := make([]string, 0, len(table.data))//todo keys shood be objectsTable struct
	    for k := range table.data {
	        keys = append(keys, k)
	    }
		table.m.RUnlock()
    		c.Write([]byte("[" + strings.Join(keys, ", ") + "]" + "\n"))
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func getVal(c net.Conn, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 2 {
		table := getTable(tables)
		table.m.RLock()
		value, ok := table.data[query_split[1]]
		table.m.RUnlock()
		if ok {
			c.Write([]byte(string("L = " + string(value[0]) +"\n")))
			c.Write([]byte(string("South : " + string(value[1]) +"\n")))
			c.Write([]byte(string("W = " + string(value[2]) +"\n")))
		} else {
			c.Write([]byte(string("key does not exist\n")))
		}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}



func setVal(c net.Conn, tablechan chan<- Table, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 5 {
		table := getTable(tables)
		if table == nil{
			table = NewTable(make(map[string][3]int))
			*tables = table
		}
		table.m.Lock()
		valuel, _ := strconv.Atoi(query_split[2])
		valuew, _ := strconv.Atoi(query_split[3])
		valuel_b, _ := strconv.Atoi(query_split[4])
		values := [3]int{  valuel, valuew ,  valuel_b,}
		table.data[query_split[1]] = values
		table.m.Unlock()
		c.Write([]byte(string("OK\n")))
		tablechan <- *table
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func get_in_range (c net.Conn, tablechan chan <- Table, tables *TablesInMemory, query_split []string){
	if len(query_split) == 3{
		table := getTable(tables)
		table.m.RLock()
		value, _ := table.data[query_split[1]]
		table.m.RUnlock()
		f_val, _ := strconv.Atoi(query_split[2])
		var length float64
		for k, v := range table.data{
			length = math.Sqrt(float64(math.Pow(float64(value[1]-v[1]), 2))+float64(math.Pow(float64(value[0]-v[0]), 2)))
			if length <= float64(f_val){
				c.Write([]byte( k + "\n"))
			}
		}
	}
}

func delKey(c net.Conn, tablechan chan<- Table, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 2  {
		table := getTable(tables)
		
		table.m.RLock()
		_, ok := table.data[query_split[1]]
		table.m.RUnlock()
		if ok {
			table.m.Lock()
			delete(table.data, query_split[1])
			table.m.Unlock()
			c.Write([]byte(string("OK\n")))
			tablechan <- *table
		} else {
			c.Write([]byte(string("key does not exist\n")))
		}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func handleRequest(c net.Conn, tablechan chan<- Table, tables *TablesInMemory, query string) {
	query_split := strings.Fields(query)

	if len(query_split) >= 2 {
		switch strings.ToLower(query_split[0]) {
			case "set":
				setVal(c, tablechan, tables, query_split)
			case "get":
				getVal(c, tables, query_split)
			case "del":
				delKey(c, tablechan, tables, query_split)
			case "get_in_range":
				get_in_range(c, tablechan, tables, query_split)
			
			default:
				c.Write([]byte(string("Unknown command\n")))
		}
	} else if len(query_split) == 1 {
		switch strings.ToLower(query_split[0]) {
			case "exit":
				exit(c)
			case "help":
				help(c)
			case "keys":
				getKeys(c, tables, query_split)
			default:
				c.Write([]byte(string("Unknown command\n")))
			}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func handleConnection(c net.Conn, tablechan chan<- Table, tables *TablesInMemory) {
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if (err != nil) || (n == 0) {
			break
		} else {
			go handleRequest(c, tablechan, tables, string(buf[0:n]))
		}
	}
}

func main() {

	ln, err := net.Listen("tcp", ":2222")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	var tables TablesInMemory
	
	tablechan := make(chan Table)
	go EncodeJSON(tablechan)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ok")
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		
		go handleConnection(conn, tablechan, &tables)
	}
}
