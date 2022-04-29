package scanner

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"michaelknudsen.com/ipscanner/ipgenerator"
)

type Server struct {
	ip ipgenerator.Ip
}

func (sc *Server) formatIpIntoUrl() (string, string) {
	return "http://" + sc.ip.AsString(), "https://" + sc.ip.AsString()
}

func getFromProtocol(protocol string, ip string) string {
	res, err := http.Get(protocol + ip)
	if err != nil {
		return err.Error()
	} else {
		defer res.Body.Close()
		switch res.StatusCode {
		case 200:
			b, _ := ioutil.ReadAll(res.Body)
			return string(b)
		default:
			return strconv.FormatInt(int64(res.StatusCode), 10)
		}
	}
}

func (sc *Server) testForHttps() string {
	return getFromProtocol("https://", sc.ip.AsString())
}
func (sc *Server) testForHttp() string {
	return getFromProtocol("http://", sc.ip.AsString())
}

func Main() {
	scanner := Server{}
	scanner.ip.SetBytes([...]int{1, 1, 1, 1})
	http_res := scanner.testForHttp()
	f, _ := os.Create("testPages/http_1.1.1.1.html")
	f.Write([]byte(http_res))
	f.Close()

	http_res = scanner.testForHttps()
	f, _ = os.Create("testPages/https_1.1.1.1.html")
	f.Write([]byte(http_res))
	f.Close()

}
