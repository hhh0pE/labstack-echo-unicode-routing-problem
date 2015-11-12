package main

import (
	"github.com/labstack/echo"
	"net/url"
	"fmt"
)

var template string

func DefaultHandler(c *echo.Context) error {
	unescaped, err := url.QueryUnescape(c.Request().URL.String())
	if err != nil {
		panic("Can't unescape URL")
	}

	fmt.Println("Original URL:", c.Request().URL.String())
	fmt.Println("Cyrillic URL:", unescaped)
	fmt.Println()

	c.HTML(200, `<h1>`+unescaped+`</h1>`+template)
	return nil
}

func main() {

	routes := []string{"/", "/ц/", "/у/", "/к/", "/г/", "/ш/", "/щ/", "/х/", "/ъ/", "/ф/", "/ы/", "/р/", "/э/", "/я/", "/ч/", "/с/", "/т/", "/ь/", "/тарифы/", "/калькулятор/", "/киев/", "/харьков/"}
	error_routes := []string{"/й/", "/е/", "/н/", "/з/", "/в/", "/а/", "/п/", "/о/", "/л/", "/д/", "/ж/", "/м/", "/н/", "/и/", "/б/", "/ю/", "/днепропетровск/", "/львов/", "/волынь/", "/одесса/", "/луцк/"}
	for _, route := range routes {
		template += `<a href="`+route+`">`+route+`</a><br />`
	}

	template += `<br /><br /><div>Must works but 404:<br />`
	for _, route := range error_routes {
		template += `<a href="`+route+`">`+route+`</a><br />`
	}
	template += `</div>`

	r := echo.New()
	r.Get("/", DefaultHandler)
	r.Get("/города/", DefaultHandler)

	t := r.Group("/тарифы")
	t.Get("/", DefaultHandler)

	c := r.Group("/калькулятор")
	c.Get("/", DefaultHandler)

	city := r.Group("/:city")
	city.Get("/", DefaultHandler)

	r.Run(":9090")

}
