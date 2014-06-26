Track all the beer!

Input via admin panel
Get updates when certain beers are low!

## API  
- GET /beer  

- POST /beer
  - name
  - desc
  - type
  - quantity
  - upc

- PUT /beer/:upc
  - name
  - desc
  - type
  - quantity
  - upc

- DELETE /beer/:upc

- POST /checkout
  - upc
  - timestamp
  - loc


## MYSQL

create database beerventory;
create table beer (upc int not null primary key, name varchar(100) not null, type varchar(100) not null, qty int);
create table location (id int not null auto_increment primary key, name varchar(100));
create table events (id int not null auto_increment primary key, created timestamp default NOW(), upc int not null, location_id int, type varchar(100));
