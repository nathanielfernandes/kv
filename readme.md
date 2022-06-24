# kv

A shared public ephemeral kv store ;)

*what could go wrong?*


#### info
keys and values will be trunicated
- keys have a max size of 256 characters
- values have a max size of 512 characters


#### routes
```
GET /kv/:key        // fetch a value
POST /kv/:key       // set a value

GET /r/:key         // redirect to value
```




