# kv

A shared public ephemeral kv store ;)

*what could go wrong?*


#### info
key value pairs will expire after about 6 hours of inactivity. GETting a value will reset it's keys expirey

key value pairs cannot be overwritten untill they expire.

keys and values will be trunicated
- keys have a max size of 256 characters
- values have a max size of 512 characters


#### routes
```
GET /kv/:key             // fetch a value
POST /kv/:key/:value     // set a value

GET /r/:key              // redirect to value
```




