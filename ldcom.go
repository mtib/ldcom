package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	number = flag.Int("num", 37, "number of ludum dare")
	uniqid = flag.Int("id", 0, "unique id of entry")
)

func main() {
	flag.Parse()
	resp, _ := http.Get(fmt.Sprintf("http://ludumdare.com/compo/ludum-dare-%d/?action=preview&uid=%d", *number, *uniqid))
	body, _ := ioutil.ReadAll(resp.Body)
	sb := string(body)
	resp.Body.Close()
	title := regexp.MustCompile("<h2 .*?>(.*?)</h2>").FindAllStringSubmatch(sb, 1)[0][1]
	comst := "<div class = 'comment'>"
	comen := "</p></div>"
	formst := "<p>You must sign in to comment.</p>"
	cb := sb[strings.Index(sb, comst)+len(comst) : strings.LastIndex(sb, formst)-len(comen)]
	comarr := strings.Split(cb, comen+comst)
	fmt.Println("#", title)
	fmt.Println("## Comments")
	// fmt.Printf("number of comments: %d\n", len(comarr))
	for k := range comarr {
		ns := strings.Replace(gtrim(comarr[k], "<", ">", "\n")[4:], "\n\n", "\n", -1)
		fmt.Println("### " + ns + "\n")
	}
}

func gtrim(s, lsep, rsep, nsep string) string {
	if len(s) == 0 {
		return ""
	}
	lin := strings.Index(s, lsep)
	rin := strings.Index(s, rsep)
	if lin == -1 {
		// done
		return s
	}
	if lin > rin {
		return "@ERROR@"
	}
	return s[:lin] + nsep + gtrim(s[rin+len(rsep):], lsep, rsep, nsep)
}
