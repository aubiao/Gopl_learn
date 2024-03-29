package ChapterOne

import (
	"fmt"
	"io"
	"strings"

	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url,"http://")||strings.HasPrefix(url,"https://")){
			url = strings.Join([]string{"https://",url},"")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout,"aaaa\n")
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()

		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", b)
		fmt.Printf("%s\n", resp.Status)
	}
}
