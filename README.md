# crm-backend
Final project for Udacity Go Nanodegree

## Installation and running

First install Chi

`go get -u github.com/go-chi/chi/v5`

Run

`go run`

Now the API is ready on your localhost

## REST API

### Getting a single customer through a /customers/{id} path

- Endpoint

GET `localhost:3000/customers/1`

- Sample response

```
{
    "id": 1,
    "name": "Logan Roy",
    "role": "CEO",
    "email": "logan@waystar.com",
    "phone": "212-111-1111"
}
```

### Getting all customers through a the /customers path

- Endpoint

GET localhost:3000/customers/1

- Sample response

```
{
    "1": {
        "id": 1,
        "name": "Logan Roy",
        "role": "CEO",
        "email": "logan@waystar.com",
        "phone": "212-111-1111"
    },
    "2": {
        "id": 2,
        "name": "Tom Wambsgans",
        "role": "Head of Media",
        "email": "twambsgans@waystar.com",
        "phone": "212-222-2222",
        "contacted": true
    },
    "3": {
        "id": 3,
        "name": "Nan Pierce",
        "role": "CEO",
        "email": "nan@pgm.com",
        "phone": "212-333-3333"
    }
}
```

### Creating a customer through a /customers path

- Endpoint

POST `localhost:3000/customers`

- Request body

```
{
    "id": 4,
    "name": "Sioban Roy",
    "role": "Nepo Baby",
    "email": "shiv@waystar.com",
    "phone": "212-444-4444"
}
```

- Sample response

```
{
    "id": 4,
    "name": "Sioban Roy",
    "role": "Nepo Baby",
    "email": "shiv@waystar.com",
    "phone": "212-444-4444"
}
```

### Updating a customer through a /customers/{id} path

- Endpoint

PUT `localhost:3000/customers/1`

- Request body

```
{
    "id": 1,
    "name": "Logan Roy",
    "role": "Former CEO",
    "email": "logan@waystar.com",
    "phone": "212-111-1111"
}
```

- Sample response

```
{
    "id": 1,
    "name": "Logan Roy",
    "role": "Former CEO",
    "email": "logan@waystar.com",
    "phone": "212-111-1111"
}
```

### Deleting a customer through a /customers/{id} path

- Endpoint

DELETE `localhost:3000/customers/1`

- Sample response

```
{
    "id": 1,
    "name": "Logan Roy",
    "role": "Former CEO",
    "email": "logan@waystar.com",
    "phone": "212-111-1111"
}
```