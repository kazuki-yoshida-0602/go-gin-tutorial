package main

type article struct {
  ID      int    `json:"id"`
  Title   string `json:"title"`
  Content string `json:"content"`
}

func getAllArticles() []article {
  return articleList
}
