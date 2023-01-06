package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type GitStatus struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		ID                 string      `json:"id"`
		Name               string      `json:"name"`
		Status             string      `json:"status"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		Position           int         `json:"position"`
		Description        string      `json:"description"`
		Showcase           bool        `json:"showcase"`
		StartDate          interface{} `json:"start_date"`
		GroupID            interface{} `json:"group_id"`
		PageID             string      `json:"page_id"`
		Group              bool        `json:"group"`
		OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
	} `json:"components"`
}

func main() {

	results := GitStatus{}

	response, err := http.Get("https://www.githubstatus.com/api/v2/components.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(body, &results)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	for _, result := range results.Components {
		if result.Status != "operational" {
			fmt.Printf("%#v current status is %v\n", result.Name, result.Status)
		}
	}

	//fmt.Println(results)

}
