
## Initial Users Endpoints

Password will be salted and hashed which is fine for this
Still don't put any actual passwords in it

/users/:id will allow /users/current which is actually preffered when possible

rest_to_user_relationship
{
    "user"
    "rest"
    "permission" 
}

### POST /register
**Request**
{
    "email": 
    "accountType": "customer | manager",
    "firstName": 
    "lastName": 
    "password": 
}

**Responses**
*200*
no response body

kick back a JWT to be used in other authorised calls, FE will send this in the `Authorization` header

*400*
{
    "error": "msg to display to the user, eg email already in use"
}

*409*
on duplicate error

### POST /login
**Request**
{
    "email": 
    "password":
}

**Responses**
*200*
no response body

kick back a JWT to be used in other authorised calls, FE will send this in the `Authorization` header

*400*
{
    "error": "msg to display to the user, eg password incorrect"
}

### GET /users/current 
**Request**
no request body

**Response**
{
    "id": 
    "accountType": 
    "email": 
    "firstName": 
    "lastName": 
}

### PUT /users/current
**Request**
{
    "email": 
    "firstName": 
    "lastName": 
}

**Response**
*200*
{
    "email": 
    "firstName": 
    "lastName": 
}

*400*
{
    "error": 
}

**Future**
### POST /restaurants

