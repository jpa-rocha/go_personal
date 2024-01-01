package utilities

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/fs"
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"

	"adamastor/public"
)

type CVArticle struct {
	Slug        string
	Institution string
	Location    string
	Country     string
	StartDate   string
	EndDate     string
	Position    string
	Type        string
	Body        templ.Component
}

func HandleCVArticles() []CVArticle {
	const base = "assets/static/content/cv"
	content, err := fs.Sub(public.Assets, base)
	if err != nil {
		log.Println(err)
		return []CVArticle{}
	}

	files, err := fs.ReadDir(content, ".")
	if err != nil {
		log.Println(err)
		return []CVArticle{}
	}
	articles := []CVArticle{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileContent, err := fs.ReadFile(content, file.Name())
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}

		scanner := bufio.NewScanner(bytes.NewReader(fileContent))
		headerConverter := fillArticle(scanner)
		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(headerConverter["body"]), &buf); err != nil {
			log.Fatalf("failed to convert markdown to HTML: %v", err)
		}
		body := Unsafe(buf.String())
		article := CVArticle{
			Slug:        headerConverter["slug"],
			Institution: headerConverter["institution"],
			Location:    headerConverter["location"],
			Country:     headerConverter["country"],
			StartDate:   headerConverter["start_date"],
			EndDate:     headerConverter["end_date"],
			Position:    headerConverter["position"],
			Type:        headerConverter["type"],
			Body:        body,
		}
		articles = append(articles, article)
	}

	return articles
}

func fillArticle(scanner *bufio.Scanner) map[string]string {
	meta := false
	headerConverter := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		// Get Header if it exists
		if line == "" {
			continue
		}
		if line == "---" {
			if !meta {
				meta = true
				continue
			} else {
				break
			}
		}
		if meta {
			split := strings.Split(line, ":")
			headerConverter[split[0]] = strings.Trim(split[1], " ")
		}
	}
	var body string
	for scanner.Scan() {
		line := scanner.Text()
		body += line + "\n"
	}
	headerConverter["body"] = body

	return headerConverter
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		_, err := io.WriteString(w, html)
		return err
	})
}

func SplitExp(exp []CVArticle, occType string) []CVArticle {
	ret := []CVArticle{}
	for _, en := range exp {
		if en.Type == occType {
			ret = append(ret, en)
		}
	}
	sort.Slice(ret, func(p, q int) bool {
		return ret[p].StartDate < ret[q].StartDate
	})
    slices.Reverse(ret)
	return ret
}
