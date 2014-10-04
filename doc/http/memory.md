# Memory API

## Remember a thing.

```
POST /thing
```

### JSON Input

| Name | Type | Description |
|:--------:|:--------:|:---------------:|
| thing | string | **Required**. The thing you want to remember. |


### Response

```
Status 201 Created
{
  "id": "1234-5678-90abcd"
}
```

## Remember a thing with custom ID.

```
PUT /thing/:id
```


### JSON Input

| Name | Type | Description |
|:--------:|:--------:|:---------------:|
| thing | string | **Required**. The thing you want to remember. |


### Response

```
Status 200 OK
{
  "id": "custom-key"
}
```


## Retrieve a thing by ID.

```
GET /thing/:id
```


### Response

```
Status 200 OK
{
  "id": "1234-5678-90abcd",
  "thing": "foobar"
}
```


## Remember a list of thing.

```
POST /list/thing
```

### JSON Input

| Name | Type | Description |
|:--------:|:--------:|:---------------:|
| things | array | **Required**. The list of things you want to remember. |


### Response

```
Status 201 Created
{
  "id": "1234-5678-90abcd"
}
```


### Remember a list of things with custom ID.

```
PUT /list/thing/:id
```


### JSON Input

| Name | Type | Description |
|:--------:|:--------:|:---------------:|
| things | array | **Required**. The list of things you want to remember. |


### Response

```
Status 200 OK
{
  "id": "custom-key"
}
```

## Push a thing to a list

```
PUT /list/thing/:id/item
```


### JSON Input

| Name | Type | Description |
|:--------:|:--------:|:---------------:|
| thing | string | **Required**. The thing you want to push to the list. |


### Response

```
Status 200 OK
{
  "id": "custom-key"
}
```


## Retrieve a list of things by ID.

```
GET /list/thing/:id
```

### Response

```
Status 200 OK
{
  "id": "1234-5678-90abcd",
  "things": ["foo", "bar"]
}
```
