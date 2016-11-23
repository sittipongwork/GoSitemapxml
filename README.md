# The sitemap.xml for golang , SEO (Search Engine Optimization)


### Installing
Use Golang CLI
```go
go get github.com/sittipongwork/GoSitemapxml
```
or

Github Repository

[https://github.com/sittipongwork/GoSitemapxml.git](https://github.com/sittipongwork/GoSitemapxml.git)

### Import Package
example : Import to go
```go
  import "github.com/sittipongwork/GoSitemapxml"
```

### Full Example include Create Sitemap Container, Addsite, Parse struct to xml
example
```go
  //Create Sitemap Container for keep data xml
  sitemap := gositemap.CreateSitemapContainer("https://www.sitemaps.org/schemas/sitemap/0.9")

  //Add Site you want to show in sitemap.xml
  sitemap.AddSite("https://website.com/", "2016-11-23")
  //You can add more one site : Example
  sitemap.AddSite("https://website.com/blog/1", "2016-11-23")
  sitemap.AddSite("https://website.com/blog/2", "2016-11-23")
  sitemap.AddSite("https://website.com/blog/3", "2016-11-23")
  sitemap.AddSite("https://website.com/contact", "2016-11-23")

  //Mashal XML Data
  output, _ := xml.MarshalIndent(sitemap, "  ", "    ")
  //Print Output
  os.Stdout.Write(output)

  //NOTE : in iris framework
  //you can run command , Dont MashalIndent
  //iris.XML(iris.StatusOK, sitemap)
```


output
```xml
<urlset xmlns="https://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://website.com/</loc>
    <lastmod>2016-11-23</lastmod>
  </url>
  <url>
    <loc>https://website.com/blog/1</loc>
    <lastmod>2016-11-23</lastmod>
  </url>
  <url>
    <loc>https://website.com/blog/2</loc>
    <lastmod>2016-11-23</lastmod>
  </url>
  <url>
    <loc>https://website.com/blog/3</loc>
    <lastmod>2016-11-23</lastmod>
  </url>
  <url>
    <loc>https://website.com/contact</loc>
    <lastmod>2016-11-23</lastmod>
  </url>
</urlset>
```
