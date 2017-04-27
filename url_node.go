package main

import (
	// "fmt"
	"bytes"
	"os"
	"strings"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"encoding/json"
)

type UrlNode struct {
	BaseUrl *url.URL `json:"-"`
	UrlString string `json:"url"`
	// For deterministic ordering
	StaticUrls []string `json:"static_urls"`
	LinkedUrls []string `json:"linked_urls"`

	// For constant time de-duping
	staticUrlsMap map[string]bool
	linkedUrlsMap map[string]bool
}

func NewUrlNode(u string) (*UrlNode, error) {
	u = cleanupUrl(u)
	formattedUrl, err := url.ParseRequestURI(u)
	if err != nil {
		// TODO(riley): proper handling if bad URL given?
		return nil, err
	}
	node := UrlNode{
		BaseUrl: formattedUrl,
		UrlString: formattedUrl.String(),
		StaticUrls: make([]string, 0),
		LinkedUrls: make([]string, 0),
		staticUrlsMap: map[string]bool{},
		linkedUrlsMap: map[string]bool{},
	}
	return &node, nil
}

func (u *UrlNode) String() string {
	return u.BaseUrl.String()
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
				// Avoid embedded data
				if strings.HasPrefix(val, "http") {
					u.addStaticUrl(val)
				}
			case "script":
				val = getTokenKey(token, "src")
				u.addStaticUrl(val)
		}
	}
	// Step 3: Printing is called by the main goroutine
}

func (u *UrlNode) addLinkedUrl(foundUrlStr string) {
	if foundUrlStr == "" {
		return
	}
	foundUrlStr = cleanupUrl(foundUrlStr)
	foundUrl, err := url.Parse(foundUrlStr)
	if err != nil {
		return // Skip silently
	}
	resolvedUrl := u.BaseUrl.ResolveReference(foundUrl)
	if u.BaseUrl.Host != resolvedUrl.Host {
		return // Skip URLs on other domains
	}
	addUnique(&u.LinkedUrls, &u.linkedUrlsMap, resolvedUrl.String())
}

func (u *UrlNode) addStaticUrl(foundUrlStr string) {
	if foundUrlStr == "" {
		return
	}
	foundUrl, err := url.Parse(foundUrlStr)
	if err != nil {
		return // Skip silently
	}
	resolvedUrl := u.BaseUrl.ResolveReference(foundUrl)
	addUnique(&u.StaticUrls, &u.staticUrlsMap, resolvedUrl.String())
}

func (u *UrlNode) PrintResults() {
	// fmt.Println("printing")
	// fmt.Println(u.LinkedUrls)
	outJson, err := json.Marshal(u)
	if err != nil {
		// Don't know why this would fail
		panic(err)
	}

	var out bytes.Buffer
	json.Indent(&out, outJson, "", "  ")
	out.WriteTo(os.Stdout)
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

// Remove anchors and trailing slashes
func cleanupUrl(urlStr string) string {
	return strings.TrimRight(strings.Split(urlStr, "#")[0], "/")
}
