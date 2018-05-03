package checker

import (
	"net/http"
	"encoding/json"
	"errors"
)

type Github struct {
}

type tag struct {
	Name string `json:"name"`
}

func NewGithub() *Github {
	return &Github{}
}

func (github *Github) Check(repo string,ignore ...string) (string,error) {
	resp,err:=http.Get("https://api.github.com/repos/"+repo+"/tags")
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	
	var tags []tag

	if err := json.NewDecoder(resp.Body).Decode(&tags);err!=nil{
		return "",err
	}

	t:
	for _,tag := range tags{
		for _,ignore := range ignore {
			if tag.Name == ignore {
				continue t
			}
		}
		return tag.Name,nil
	}
	return "",errors.New("tag list is null")
}