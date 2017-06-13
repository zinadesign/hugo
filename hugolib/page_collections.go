// Copyright 2016 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"fmt"
	"sync"
	"time"
	"github.com/zinadesign/hugo/helpers"
	"strings"
)


// PageCollections contains the page collections for a site.
type PageCollections struct {
	// Includes only pages of all types, and only pages in the current language.
	Pages Pages
	PagesByUrl map[string]*Page

	// Includes all pages in all languages, including the current one.
	// Inlcudes pages of all types.
	AllPages Pages

	// A convenience cache for the traditional index types, taxonomies, home page etc.
	// This is for the current language only.
	indexPages Pages

	// A convenience cache for the regular pages.
	// This is for the current language only.
	RegularPages Pages

	// A convenience cache for the all the regular pages.
	AllRegularPages Pages

	// Includes absolute all pages (of all types), including drafts etc.
	rawAllPages Pages
}

func (c *PageCollections) refreshPageCaches() {
	c.indexPages = c.findPagesByKindNotIn(KindPage, c.Pages)
	c.RegularPages = c.findPagesByKindIn(KindPage, c.Pages)
	c.AllRegularPages = c.findPagesByKindIn(KindPage, c.AllPages)
}

func newPageCollections() *PageCollections {
	return &PageCollections{}
}

func newPageCollectionsFromPages(pages Pages) *PageCollections {
	return &PageCollections{rawAllPages: pages}
}

func (c *PageCollections) getPage(typ string, path ...string) *Page {
	pages := c.findPagesByKindIn(typ, c.Pages)

	if len(pages) == 0 {
		return nil
	}

	if len(path) == 0 && len(pages) == 1 {
		return pages[0]
	}

	for _, p := range pages {
		match := false
		for i := 0; i < len(path); i++ {
			if len(p.sections) > i && path[i] == p.sections[i] {
				match = true
			} else {
				match = false
				break
			}
		}
		if match {
			return p
		}
	}

	return nil
}

func (*PageCollections) findPagesByKindIn(kind string, inPages Pages) Pages {
	var pages Pages
	for _, p := range inPages {
		if p.Kind == kind {
			pages = append(pages, p)
		}
	}
	return pages
}

func (*PageCollections) findPagesByKindNotIn(kind string, inPages Pages) Pages {
	var pages Pages
	for _, p := range inPages {
		if p.Kind != kind {
			pages = append(pages, p)
		}
	}
	return pages
}

func (c *PageCollections) findPagesByKind(kind string) Pages {
	return c.findPagesByKindIn(kind, c.Pages)
}

var (
	pageUrlMapLock sync.RWMutex
)

func (c *PageCollections) findPageByUrl(url string) (Page, error) {
	pageUrlMapLock.Lock()
	//if strings.Contains(url, "about") {
	//	print(url)
	//}
	defer pageUrlMapLock.Unlock()
	if len(c.PagesByUrl) == 0 {
		c.PagesByUrl = make(map[string]*Page)
		for _, p := range c.AllPages {
			c.PagesByUrl[p.URL()] = p
		}
	}
	if custom_url, ok := custom_taxonomy_urls[url]; ok {
		url = custom_url
	}
	res_page, ok := c.PagesByUrl[url]
	if ok {
		return *res_page, nil
	}
	var page Page
	return page, fmt.Errorf("Page with url %s not found", url)
}
type TaxonomyTermInfo struct {
	Title string
	Date time.Time
	Weight int
	URL string
}

func (c *PageCollections) GetTermInfo(taxonomy_name string, term_name string) (TaxonomyTermInfo) {
	term_name_url_encoded := helpers.CurrentPathSpec().URLize(term_name)
	url := fmt.Sprintf("/%s/%s/", taxonomy_name, term_name_url_encoded)
	if p, ok := c.findPageByUrl(url); ok == nil {
		return TaxonomyTermInfo{Title: p.Title, URL: p.URL()}
	}
	return TaxonomyTermInfo{Title: term_name, URL: url}
}
func (c *PageCollections) PageIsActive(current_page_url, taget_page_url string) (bool, error) {
	if taget_page_url == current_page_url {
		return true, nil
	}
	if taget_page_url == "/" || strings.HasSuffix(current_page_url, "/404/") {
		return false, nil
	}
	current_page, ok := c.findPageByUrl(current_page_url)
	if ok != nil {
		return false, ok
	}
	breadcrumbs, _ := current_page.Breadcrumbs()
	for _, breadcrumb := range breadcrumbs {
		if breadcrumb.URL == taget_page_url {
			return true, nil
		}
	}
	return false, nil
}

func (c *PageCollections) addPage(page *Page) {
	c.rawAllPages = append(c.rawAllPages, page)
}

func (c *PageCollections) removePageByPath(path string) {
	if i := c.rawAllPages.FindPagePosByFilePath(path); i >= 0 {
		c.rawAllPages = append(c.rawAllPages[:i], c.rawAllPages[i+1:]...)
	}
}

func (c *PageCollections) removePage(page *Page) {
	if i := c.rawAllPages.FindPagePos(page); i >= 0 {
		c.rawAllPages = append(c.rawAllPages[:i], c.rawAllPages[i+1:]...)
	}
}

func (c *PageCollections) replacePage(page *Page) {
	// will find existing page that matches filepath and remove it
	c.removePage(page)
	c.addPage(page)
}
