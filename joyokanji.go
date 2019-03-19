package joyokanji

import "github.com/PuerkitoBio/goquery"

func List() ([]rune, error) {
	d, err := goquery.NewDocument("https://ja.wikipedia.org/wiki/%E5%B8%B8%E7%94%A8%E6%BC%A2%E5%AD%97%E4%B8%80%E8%A6%A7")
	if err != nil {
		return nil, err
	}

	var ks []rune
	tr := d.Find("#mw-content-text div table tbody tr")
	tr.Each(func(n int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() < 8 || td.Eq(7).Text() != "" {
			return
		}

		k := td.Eq(1).Find("a.extiw").Text()
		if k == "" {
			return
		}

		ks = append(ks, []rune(k)[0])
	})

	return ks, nil
}
