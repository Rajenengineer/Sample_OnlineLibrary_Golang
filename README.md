# Online Library

## Description :
This Application work as online library system which stores the file in system in particular format.
* For a file storage it takes four arguments 
1.fileName
2.author
3.region
4.category
and then this file is stored as filename_author_category_region.txt. only txt format is supported yet.
Also give options to download a file from system with full details , not on regex basis, and all data is sent over http, no external link is given.

user can also search for files meta data with given name, and if there are multiple filename with same name then all meta data is sent, it like searching a file name and then taking a decision basis on the author/category/region

user can delete the file from system via putting complete details of file

*Right Now there is no concept of super user and role, like only admin or author can delete the file, this is simple online library system wihout scope


## Technologies :
golang 1.15, file system pattern, http interfaces(inbuilt), go routine , go mod concept is being used

## Deployment :
For deplyment just need to go to main directory ( online-library), then run the command "go run main.go" , it will install all the necessary tool and utility that are in use.

## Api
* for upload : 
  api : "/api/v1/file" 
  method : post ( send data in multipart form)
  fields: file [file.txt], author[author], category[category], region[region]
  resp : 200OK

 * for download :
  api : "/api/v1/file" 
  method : get ( send data in form)
  fields: fileName [file.txt], author[author], category[category], region[region]
  resp : byte data 

 * for delete
    api : "/api/v1/file" 
    method : delete ( send data in form)
    fields: fileName [file.txt], author[author], category[category], region[region]
    resp :200OK

  * For Query
    api : /api/v1/query
    method : GET
    fields : fileName[fileName]
    resp : []{ fileName: "abc.txt", author: "dummy", category: "comedy", region:"west_SA"}  

## Constraints to follow
* Don't use underscore in between the words for now (_ is a key words) for the api fields { fileName, author, category, region}
use single word

* No trailing spaces are handled, & all data operations are case sensitive

