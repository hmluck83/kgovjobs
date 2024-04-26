package recruiment

import (
	"strings"
	"time"
)

type Element struct {
	posting       string
	organization  string
	postURL       string 
	postDate      time.Time
	deadLine      time.Time
}

func ElementNew(posting, postURL, organization, postDate, deadLine string) *Element {
	// TODO: 입력값 검증 필요
	postDateTime, _ := time.Parse(time.DateOnly, strings.TrimSpace(postDate))
	deadLineTime, _ := time.Parse(time.DateOnly, strings.TrimSpace(deadLine))
	
	var url string
	if len(postURL) < 36 {
		url = ""	
	} else {
		url = "https://www.gojobs.go.kr/apmView.do?empmnsn=" + strings.TrimSpace(postURL)[30:36]
	}

	return &Element{
		posting:      strings.TrimSpace(posting),
		organization: strings.TrimSpace(organization),
		postURL:      url,
		postDate:     postDateTime,
		deadLine:     deadLineTime,
	}
}

func (e *Element) String() string {
	return e.posting + " " + e.organization + " " + e.postURL + " " + e.postDate.String() + " " + e.deadLine.String()
}

func (e *Element) HTMLString() string {
	return "<a href=\"" + e.postURL + "\">" + e.posting + "</a>\n" + e.organization + "\n" + "\n\n"
}

func (e *Element) BeforeDay(n time.Time) bool {
	return e.postDate.Before(n) 
}

func (e *Element) AfterDay(n time.Time) bool {
	return e.postDate.After(n)
}