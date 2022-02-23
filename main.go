package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Publisher struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int32  `json:"age,omitempty"`
}

type Review struct {
	Comment string `json:"comment,omitempty"`
	Rating  int    `json:"rating,omitempty"`
}

type Article struct {
	ArticleId   int32     `json:"article_id,omitempty"`
	ArticleName string    `json:"article_name,omitempty"`
	Publisher   Publisher `json:"publisher,omitempty"`
	Reviews     []Review  `json:"reviews,omitempty"`
}

type Articles struct {
	Articles []Article
}

func GetJson() *Articles {

	file, err := os.Open("./articles.json")
	if err != nil {
		panic(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	var articles Articles
	if err := json.NewDecoder(file).Decode(&articles); err != nil {
		panic(err.Error())
	}
	return &articles
}

func (articles *Articles) PrintRaw() {
	fmt.Printf("Printed Struct:%+v\n\n", articles)
}

func (articles *Articles) PrintJson() {
	strJson, err := json.MarshalIndent(articles, "", " ")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(string(strJson))
	}
}

func (article *Article) PrintJson() {
	strJson, err := json.MarshalIndent(article, "", " ")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Printed Json Payload: %s\n", strJson)
	}
}

func (articles *Articles) GetHighestRated() (*Article, float32) {

	var bestRating float32
	var bestArticle Article
	for _, article := range articles.Articles {
		var tempRating float32
		var tempArticle Article
		if article.GetAverage() > bestRating {
			tempRating = article.GetAverage()
			tempArticle = article
		} else {
			continue
		}
		bestRating = tempRating
		bestArticle = tempArticle
	}
	return &bestArticle, bestRating
}

func (article *Article) GetAverage() float32 {
	var rating float32
	for _, review := range article.Reviews {
		rating += float32(review.Rating)
	}
	return rating / float32(len(article.Reviews))
}

func main() {
	articles := GetJson()
	articles.PrintRaw()
	articles.PrintJson()
	article, rating := articles.GetHighestRated()
	fmt.Printf("Best Rating%.2f\n", rating)
	article.PrintJson()
}
