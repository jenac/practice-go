//go run github_cli.go -a=field
//go run github_cli.go -a=jenac


package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type GithubAccountInfo struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUsername   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func GithubUserStorager(name string) GithubAccountInfo {
	url := fmt.Sprintf("https://api.github.com/users/%s", name)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "My Golang Browser")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return GithubAccountInfo{}
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	var account GithubAccountInfo
	err = json.Unmarshal([]byte(result), &account)
	if err != nil {
		fmt.Println(err)
		return GithubAccountInfo{}
	}
	fmt.Println(string([]byte(result)))
	return account
}

func GithubUserFields() {
	var fields []string
	account := GithubAccountInfo{}
	s := reflect.TypeOf(&account).Elem()
	for i := 0; i < s.NumField(); i++ {
		fields = append(fields, s.Field(i).Name)
	}
	fieldsJson, _ := json.MarshalIndent(fields, "", "  ")
	fmt.Println(string(fieldsJson))
}

func FlagHelper() {
	var Account string
	flag.StringVar(&Account, "a", "what", "show gith hub account user info")
	flag.Parse()

	if Account == "field" {
		GithubUserFields()
		return
	} else {
		GithubUserStorager(Account)
		return
	}
}

func main() {
	FlagHelper()
}