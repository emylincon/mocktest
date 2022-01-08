package main

import (
	"fmt"
	"log"
)

func getGithubUser(url string, username string) (User, error) {
	url = fmt.Sprintf("%s/%s", url, username)
	body, err := Requests(url)
	if err != nil {
		log.Println("request error: ", err)
		return User{}, err
	}
	user, err := convertArticleToStruct(body)
	if err != nil {
		log.Println("convertion error: ", err)
		return User{}, err
	}
	return *user, nil
}

func getGithubName(url string, username string) (name string, err error) {
	user, err := getGithubUser(url, username)
	if err != nil {
		return name, err
	} else {
		return user.Name, nil
	}
}

func getGithubUserCompany(url string, username string) (name string, err error) {
	user, err := getGithubUser(url, username)
	if err != nil {
		return name, err
	} else {
		return user.Company, nil
	}

}

func main() {
	url := "https://api.github.com/users"
	username := "emylincon"
	user, _ := getGithubUser(url, username)
	log.Printf("\nName: %s \nCompany: %s \nLocation: %s\n", user.Name, user.Company, user.Location)
	log.Println(getGithubName(url, username))
	log.Println(getGithubUserCompany(url, username))
}
