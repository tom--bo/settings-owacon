how to use alp
======

about alp, read `alp --help`  (github)[https://github.com/tkuchiki/alp]


##normal use
you can sort by each params or in reverse
```
cat access.log |alp --sum -r 
```


##aggregate by url
```
cat access.log| alp --aggregates=/memo/.*,/recent/.* --sum -r 
```
