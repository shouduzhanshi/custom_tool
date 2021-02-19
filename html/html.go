package html

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)
func GetHtmlTitle(data io.Reader) string {
	doc ,err := goquery.NewDocumentFromReader(data)
	if err ==nil{
		return doc.Find("head").Find("title").Text()
	}
	return ""
}
