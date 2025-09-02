package main

import (
	"encoding/json"
	"fmt"
)

const absoluteTS string = "15:04:05.000"

/* Print Traffic logs in JSON format to stdio for journal.d to pickup */
func PrintLogAsJSON(req_log RequestLog) {
	entry_json, err := json.Marshal(req_log)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s\n", string(entry_json))
}
