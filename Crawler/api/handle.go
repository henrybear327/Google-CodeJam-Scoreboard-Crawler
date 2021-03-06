package api

import (
	"encoding/json"
	"log"
)

type handleSearchPayload struct {
	/*
		{"nickname":"henrybear327","scoreboard_page_size":10}
	*/
	Handle   string `json:"nickname"`
	PageSize int    `json:"scoreboard_page_size"`
}

func getHandleSearchPayload(handle string) string {
	payload := handleSearchPayload{Handle: handle, PageSize: 1}
	res, err := json.Marshal(payload)
	handleErr(err)
	return encodeToBase64(res)
}

func (data *ContestData) fetchHandleResult(handle string, ch chan userScore) {
	param := make([]interface{}, 1)
	param[0] = handle
	response := fetchAPIResponse(specificHandleType, data.contestID, param).(*scoreboardResponse)

	if len(response.UserScores) != 1 {
		log.Printf("Incorrect user count (%v). Should be 1", len(response.UserScores))
		ch <- userScore{isEmpty: true}
		return
	}

	ch <- response.UserScores[0]
}
