// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package scraper // import "miniflux.app/v2/internal/reader/scraper"

// List of predefined scraper rules (alphabetically sorted)
// domain => CSS selectors
var predefinedRules = map[string]string{
	"bbc.co.uk":            "div.vxp-column--single, div.story-body__inner, ul.gallery-images__list",
	"blog.cloudflare.com":  "div.post-content",
	"cbc.ca":               ".story-content",
	"darkreading.com":      "#article-main:not(header)",
	"developpez.com":       "div[itemprop=articleBody]",
	"dilbert.com":          "span.comic-title-name, img.img-comic",
	"explosm.net":          "div#comic",
	"financialsamurai.com": "article",
	"francetvinfo.fr":      ".text",
	"github.com":           "article.entry-content",
	"heise.de":             "header .article-content__lead, header .article-image, div.article-layout__content.article-content",
	"igen.fr":              "section.corps",
	"ikiwiki.iki.fi":       ".page.group",
	"ilpost.it":            ".entry-content",
	"ing.dk":               "section.body",
	"lapresse.ca":          ".amorce, .entry",
	"lemonde.fr":           "article",
	"lepoint.fr":           ".art-text",
	"lesjoiesducode.fr":    ".blog-post-content img",
	"lesnumeriques.com":    ".text",
	"linux.com":            "div.content, div[property]",
	"mac4ever.com":         "div[itemprop=articleBody]",
	"monwindows.com":       ".blog-post-body",
	"motorsport.com":       "h2.text-article-description, div.ms-entity-promo, div.ms-article-content > :not(.relatedContent)",
	"npr.org":              "#storytext",
	"oneindia.com":         ".io-article-body",
	"opensource.com":       "div[property]",
	"openingsource.org":    "article.suxing-popup-gallery",
	"osnews.com":           "div.newscontent1",
	"phoronix.com":         "div.content",
	"pitchfork.com":        "#main-content",
	"pseudo-sciences.org":  "#art_main",
	"quantamagazine.org":   ".outer--content, figure, script",
	"raywenderlich.com":    "article",
	"royalroad.com":        ".author-note-portlet,.chapter-content",
	"slate.fr":             ".field-items",
	"smbc-comics.com":      "div#cc-comicbody, div#aftercomic",
	"swordscomic.com":      "img#comic-image, div#info-frame.tab-content-area",
	"techcrunch.com":       "div.article-entry",
	"theoatmeal.com":       "div#comic",
	"theregister.com":      "#top-col-story h2, #body",
	"theverge.com":         "h2.inline:nth-child(2),h2.duet--article--dangerously-set-cms-markup,figure.w-full,div.duet--article--article-body-component",
	"turnoff.us":           "article.post-content",
	"universfreebox.com":   "#corps_corps",
	"version2.dk":          "section.body",
	"wdwnt.com":            "div.entry-content",
	"webtoons.com":         ".viewer_img,p.author_text",
	"wired.com":            "main figure, article",
	"zeit.de":              ".summary, .article-body",
	"zdnet.com":            "div.storyBody",
}
