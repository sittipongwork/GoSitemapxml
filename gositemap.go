package gositemap

import "encoding/xml"

/*
* [SitemapContainer] SitemapXML ROOT Document
* SitemapTag      {xml.Name `xml:"urlset"`}     Root tag xml for sitemap <urlset>
* XmlnsAttr       {string `xml:"xmlns,attr"`}   Attribute xmlns="{string}" in xml <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
 */
type SitemapContainer struct {
	XMLName   xml.Name  `xml:"urlset"`
	XmlnsAttr string    `xml:"xmlns,attr"`
	Url       []UrlSite `xml:"url"`
}

/*
* [UrlSite] Url of website in SitemapContainer
* Loc         {string `xml:"loc"`}        Url of site
* LastMod     {string `xml:"lastmod"`}    Date Last modify website format should like 2016-12-30 (YYYY-MM-DD) and String type
 */
type UrlSite struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

/*
* [CreateSitemapContainer] Create Container for keep sitemaps
* @param    sitemapversion   {string}     data sitemap version in xmlns e.g: "http://www.sitemaps.org/schemas/sitemap/0.9"
* @return   &SitemapContainer             Container of sitemap
 */
func CreateSitemapContainer(sitemapversion string) *SitemapContainer {
	return &SitemapContainer{XmlnsAttr: sitemapversion}
}

/*
* [SitemapContainer.AddSite] Add Site to container
* @param    site       {string}      website in loc
* @param    lastmod    {string}      date last modify website
*
* @append   SitemapContainer.Url    for keep data in SitemapContainer
* @return   No
 */
func (sm *SitemapContainer) AddSite(site string, lastmod string) {
	sm.Url = append(sm.Url, UrlSite{Loc: site, LastMod: lastmod})
}
