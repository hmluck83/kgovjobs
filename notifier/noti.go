package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hmluck83/kgovjobs/recruiment"
)

type notifi struct {
	Text       string `json:"text"`
	Chat_id    string `json:"chat_id"`
	Parse_mode string `json:"parse_mode"`
}

func sendMessage(s string) {
	noti := notifi{
		Text:	   s,
		// 내 채널 아이디임  
		Chat_id:   "-1002020654979",
		Parse_mode: "html",
	}

	jsonNoti, err := json.Marshal(noti)
	if err != nil {
		fmt.Println(err)
		return
	}

	// URL Building
	// TODO 존나 비효율적
	botAPI := os.Getenv("TBOT_API")
	telegramURL := "https://api.telegram.org/bot" + botAPI + "/sendmessage"

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", telegramURL, bytes.NewBuffer(jsonNoti))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 헤더 설정
	req.Header.Set("Content-Type", "application/json")

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// HTTP 요청 전송
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 응답 데이터 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 응답 출력
	fmt.Println(string(body))
}

func Send(eF *[]recruiment.Element) {
	var sb strings.Builder

	for i := range *eF {
		sb.WriteString((*eF)[i].HTMLString())
		if i % 20 == 0 {
			sendMessage(sb.String())
			sb.Reset()
		}
	}

	if sb.Len() > 0 {
		sendMessage(sb.String())
	}

}


