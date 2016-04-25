package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputPath = flag.String("in", "", "input json file path")
)

type Dataset struct {
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Tags     []Tag     `json:tags`
	Comments []Comment `json:comments`
}

type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:versions`
}

type User struct {
	Id              string `json:id`
	PermanentId     int    `json:permanent_id`
	ProfileImageUrl string `profile_image_url`
}

type Comment struct {
	Body      string `json:body`
	CreatedAt string `json:created_at`
	UpdateAt  string `json:updated_at`
	User      User   `json:user`
}

type Article struct {
	Results []Dataset `json:"articles"`
}

func main() {
	flag.Parse()

	if !Exsits(*inputPath) {
		fmt.Printf("not exist input file path : %s \n", *inputPath)
		return
	}

	// Loading jsonfile
	file, err := ioutil.ReadFile(*inputPath)
	if err != nil {
		panic(err)
	}

	var article Article
	jsonErr := json.Unmarshal(file, &article)
	if jsonErr != nil {
		panic(jsonErr)
	}
	for _, result := range article.Results {

		fileName := strings.Replace(result.Title, "/", "", -1)
		fw, err := os.Create("./output/" + fileName + ".md")
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(fw)

		fmt.Fprint(w, OutputValue(result))
		w.Flush()
		fw.Close()
	}

	fmt.Println("処理が終了しました")
}

func Exsits(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func OutputValue(result Dataset) string {

	var buffer bytes.Buffer

	buffer.WriteString("#" + result.Title)
	buffer.WriteString("\n")
	buffer.WriteString(result.Body)
	buffer.WriteString("\n")
	if len(result.Comments) > 0 {
		buffer.WriteString("##Comments")
		buffer.WriteString("\n")
		for _, val := range result.Comments {
			buffer.WriteString(val.Body)
			buffer.WriteString("\n")
			buffer.WriteString(val.User.Id)
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
