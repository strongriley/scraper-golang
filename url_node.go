package main

import (
	"fmt"
	"bytes"
	"os"
	"strings"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"encoding/json"
)

type UrlNode struct {
	url *url.URL `json:"url"`
	// For deterministic ordering
	staticUrls []string `json:"static_urls"`
	linkedUrls []string

	// For constant time de-duping
	staticUrlsMap map[string]bool
	linkedUrlsMap map[string]bool
}

type PrintUrl struct {
	Url string  `json:"url"`
	StaticAssets []string  `json:"static_assets"`
}

func NewUrlNode(u string) (*UrlNode, error) {
	// TODO(riley): Don't repeat yourself. See parsing below
	u = strings.TrimRight(strings.Split(u, "#")[0], "/")
	formattedUrl, err := url.ParseRequestURI(u)
	if err != nil {
		// TODO(riley): proper handling if bad URL given?
		return nil, err
	}
	node := UrlNode{
		url: formattedUrl,
		staticUrls: []string{},
		linkedUrls: []string{},
		staticUrlsMap: map[string]bool{},
		linkedUrlsMap: map[string]bool{},
	}
	return &node, nil
}

func (u *UrlNode) String() string {
	return u.url.String()
}

func (u *UrlNode) Process() {
	// Step 1: HTTP Request -----------
	resp, err := http.Get(u.String())
	if err != nil {
		return // Fail silently
	}
	bod := resp.Body
	defer bod.Close()
	page := html.NewTokenizer(bod)

	// Step 2: HTML Parsing -----------
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := page.Token()
		var val string
		switch token.Data {
			case "a":
				val = getTokenKey(token, "href")
				u.addLinkedUrl(val)
			case "link":
				val = getTokenKey(token, "href")
				rel := getTokenKey(token, "rel")
				if rel == "stylesheet" {
					u.addStaticUrl(val)
				}
			case "img":
				val = getTokenKey(token, "src")
				u.addStaticUrl(val)
			case "script":
				val = getTokenKey(token, "src")
				u.addStaticUrl(val)
		}
	}
	// Step 3: Printing is called by the main thread
}

func (u *UrlNode) addLinkedUrl(foundUrlStr string) {
	if foundUrlStr == "" {
		return
	}
	// Remove anchors and trailing slashes
	foundUrlStr = strings.TrimRight(strings.Split(foundUrlStr, "#")[0], "/")
	foundUrl, err := url.Parse(foundUrlStr)
	if err != nil {
		return // Skip silently
	}
	resolvedUrl := u.url.ResolveReference(foundUrl)
	if u.url.Host != resolvedUrl.Host {
		return // Skip URLs on other domains
	}
	addUnique(&u.linkedUrls, &u.linkedUrlsMap, resolvedUrl.String())
}

func (u *UrlNode) addStaticUrl(foundUrlStr string) {
	if foundUrlStr == "" {
		return
	}
	foundUrl, err := url.Parse(foundUrlStr)
	if err != nil {
		return // Skip silently
	}
	resolvedUrl := u.url.ResolveReference(foundUrl)
	addUnique(&u.staticUrls, &u.staticUrlsMap, resolvedUrl.String())
}

func (u *UrlNode) PrintResults() {
	outObj := &PrintUrl{
		Url: u.url.String(),
		StaticAssets: u.staticUrls,
	}
	outJson, err := json.Marshal(outObj)
	if err != nil {
		// Don't know why this would fail
		panic(err)
	}

	var out bytes.Buffer
	json.Indent(&out, outJson, "", "  ")
	out.WriteTo(os.Stdout)
	fmt.Print("\n")
}

// Used for both addLinkedUrl and addStaticUrl. DRY
func addUnique(urlList *[]string, urlMap *map[string]bool, newUrl string) {
	_, exists := (*urlMap)[newUrl]
	if exists {
		return
	}
	(*urlMap)[newUrl] = true
	*urlList = append(*urlList, newUrl)
}

func getTokenKey(token html.Token, key string) string {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
