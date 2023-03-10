<div align="center">

![alt text](https://github.com/MeiSastraJayadi/acacia/blob/master/acacia-wordmark.png "Acacia's Logo")

<div id="badges">
  <img src="https://img.shields.io/static/v1?logo=Go&label=Golang&message=Golang%20Package&color=blue" alt="golang"/>
  <img src="https://img.shields.io/github/license/MeiSastraJayadi/acacia" alt="license"/>
</div>

</div>

# About 

#### Acacia is lightweight package that will help you to route http server. Acacia route the path based on tree data structure. This package also support parameters in the path, so it will help you to build http routing with some parameters. in the path

--- 

# Install 

```console
$ go get github.com/MeiSastraJayadi/acacia

```

----

# Router
## Create New Router
```golang
router := multiplexer.NewRouter("/")
```
## Add Prefix to the Router
##### SetPrefix will add a prefix into the router path. the parameter that will be accept in function SetPrefix is just a string 
```golang
router := multiplexer.NewRouter("/")
router.SetPrefix("product")
```
## Add Some Methods 
##### With Router.Methods(), some methods will be added into the router. The methods can be added more than just one method if the router can handle more than one methods
```golang
router := multiplexer.NewRouter("/")
routerPost := router.Methods(http.MethodsPost)
```
##### Or
```golang
router := multiplexer.NewRouter("/")
routerPost := router.Methods(http.MethodsPost, http.MethodGet)
```
## Add Subrouter
##### A router can be added some subrouter from another router. To do that, what you can do is
```golang
router := multiplexer.NewRouter("/")
routerPost := router.Methods(http.MethodsPost, http.MethodGet)

subrouter := multiplexer.NewRouter("/") 
subrouter.SetPrefix("product")
subrouterPost := subrouter.Methods(http.MethodGet)
router.SubRouter(subrouter)
```


