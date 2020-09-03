package camMysql

import (
	"database/sql"
	"fmt"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
)

type MysqlComponent struct {
	component.Component
	config *MysqlComponentConfig
}

// init
func (comp *MysqlComponent) Init(configI camStatics.ComponentConfigInterface) {
	comp.Component.Init(configI)
	var ok bool
	comp.config, ok = configI.(*MysqlComponentConfig)
	if !ok {
		camStatics.App.Fatal("MysqlComponent", "invalid config")
		return
	}
}

// start
func (comp *MysqlComponent) Start() {
	comp.Component.Start()
	comp.test()
}

// stop
func (comp *MysqlComponent) Stop() {
	defer comp.Component.Stop()
}

func (comp *MysqlComponent) test() {
	db, err := sql.Open("mysql", "root:123456@/cam")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
