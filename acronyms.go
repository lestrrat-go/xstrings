package xstrings

import "strings"

var acronyms map[string]string

func init() {
	// Copied from golint (`commonInitialisms`)
	names := []string{
		"ACL",
		"API",
		"ASCII",
		"CPU",
		"CSS",
		"DNS",
		"EOF",
		"GUID",
		"HTML",
		"HTTP",
		"HTTPS",
		"ID",
		"IP",
		"JSON",
		"LHS",
		"QPS",
		"RAM",
		"RHS",
		"RPC",
		"SLA",
		"SMTP",
		"SQL",
		"SSH",
		"TCP",
		"TLS",
		"TTL",
		"UDP",
		"UI",
		"UID",
		"UUID",
		"URI",
		"URL",
		"UTF8",
		"VM",
		"XML",
		"XMPP",
		"XSRF",
		"XSS",
	}

	acronyms = make(map[string]string)
	for _, s := range names {
		// We edit the keys to Look Like This (uppercase first) which is how we will be querying it
		acronyms[UcFirst(strings.ToLower(s))] = s
	}
}
