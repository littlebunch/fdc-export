# fdc-export
Exports Food Data Central documents in a Couchbase bucket to a file as JSON.
## Building
go version 11 or greater is required.    
Clone this repo into any location *other* than your $GOPATH:   
```
git clone git@github.com:littlebunch/fdc-export.git
```
and cd to the repo root, e.g.:
```
cd ~/fdc-export
```
Initialize the module:     
```
go mod init github.com/littlebunch/fdc-export
```
Build the module:    
```
go get 
```
You should now be ready to build or run an executable.   
## Running
```
go run exporter.go
```
Options:   
Required:     
```-o <outfile.json>```
DOCTYPE can be one of FOOD, NUTDATA, or NUT:   
```-t <DOCTYPE`>``
Optional:
```-s <offset> //Document offset to begin scan, default is 0
```
```
-n <export number> // Maximum # of docs to export, if not specified all docs will be written
```