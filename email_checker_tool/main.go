package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("domain, has MX, has SPF, spr Record", "has DMARC", "dmarc Record")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
		os.Exit(1)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarc, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("ErrorL %v\n", err)
	}

	for _, record := range dmarc {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println(domain, hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)
}