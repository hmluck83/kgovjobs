package retriever

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hmluck83/kgovjobs/recruiment"
)

func Retrieve() (*[]recruiment.Element, error) {
    data := url.Values{}
	data.Set("searchJobsecode", "020")
    data.Set("empmnsn", "0")
    data.Set("prgl", "apmList")
    data.Set("searchWorkareaname", "전체")
    data.Set("areanm", "전국")
    data.Set("menuNo", "401")
    data.Set("selMenuNo", "400")
    data.Set("isShowBtn", "N")
    data.Set("serachAreaClassCd", "00000")
    data.Set("searchWorkareacode", "00000")

    gojobsUrl := "https://www.gojobs.go.kr/apmList.do"

	// 채용정보가 저장된 요소
	rE := []recruiment.Element{}

	// 오늘
	today := time.Now().Truncate(24 * time.Hour).Add(-24 * time.Hour)

	breakFlag := false

	// 일단 제대로 작동되는지 확인하기 우해 50번만 강제로 돌림 '-'
	for i := 1; i < 50; i++ {
		data.Set("pageIndex", fmt.Sprintf("%d", i))

		fmt.Println("page: ", i)
		requestBody := bytes.NewBufferString(data.Encode())

		// HTTP 요청 생성
		req, err := http.NewRequest("POST", gojobsUrl, requestBody)
		if err != nil {
			return nil, err
		}

		// 요청 헤더 설정
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// HTTP 클라이언트 생성 및 요청 전송
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		// 200 OK가 아닌 경우 에러 출력
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			return nil, err
		}

		resp.Body.Close()
		
		doc.Find("#apmTbl tbody tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
			td := s.Find("td")
			el := recruiment.ElementNew(td.Eq(1).Text(), td.Eq(1).Find("a").AttrOr("href", ""), td.Eq(2).Text(), td.Eq(3).Text(), td.Eq(4).Text())
			if !el.AfterDay(today) {
				// 오늘을 기준으로 24시간 이전(어제)를 기준으로 비교 함에 따라
				// 24시간 전 보다 이후일 경우 오늘임 
				// fmt.Println("오늘자 자료")
				if el.BeforeDay(today) {
					breakFlag = true
					return false
				}
				rE = append(rE, *el)

			} 
			return true
		})

		if breakFlag {
			break
		}
	}

	return &rE, nil
}