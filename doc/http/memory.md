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