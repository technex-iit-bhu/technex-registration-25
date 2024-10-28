# Workshop Management API Documentation

This documentation describes the workshop management API, covering the key handlers and routes for creating,
retrieving, updating, and deleting workshops. The API supports bulk operations and handles workshop data in MongoDB.

## Table of Contents
- [Insert Workshop](#insert-workshop)
- [Get Workshop by Name](#get-workshop-by-name)
- [Get Workshop by ID](#get-workshop-by-id)
- [Get All Workshops](#get-all-workshops)
- [Update Workshop](#update-workshop)
- [Delete Workshop by Name](#delete-workshop-by-name)
- [Bulk Insert Workshops](#bulk-insert-workshops)

---

## Insert Workshop

### Route:
`POST /workshops/insertWorkshop`

### Description:
This handler allows you to create a new workshop by providing the workshop details in the request body.

### Request Body:
```json
{
  "title": "Workshop Title",
  "description": "Workshop Description",
  "sub-description": "Additional Information",
  "start_date": "2024-12-01",
  "end_date": "2024-12-31",
  "github": "https://github.com/technex-iit-bhu/workshops/workshop1.md"
}
```

### Structure:
```go
func InsertWorkshop(c *fiber.Ctx) error
```
- Parses the request body to extract workshop details.
- Inserts the workshop into the MongoDB collection.
- Returns a success message with the inserted workshop ID.

### Response:
```json
{
  "message": "Workshop inserted successfully",
  "workshop_id": "ObjectID"
}
```

---

## Get Workshop by Name

### Route:
`GET /workshops/:name`

### Description:
This handler retrieves a specific workshop based on its name.

### Parameters:
- `name`: The title of the workshop to retrieve.

### Structure:
```go
func GetWorkshopDetails(c *fiber.Ctx) error
```
- Extracts the workshop name from the URL parameters.
- Queries the MongoDB collection for the corresponding workshop.
- Returns the workshop details if found.

### Response:
```json
{
  "workshop": {
    "title": "Workshop Title",
    "description": "Workshop Description",
    "sub-description": "Additional Information",
    "start_date": "2024-12-01",
    "end_date": "2024-12-31",
    "github": "https://github.com/technex-iit-bhu/workshops/workshop1.md"
  }
}
```

---

## Get Workshop by ID

### Route:
`GET /workshops/:id`

### Description:
This handler retrieves a specific workshop based on its MongoDB ObjectID.

### Parameters:
- `id`: The ObjectID of the workshop to retrieve.

### Structure:
```go
func GetWorkshopByID(c *fiber.Ctx) error
```
- Extracts the workshop ID from the URL parameters.
- Queries the MongoDB collection for the corresponding workshop.
- Returns the workshop details if found.

### Response:
```json
{
  "workshop": {
    "id": "ObjectID",
    "title": "Workshop Title",
    "description": "Workshop Description",
    "sub-description": "Additional Information",
    "start_date": "2024-12-01",
    "end_date": "2024-12-31",
    "github": "https://github.com/technex-iit-bhu/workshops/workshop1.md"
  }
}
```

---

## Get All Workshops

### Route:
`GET /workshops`

### Description:
This handler retrieves all the workshops stored in the MongoDB collection.

### Structure:
```go
func GetAllWorkshops(c *fiber.Ctx) error
```
- Fetches all the workshops from the MongoDB collection.
- Returns the list of all workshops.

### Response:
```json
{
  "workshops": [
    {
      "id": "ObjectID",
      "title": "Workshop 1",
      "description": "Workshop 1 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/workshops/workshop1.md"
    },
    {
      "id": "ObjectID",
      "title": "Workshop 2",
      "description": "Workshop 2 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/workshops/workshop2.md"
    }
  ]
}
```

---

## Update Workshop

### Route:
`PATCH /workshops/updateWorkshop`

### Description:
This handler updates an existing workshop by its MongoDB ObjectID. Only the fields provided in the request body will be updated, leaving other fields unaffected.

### Parameters:
- `id`: The ObjectID of the workshop to update (passed as a parameter in the URL).

### Request Body:
```json
{
  "title": "Updated Workshop Title",
  "description": "Updated Workshop Description",
  "sub-description": "Updated Information",
  "start_date": "2024-12-01",
  "end_date": "2024-12-31",
  "github": "https://github.com/technex-iit-bhu/workshops/workshop2.md"
}
```

### Structure:
```go
func UpdateWorkshop(c *fiber.Ctx) error
```
- Extracts the workshop ID from the URL parameters.
- Parses the request body to determine which fields to update.
- Updates the workshop in the MongoDB collection.
- Returns a success message indicating the workshop was updated.

### Response:
```json
{
  "message": "Workshop updated successfully"
}
```

---

## Delete Workshop by Name

### Route:
`DELETE /workshops/deleteWorkshop`

### Description:
This handler deletes an workshop by its name.

### Request Body:
```json
{
  "name": "Workshop Title"
}
```

### Structure:
```go
func DeleteWorkshop(c *fiber.Ctx) error
```
- Parses the request body to get the workshop name.
- Deletes the workshop with the matching name from the MongoDB collection.
- Returns a success message indicating the number of deleted workshops.

### Response:
```json
{
  "message": "Workshop deleted",
  "deleted": 1,
  "name": "Workshop Title"
}
```

---

## Bulk Insert Workshops

### Route:
`POST /workshops/insertWorkshops`

### Description:
This handler allows inserting multiple workshops in bulk.

### Request Body:
```json
{
  "workshops": [
    {
      "title": "Workshop 1",
      "description": "Workshop 1 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/workshops/workshop1.md"
    },
    {
      "title": "Workshop 2",
      "description": "Workshop 2 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/workshops/workshop2.md"
    }
  ]
}
```

### Structure:
```go
func BulkInsertWorkshops(c *fiber.Ctx) error
```
- Parses the request body to get the list of workshops.
- Inserts the workshops in bulk into the MongoDB collection.
- Returns a success message with the inserted workshop IDs.

### Response:
```json
{
  "message": "Workshops inserted successfully",
  "insertedIDs": ["ObjectID1", "ObjectID2"]
}
```

To support the new functionality of fetching multiple workshops by their IDs, you can add a new route in your router. Since the request requires passing multiple IDs, a `POST` route would be ideal because the list of IDs will be sent in the request body.

Here is the suggested route:

### New Route for Fetching Multiple Workshops by IDs:
Add the following line in your `Route` function to define a new route for the `GetWorkshopsByID` handler:

```go
workshops.Post("/getWorkshopsByIds", workshop_handler.GetWorkshopsByID)
```

### Updated Router:
```go
workshops := api.Group("/workshops")
workshops.Get("/", workshop_handler.GetAllWorkshops)
workshops.Get("/:name", workshop_handler.GetWorkshopDetails)
workshops.Get("/:id", workshop_handler.GetWorkshopByID)
workshops.Post("/insertWorkshop", workshop_handler.InsertWorkshop)
workshops.Post("/insertWorkshops", workshop_handler.BulkInsertWorkshops)
workshops.Post("/getWorkshopsByIds", workshop_handler.GetWorkshopsByID) // New route for fetching multiple workshops by IDs
workshops.Delete("/deleteWorkshop", workshop_handler.DeleteWorkshop)
workshops.Patch("/updateWorkshop", workshop_handler.UpdateWorkshop)
```

### Documentation to be Appended:

---

## Get Workshops by IDs

### Route:
`POST /workshops/getWorkshopsByIds`

### Description:
This handler retrieves multiple workshops based on their MongoDB ObjectIDs provided in the request body.

### Request Body:
```json
{
  "ids": ["ObjectID1", "ObjectID2", "ObjectID3"]
}
```

### Structure:
```go
func GetWorkshopsByID(c *fiber.Ctx) error
```
- Parses the request body to get the list of workshop IDs.
- Converts the IDs from string format to MongoDB ObjectIDs.
- Queries the MongoDB collection for the workshops matching the provided IDs.
- Returns the list of workshops found.
- Body Format of Ids is as follow : `{"ids": ["id1", "id2", "id3"]}`

### Response:
```json
{
  "workshops": [
    {
      "id": "ObjectID1",
      "title": "Workshop 1",
      "description": "Workshop 1 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/workshops/workshop1.md"
    },
    {
      "id": "ObjectID2",
      "title": "Workshop 2",
      "description": "Workshop 2 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/workshops/workshop2.md"
    }
  ]
}
``` 
