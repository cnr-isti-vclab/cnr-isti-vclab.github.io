//import json package to parse json file
package main

//fix package command-line-arguments is not a main package problem
import (
	"encoding/json"
	"text/template"
	"strings"	
	"io/ioutil"
	"log"
	"os"
)

//define class Activity
type Activity struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Authors  string `json:"authors"`
	Years    string `json:"years"`
	Url      string `json:"url"`
	Body     string `json:"body"`
	Tags     string `json:"tags"`
}




//open and parse the file activities.json
func GetActivities() []Activity {
	file, err := os.Open("activities.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var activities []Activity
	json.Unmarshal(byteValue, &activities)

	return activities
}

//define a template string for the index.md file
const indexTemplate = `---
title: "{{.Title}}"
summary: "[{{.Years}}] {{.Subtitle}} <p onclick='this.style.display=\"block\"; event.preventDefault();' style='overflow: hidden; display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical;'>{{.Body}}</p>"
authors: [{{.Authors}}]
tags: 
{{.Tags}}

show_date: false

external_link: "{{.Url}}"

image:
  caption: ""
  focal_point: ""
  preview_only: false

url_code: ""
url_pdf: ""
url_slides: ""
url_video: ""

slides: ""
---
<p>{{.Body}}</p>
<button onclick="console.log('a')">Show More</button>
`

//create a folder for each entry in the activities.json file
func CreateFolder(activity Activity) {
	var folder = activity.Id
	err := os.Mkdir(folder, 0755)
	if err != nil {
		//just print a warning if the folder already exists
		log.Println(err)
	}
	//create the index.md file, overwrite in case it already exists
	file, err := os.Create(folder + "/index.md")
	if err != nil {
		log.Fatal(err)
	}
	//convert activity.tags string to an array
	var tmp = strings.Split(activity.Tags, ",")
	//map each tag to a string starting with -
	for i, tag := range tmp {
		tmp[i] = "- " + strings.TrimSpace(tag)
	}
	//join the tags array to a string, each line starts with -
	activity.Tags = strings.Join(tmp, "\n")


	tmpl, err := template.New("activity").Parse(indexTemplate)
	if err != nil { panic(err) }
	err = tmpl.Execute(file, activity)
	if err != nil { panic(err) }

	//print index to stdout

	
	/*_, err = file.WriteString(index)
	if err != nil {
		log.Fatal(err)
	} */
	defer file.Close() 
}

//main function to create a folder for each entry in the activities.json file
func main() {
	activities := GetActivities()

	for _, activity := range activities {
		CreateFolder(activity)
	}
}



