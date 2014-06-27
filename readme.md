Track all the beer!

Input via admin panel
Get updates when certain beers are low!

## API  
- GET /beer  
```{"beers":[{"upc":"", "name":"", "desc":"", "type":"", "quantity":""}]}```

- GET /beer/:upc

```{"upc":"", "name":"", "desc":"", "type":"", "quantity":""}```

- POST /beer
  - name
  - desc
  - type
  - quantity
  - upc

```{"upc":"", "name":"", "desc":"", "type":"", "quantity":""}```

- PUT /beer/:upc
  - name
  - desc
  - type
  - quantity
  - upc

```{"upc":"", "name":"", "desc":"", "type":"", "quantity":""}```

- DELETE /beer/:upc

```{"deleted":true}```

- POST /checkout
  - upc
  - timestamp
  - loc

```{"upc":"", "timestamp":"", "loc":""}```


## MYSQL

mysql -u <username> -p beerventory < seed_data.sql
