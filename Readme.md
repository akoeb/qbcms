# another blog engine #

* golang && gin
* with templates
* returns html, rss or json
* everything stored in sqlite db
* integrated in memory cache
* allows upload of templates and static assets
* checks templates for correct usage of all vars for all posts / pages
* checks whether assets are used
* revisioning of templates, posts, 


## entities ##

* post
  * id
  * title
  * description
  * text (only basic html allowed)
  * status
  * author, fk to user

* page
  * id
  * name
  * URL
  * template, fk to template
  * list of posts

* template
  * id
  * hierarchy, unique identificator to find template
  * content
  
* asset
  * id
  * path
  * type
  * binary content

* user
  * id
  * name
  * login
  * passwd

## workflows ##

**1. Authors**

* add new post (automatically creates single post page, optionally puts post on other pages as well)
* delete post
* modify post
* change status of post (publish / un-publish)

* add new page
* delete page (does not delete the posts of page, but maybe it un-publishes)
* modify page
* add post to page
* remove post from page

* upload template
* change template 
* delete template (only if no posts are bound)

* upload asset
* change asset (for css/js/)
* delete asset

* change password
* delete account

(new users over command line for a start)

**automatic**

* on request: check which page for url, verify page is in cache and return
* adding a post also creates at least a single-post page
* modif
