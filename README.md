# fdc-export
Exports Food Data Central documents in a Couchbase bucket to a file as JSON.  (See https://github.com/littlebunch/fdc-api/blob/master/README.md for how to create a FDC datastore.)    
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

## Configuration
Configuration is minimal and can be in a YAML file or envirnoment variables which override the config file.   

```
couchdb:   
  url:  localhost   
  bucket: gnutdata   //default  bucket    
  fts: fd_food  // default full-text index   
  user: <your_user>    
  pwd: <your_password>    

```
      
Environment   
```
COUCHBASE_URL=localhost   
COUCHBASE_BUCKET=gnutdata   
COUCHBASE_FTSINDEX=fd_food   
COUCHBASE_USER=user_name   
COUCHBASE_PWD=user_password   
```
## Running
```
go run exporter.go
```
### Options   
Required:          
```
-t <DOCTYPE`> //DOCTYPE can be one of FOOD, NUTDATA, or NUT
```     
Optional:
```
-c <config.yml> // Configuration file, defaults to config.yml
```
```
-o <outfile.json> // Output file, default is out.json
```
```
-s <offset> //Document offset to begin scan, default is 0
```
```
-n <export number> // Maximum # of docs to export, if not specified all docs will be written
```

