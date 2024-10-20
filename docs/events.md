# Event Management API Documentation

This documentation describes the event management API, covering the key handlers and routes for creating, retrieving, updating, and deleting events. The API supports bulk operations and handles event data in MongoDB.

## Table of Contents
- [Insert Event](#insert-event)
- [Get Event by ID](#get-event-by-id)
- [Get All Events](#get-all-events)
- [Update Event](#update-event)
- [Delete Event by Name](#delete-event-by-name)
- [Bulk Insert Events](#bulk-insert-events)
- [Bulk Update Events](#bulk-update-events)

## Insert Event

### Route:
`POST /events`

### Description:
This handler allows you to create a new event by providing the event details in the request body.

### Request Body:
```json
{
  "name": "Event Name",
  "desc": "Event Description",
  "startDate": "2024-12-01T00:00:00Z",
  "endDate": "2024-12-31T00:00:00Z",
  "github": "https://github.com/event_repo"
}
```

### Structure:
```go
func InsertEvent(c *fiber.Ctx) error
```
- Parses the request body to extract event details.
- Inserts the event into the MongoDB collection.
- Returns a success message with the inserted event ID.

### Response:
```json
{
  "message": "Event inserted successfully",
  "event_id": "ObjectID"
}
```

---

## Get Event by ID

### Route:
`GET /events/:id`

### Description:
This handler retrieves a specific event based on its MongoDB ObjectID.

### Parameters:
- `id`: The ObjectID of the event to retrieve.

### Structure:
```go
func GetEventByID(c *fiber.Ctx) error
```
- Extracts the event ID from the URL parameters.
- Queries the MongoDB collection for the corresponding event.
- Returns the event details if found.

### Response:
```json
{
  "event": {
    "id": "ObjectID",
    "name": "Event Name",
    "desc": "Event Description",
    "startDate": "2024-12-01T00:00:00Z",
    "endDate": "2024-12-31T00:00:00Z",
    "github": "https://github.com/event_repo"
  }
}
```

---

## Get All Events

### Route:
`GET /events`

### Description:
This handler retrieves all the events stored in the MongoDB collection.

### Structure:
```go
func GetEvents(c *fiber.Ctx) error
```
- Fetches all the events from the MongoDB collection.
- Returns the list of all events.

### Response:
```json
{
  "events": [
    {
      "id": "ObjectID",
      "name": "Event 1",
      "desc": "Event 1 Description",
      "startDate": "2024-12-01T00:00:00Z",
      "endDate": "2024-12-31T00:00:00Z",
      "github": "https://github.com/event1_repo"
    },
    {
      "id": "ObjectID",
      "name": "Event 2",
      "desc": "Event 2 Description",
      "startDate": "2024-12-01T00:00:00Z",
      "endDate": "2024-12-31T00:00:00Z",
      "github": "https://github.com/event2_repo"
    }
  ]
}
```

---

## Update Event

### Route:
`PUT /events/:id`

### Description:
This handler updates an existing event by its MongoDB ObjectID. Only the fields provided in the request body will be updated, leaving other fields unaffected.

### Parameters:
- `id`: The ObjectID of the event to update.

### Request Body:
```json
{
  "name": "Updated Event Name",
  "desc": "Updated Event Description",
  "startDate": "2024-12-01T00:00:00Z",
  "endDate": "2024-12-31T00:00:00Z",
  "github": "https://github.com/updated_event_repo"
}
```

### Structure:
```go
func UpdateEvent(c *fiber.Ctx) error
```
- Extracts the event ID from the URL parameters.
- Parses the request body to determine which fields to update.
- Updates the event in the MongoDB collection.
- Returns a success message indicating the event was updated.

### Response:
```json
{
  "message": "Event updated successfully"
}
```

---

## Delete Event by Name

### Route:
`DELETE /events`

### Description:
This handler deletes an event by its name.

### Request Body:
```json
{
  "name": "Event Name"
}
```

### Structure:
```go
func DeleteEvent(c *fiber.Ctx) error
```
- Parses the request body to get the event name.
- Deletes the event with the matching name from the MongoDB collection.
- Returns a success message indicating the number of deleted events.

### Response:
```json
{
  "message": "Event deleted",
  "deleted": 1,
  "name": "Event Name"
}
```

---

## Bulk Insert Events

### Route:
`POST /events/bulk`

### Description:
This handler allows inserting multiple events in bulk.

### Request Body:
```json
{
  "events": [
    {
      "name": "Event 1",
      "desc": "Event 1 Description",
      "startDate": "2024-12-01T00:00:00Z",
      "endDate": "2024-12-31T00:00:00Z",
      "github": "https://github.com/event1_repo"
    },
    {
      "name": "Event 2",
      "desc": "Event 2 Description",
      "startDate": "2024-12-01T00:00:00Z",
      "endDate": "2024-12-31T00:00:00Z",
      "github": "https://github.com/event2_repo"
    }
  ]
}
```

### Structure:
```go
func BulkInsertEvents(c *fiber.Ctx) error
```
- Parses the request body to get the list of events.
- Inserts the events in bulk into the MongoDB collection.
- Returns a success message with the inserted event IDs.

### Response:
```json
{
  "message": "Events inserted successfully",
  "insertedIDs": ["ObjectID1", "ObjectID2"]
}
```

---

## Bulk Update Events

### Route:
`PUT /events/bulk`

### Description:
This handler allows updating multiple events in bulk.

### Request Body:
```json
{
  "events": [
    {
      "name" : "event 1",
      "description" : "event 1 description",
      "sub-description"  : "sub-description",
      "start_date" : "2024-12-01",
      "end_date"  : "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event1.md"
    },
    {
      "name" : "event 2",
      "description" : "event 2 description",
      "sub-description"  : "sub-description",
      "start_date" : "2024-12-01",
      "end_date"  : "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event2.md"
    }
  ]
}
```

### Structure:
```go
func BulkUpdateEvents(c *fiber.Ctx) error
```
- Parses the request body to get the list of events.
- Updates each event in the MongoDB collection by its `ObjectID`.
- Returns a success message indicating the events were updated.

### Response:
```json
{
  "message": "Events updated successfully"
}
```
