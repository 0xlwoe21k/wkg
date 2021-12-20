package conf

type Config struct {
	Scan *Scan

	DnsLogUrl string
}

type Scan struct {
	VulnScanRate 	int
	DomainScanRate 	int
	WebSiteScanRate int
	IPsScanRate		int
}

