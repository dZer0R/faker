package data

// Computer consists of computer information
var Computer = map[string][]string{
	"linux_processor":  {"i686", "x86_64"},
	"mac_processor":    {"Intel", "PPC", "U; Intel", "U; PPC"},
	"windows_platform": {"Windows NT 10.0", "Windows NT 6.3", "Windows NT 6.2", "Windows NT 6.1"},
	"windows_version":  {"", "; WOW64", "; Win64; x64"},
}
