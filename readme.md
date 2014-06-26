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
