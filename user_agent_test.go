package faker

import (
	"fmt"
	"testing"
)

func ExampleUserAgent() {
	Seed(11)
	fmt.Println(UserAgent())
	// Output: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36
}

func BenchmarkUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserAgent()
	}
}

func TestUserAgent(t *testing.T) {
	for i := 0; i < 100; i++ {
		UserAgent()
	}
}

func ExampleChromeUserAgent() {
	Seed(11)
	fmt.Println(ChromeUserAgent())
	// Output: Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36
}

func BenchmarkChromeUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChromeUserAgent()
	}
}

func ExampleFirefoxUserAgent() {
	Seed(11)
	fmt.Println(FirefoxUserAgent())
	// Output: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_11_2; rv:55.0) Gecko/20100101 Firefox/55.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirefoxUserAgent()
	}
}

func ExampleSafariUserAgent() {
	Seed(11)
	fmt.Println(SafariUserAgent())
	// Output: Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6
}

func BenchmarkSafariUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafariUserAgent()
	}
}

func ExampleOperaUserAgent() {
	Seed(11)
	fmt.Println(OperaUserAgent())
	// Output: Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OperaUserAgent()
	}
}
