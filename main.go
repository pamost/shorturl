package main

import (
	"bufio"
	"fmt"
	"github.com/speps/go-hashids"
	"os"
	"strings"
)

var m map[string]string

type Shortener interface {
	Shorten() string
	Resolve() string
}

type Url struct {
	Value string
}

func (u Url) Shorten() string {
	hd := hashids.NewData()
	hd.Salt = u.Value
	hd.MinLength = 5
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{len(u.Value)})
	part := strings.ToLower(e)
	path := strings.Split(u.Value, "/")
	shortUrl := path[0] + "/" + part

	m[shortUrl] = u.Value
	return shortUrl
}

func (u Url) Resolve() string {
	longUrl := m[u.Value]
	return longUrl
}

func main() {
	m = make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" {  // break the loop if text == "q"
		fmt.Print("Generate short URL (enter - 1), if check long URL (enter - 2): ")
		scanner.Scan()
		text = scanner.Text()

		if text == "1"  {
			fmt.Print("Enter your long URL (without http(s)://): ")
			scanner.Scan()
			text = scanner.Text()
			//url := Url{text}
			//url.Shorten(url.Value)
			var s Shortener = Url{text}
			s.Shorten()
			fmt.Println("map:", m)
		}

		if text == "2"  {
			fmt.Print("Enter your short URL (without http(s)://): ")
			scanner.Scan()
			text = scanner.Text()
			//url := Url{text}
			//fmt.Println(url.Resolve(url.Value))
			var s Shortener = Url{text}
			fmt.Println(s.Resolve())
		}
	}

}
