---
aliases:
- /layout/functions/
date: 2013-07-01
linktitle: Functions
menu:
  main:
    parent: layout
next: /templates/variables
prev: /templates/go-templates
title: Hugo Template Functions
weight: 20
---

Hugo uses the excellent Go html/template library for its template engine.
It is an extremely lightweight engine that provides a very small amount of
logic. In our experience, it is just the right amount of logic to be able
to create a good static website.

Go templates are lightweight but extensible. Hugo has added the following
functions to the basic template logic.

(Go itself supplies built-in functions, including comparison operators
and other basic tools; these are listed in the
[Go template documentation](http://golang.org/pkg/text/template/#hdr-Functions).)

## General

### isset
Return true if the parameter is set.
Takes either a slice, array or channel and an index or a map and a key as input.

e.g. `{{ if isset .Params "project_url" }} {{ index .Params "project_url" }}{{ end }}`

### echoParam
If parameter is set, then echo it.

e.g. `{{echoParam .Params "project_url" }}`

### eq
Return true if the parameters are equal.

e.g.

    {{ if eq .Section "blog" }}current{{ end }}

### first
Slices an array to only the first X elements.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    {{ range first 10 .Data.Pages }}
        {{ .Render "summary" }}
    {{ end }}

### where
Filters an array to only elements containing a matching value for a given field.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    {{ range where .Data.Pages "Section" "post" }}
       {{ .Content }}
    {{ end }}

*`where` and `first` can be stacked, e.g.:*

    {{ range first 5 (where .Data.Pages "Section" "post") }}
       {{ .Content }}
    {{ end }}

### in
Checks if an element is in an array (or slice) and returns a boolean.  The elements supported are strings, integers and floats (only float64 will match as expected).  In addition, it can also check if a substring exists in a string.

e.g.

    {{ if in .Params.tags "Git" }}Follow me on GitHub!{{ end }}

or

    {{ if in "this string contains a substring" "substring" }}Substring found!{{ end }}

### intersect
Given two arrays (or slices), this function will return the common elements in the arrays.  The elements supported are strings, integers and floats (only float64).

A useful example of this functionality is a 'similar posts' block.  Create a list of links to posts where any of the tags in the current post match any tags in other posts.

e.g.

    <ul>
    {{ $page_link := .Permalink }}
    {{ $tags := .Params.tags }}
    {{ range .Site.Recent }}
        {{ $page := . }}
        {{ $has_common_tags := intersect $tags .Params.tags | len | lt 0 }}
        {{ if and $has_common_tags (ne $page_link $page.Permalink) }}
            <li><a href="{{ $page.Permalink }}">{{ $page.Title }}</a></li>
        {{ end }}
    {{ end }}
    </ul>


## Math

<table class="table table-bordered">
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example</th>
</tr>
</thead>

<tbody>
<tr>
<td><code>add</code></td>
<td>Adds two integers.</td>
<td><code>{{add 1 2}}</code> → 3</td>
</tr>

<tr>
<td><code>sub</code></td>
<td>Subtracts two integers.</td>
<td><code>{{sub 3 2}}</code> → 1</td>
</tr>

<tr>
<td><code>mul</code></td>
<td>Multiplies two integers.</td>
<td><code>{{mul 2 3}}</code> → 6</td>
</tr>

<tr>
<td><code>div</code></td>
<td>Divides two integers.</td>
<td><code>{{div 6 3}}</code> → 2</td>
</tr>

<tr>
<td><code>mod</code></td>
<td>Modulus of two integers.</td>
<td><code>{{mod 15 3}}</code> → 0</td>
</tr>

<tr>
<td><code>modBool</code></td>
<td>Boolean of modulus of two integers.  <code>true</code> if modulus is 0.</td>
<td><code>{{modBool 15 3}}</code> → true</td>
</tr>
</tbody>
</table>


## Strings

### urlize
Takes a string and sanitizes it for usage in URLs, converts spaces to "-".

e.g. `<a href="/tags/{{ . | urlize }}">{{ . }}</a>`

### safeHtml
Declares the provided string as a "safe" HTML document fragment
so Go html/template will not filter it.  It should not be used
for HTML from a third-party, or HTML with unclosed tags or comments.

Example: Given a site-wide `config.toml` that contains this line:

    copyright = "© 2014 Jane Doe.  <a href=\"http://creativecommons.org/licenses/by/4.0/\">Some rights reserved</a>."

`{{ .Site.Copyright | safeHtml }}` would then output:

> © 2014 Jane Doe.  <a href="http://creativecommons.org/licenses/by/4.0/">Some rights reserved</a>.

However, without the `safeHtml` function, html/template assumes
`.Site.Copyright` to be unsafe, escaping all HTML tags,
rendering the whole string as plain-text like this:

<blockquote>
<p>© 2014 Jane Doe.  &lt;a href=&#34;http://creativecommons.org/licenses/by/4.0/&#34;&gt;Some rights reserved&lt;/a&gt;.</p>
</blockquote>

----

> ### safeCss _(coming soon in v0.13)_
> Declares the provided string as a known "safe" CSS string
> so Go html/templates will not filter it.
> "Safe" means CSS content that matches any of:
>
> 1. The CSS3 stylesheet production, such as `p { color: purple }`.
> 2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
> 3. CSS3 declaration productions, such as `color: red; margin: 2px`.
> 4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.
>
> Example: Given `style = "color: red;"` defined in the front matter of your `.md` file:
>
> * `<p style="{{ .Params.style | safeCss }}">…</p>` ⇒ `<p style="color: red;">…</p>` (Good!)
> * `<p style="{{ .Params.style }}">…</p>` ⇒ `<p style="ZgotmplZ">…</p>` (Bad!)
>
> Note: "ZgotmplZ" is a special value that indicates that unsafe content reached a
> CSS or URL context.

----

> ### safeUrl _(coming soon in v0.13)_
> Declares the provided string as a "safe" URL or URL substring (see [RFC 3986][]).
> A URL like `javascript:checkThatFormNotEditedBeforeLeavingPage()` from a trusted
> source should go in the page, but by default dynamic `javascript:` URLs are
> filtered out since they are a frequently exploited injection vector.
>
> Without `safeUrl`, only the URI schemes `http:`, `https:` and `mailto:`
> are considered safe by Go.  If any other URI schemes, e.g.&nbsp;`irc:` and
> `javascript:`, are detected, the whole URL would be replaced with
> `#ZgotmplZ`.  This is to "defang" any potential attack in the URL,
> rendering it useless.
>
> Example: Given a site-wide `config.toml` that contains this menu entry:
>
>     [[menu.main]]
>         name = "IRC: #golang at freenode"
>         url = "irc://irc.freenode.net/#golang"
>
> The following template:
>
>     <ul class="sidebar-menu">
>       {{ range .Site.Menus.main }}
>       <li><a href="{{ .Url }}">{{ .Name }}</a></li>
>       {{ end }}
>     </ul>
>
> would produce `<li><a href="#ZgotmplZ">IRC: #golang at freenode</a></li>`
> for the `irc://…` URL.
>
> To fix this, add ` | safeUrl` after `.Url` on the 3rd line, like this:
>
>       <li><a href="{{ .Url | safeUrl }}">{{ .Name }}</a></li>
>
> With this change, we finally get `<li><a href="irc://irc.freenode.net/#golang">IRC: #golang at freenode</a></li>`
> as intended.

[RFC 3986]: http://tools.ietf.org/html/rfc3986

----

### lower
Convert all characters in string to lowercase.

e.g. `{{lower "BatMan"}}` → "batman"

### upper
Convert all characters in string to uppercase.

e.g. `{{upper "BatMan"}}` → "BATMAN"

### title
Convert all characters in string to titlecase.

e.g. `{{title "BatMan"}}` → "Batman"

### highlight
Take a string of code and a language, uses Pygments to return the syntax
highlighted code in HTML. Used in the [highlight shortcode](/extras/highlighting/).
