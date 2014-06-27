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

create database beerventory;
create table beer (upc int not null primary key, name varchar(100) not null, type varchar(100) not null, qty int);
create table location (id int not null auto_increment primary key, name varchar(100));
create table events (id int not null auto_increment primary key, created timestamp default NOW(), upc int not null, location_id int, type varchar(100));
