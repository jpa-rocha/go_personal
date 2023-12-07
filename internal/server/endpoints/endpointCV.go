package endpoints

import (
	"adamastor/public"
	"bufio"
	"bytes"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

type Header struct {
    Slug string
    Institution string
    Location string
    Country string
    StartDate string
    EndDate string
    Position string
    Type string
}

func (h *Handler) HandleCV(w http.ResponseWriter, _ *http.Request) {
    // positions := []Header{}
    
	const base = "assets/static/content"
	content, err := fs.Sub(public.Assets, base)
	if err != nil {
		log.Println(err)
		return
	}

    files, err := fs.ReadDir(content, ".")
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
         fileContent, err := fs.ReadFile(content, file.Name())
        if err != nil {
            log.Fatalf("error reading file: %v", err)
        }

        scanner := bufio.NewScanner(bytes.NewReader(fileContent))
        headerConverter := map[string]string{}
        meta := false
        for scanner.Scan() {
            line := scanner.Text()
            // Get Header
            if line == "" {
                continue
            }
            if line == "---" {
                if !meta {
                    meta = true
                    continue
                } else {
                    break;
                }
            }
            if meta {
                split := strings.Split(line, ":")
                headerConverter[split[0]] = strings.Trim(split[1], " ")
            }
            
        }
        header := Header{
            Slug: headerConverter["slug"],
            Institution: headerConverter["institution"],
            Location: headerConverter["location"],
            Country: headerConverter["country"],
            StartDate: headerConverter["start_date"],
            EndDate: headerConverter["end_date"],
            Position: headerConverter["position"],
            Type: headerConverter["type"],

        }
        log.Println(header)
	}

	err = h.t.RenderTemplate(w, "cv.html", nil)
	if err != nil {
		log.Println(err)
	}
}
