---
aliases:
- /doc/variables/
- /layout/variables/
date: 2013-07-01
linktitle: Variables
menu:
  main:
    parent: layout
next: /templates/content
prev: /templates/functions
title: Template Variables
weight: 20
---

Hugo makes a set of values available to the templates. Go templates are context based. The following
are available in the context for the templates.

## Page Variables

The following is a list of most of the accessible variables which can be
defined for a piece of content. Many of these will be defined in the front
matter, content or derived from file location.

<dl>
<dt><code>.Title</code></dt><dd> The title for the content.</dd>
<dt><code>.Content</code></dt><dd>The content itself, defined below the front matter.</dd>
<dt><code>.Summary</code></dt>
<dd>A generated summary of the content for easily showing a snippet in a summary view. Note that the breakpoint can be set manually by inserting <code>&lt;!&#x2d;&#x2d;more&#x2d;&#x2d;&gt;</code> at the appropriate place in the content page.  See <a href="/content/summaries/">Summaries</a> for more details.</dd>
<dt><code>.Truncated</code></dt>
<dd>A boolean, <code>true</code> if the <code>.Summary</code> is truncated.  Useful for showing a “Read more…” link only if necessary.  See <a href="/content/summaries/">Summaries</a> for more details.</dd>
<dt><code>.Description</code></dt><dd>The description for the content.</dd>
<dt><code>.Keywords</code></dt><dd>The meta keywords for this content.</dd>
<dt><code>.Date</code></dt><dd>The date the content is associated with.</dd>
<dt><code>.PublishDate</code></dt><dd>The date the content is published on.</dd>
<dt><code>.Type</code></dt><dd>The content <a href="/content/types/">type</a> (e.g. post).</dd>
<dt><code>.Section</code></dt><dd>The <a href="/content/sections/">section</a> this content belongs to.</dd>
<dt><code>.Permalink</code></dt><dd>The Permanent link for this page.</dd>
<dt><code>.RelPermalink</code></dt><dd>The Relative permanent link for this page.</dd>
<dt><code>.LinkTitle</code></dt><dd>Access when creating links to this content. Will use <code>linktitle</code> if set in front-matter, else <code>title</code>.</dd>
<dt><code>.Taxonomies</code></dt><dd>These will use the field name of the plural form of the index (see tags and categories below).</dd>
<dt><code>.RSSLink</code></dt><dd>Link to the taxonomies’ RSS link.</dd>
<dt><code>.TableOfContents</code></dt><dd>The rendered table of contents for this content.</dd>
<dt><code>.Prev</code></dt><dd>Pointer to the previous content (based on pub date).</dd>
<dt><code>.Next</code></dt><dd>Pointer to the following content (based on pub date).</dd>
<dt><code>.PrevInSection</code></dt><dd>Pointer to the previous content within the same section (based on pub date)</dd>
<dt><code>.NextInSection</code></dt><dd>Pointer to the following content within the same section (based on pub date)</dd>
<dt><code>.FuzzyWordCount</code></dt><dd>The approximate number of words in the content.</dd>
<dt><code>.WordCount</code></dt><dd>The number of words in the content.</dd>
<dt><code>.ReadingTime</code></dt><dd>The estimated time it takes to read the content in minutes.</dd>
<dt><code>.Weight</code></dt><dd>Assigned weight (in the front matter) to this content, used in sorting.</dd>
<dt><code>.IsNode</code></dt><dd>Always false for pages.</dd>
<dt><code>.IsPage</code></dt><dd>Always true for page.</dd>
<dt><code>.Site</code></dt><dd>See Site Variables below.</dd>
</dl>

## Page Params

Any other value defined in the front matter, including taxonomies, will be made available under `.Params`.
Take for example I'm using *tags* and *categories* as my taxonomies. The following would be how I would access them:

* `.Params.tags`
* `.Params.categories`

**All Params are only accessible using all lowercase characters**

## Node Variables
In Hugo, a node is any page not rendered directly by a content file. This
includes taxonomies, lists and the homepage.

<dl>
<dt><code>.Title</code></dt><dd> The title for the content.</dd>
<dt><code>.Date</code></dt><dd>The date the content is published on.</dd>
<dt><code>.Permalink</code></dt><dd>The Permanent link for this node</dd>
<dt><code>.Url</code></dt><dd>The relative URL for this node.</dd>
<dt><code>.RSSLink</code></dt><dd>Link to the taxonomies’ RSS link</dd>
<dt><code>.Data</code></dt><dd>The data specific to this type of node.</dd>
<dt><code>.IsNode</code></dt><dd>Always true for nodes.</dd>
<dt><code>.IsPage</code></dt><dd>Always false for nodes.</dd>
<dt><code>.Site</code></dt><dd>See Site Variables below</dd>
</dl>

## Site Variables

Also available is `.Site` which has the following:

<dl>
<dt><code>.Site.BaseUrl</code></dt><dd>The base URL for the site as defined in the site configuration file.</dd>
<dt><code>.Site.Taxonomies</code></dt><dd>The <a href="/taxonomies/usage/">taxonomies</a> for the entire site. Replaces the now-obsolete <code>.Site.Indexes</code> since v0.11.</dd>
<dt><code>.Site.LastChange</code></dt><dd>The date of the last change of the most recent content.</dd>
<dt><code>.Site.Recent</code></dt><dd>Array of all content ordered by Date, newest first.  To be replaced by <code>.Site.Pages</code> starting v0.13.</dd>
<dt><code>.Site.Params</code></dt><dd>A container holding the values from the <code>params</code> section of your site configuration file. For example, a TOML config file might look like this:
<pre><code>baseurl = "http://yoursite.example.com/"

[params]
  description = "Tesla's Awesome Hugo Site"
  author = "Nikola Tesla"
</code></pre></dd>
<dt><code>.Site.Sections</code></dt><dd>Top level directories of the site.</dd>
<dt><code>.Site.Pages</code></dt><dd>All of the content pages of the site.</dd>
<dt><code>.Site.Files</code></dt><dd>All of the source files of the site.</dd>
<dt><code>.Site.Menus</code></dt><dd>All of the menus in the site.</dd>
<dt><code>.Site.Title</code></dt><dd>A string representing the title of the site.</dd>
<dt><code>.Site.Author</code></dt><dd>A map of the authors as defined in the site configuration.</dd>
<dt><code>.Site.LanguageCode</code></dt><dd>A string representing the language as defined in the site configuration.</dd>
<dt><code>.Site.DisqusShortname</code></dt><dd>A string representing the shortname of the Disqus shortcode as defined in the site configuration.</dd>
<dt><code>.Site.Copyright</code></dt><dd>A string representing the copyright of your web site as defined in the site configuration.</dd>
<dt><code>.Site.LastChange</code></dt><dd>A string representing the last time content has been updated.</dd>
<dt><code>.Site.Permalinks</code></dt><dd>A string to override the default permalink format. Defined in the site configuration.</dd>
<dt><code>.Site.BuildDrafts</code></dt><dd>A boolean (Default: false) to indicate whether to build drafts. Defined in the site configuration.</dd>
</dl>
