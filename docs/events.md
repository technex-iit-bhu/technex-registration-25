# Event Management API Documentation

This documentation describes the event management API, covering the key handlers and routes for creating,
retrieving, updating, and deleting events. The API supports bulk operations and handles event data in MongoDB.

## Table of Contents
- [Insert Event](#insert-event)
- [Get Event by Name](#get-event-by-name)
- [Get Event by ID](#get-event-by-id)
- [Get All Events](#get-all-events)
- [Update Event](#update-event)
- [Delete Event by Name](#delete-event-by-name)
- [Bulk Insert Events](#bulk-insert-events)

---

## Insert Event

### Route:
`POST /events/insertEvent`

### Description:
This handler allows you to create a new event by providing the event details in the request body.

### Request Body:
```json
{
  "title": "Event Title",
  "description": "Event Description",
  "sub-description": "Additional Information",
  "start_date": "2024-12-01",
  "end_date": "2024-12-31",
  "github": "https://github.com/technex-iit-bhu/events/event1.md"
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

## Get Event by Name

### Route:
`GET /events/:name`

### Description:
This handler retrieves a specific event based on its name.

### Parameters:
- `name`: The title of the event to retrieve.

### Structure:
```go
func GetEventDetails(c *fiber.Ctx) error
```
- Extracts the event name from the URL parameters.
- Queries the MongoDB collection for the corresponding event.
- Returns the event details if found.

### Response:
```json
{
  "event": {
    "title": "Event Title",
    "description": "Event Description",
    "sub-description": "Additional Information",
    "start_date": "2024-12-01",
    "end_date": "2024-12-31",
    "github": "https://github.com/technex-iit-bhu/events/event1.md"
  }
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
    "title": "Event Title",
    "description": "Event Description",
    "sub-description": "Additional Information",
    "start_date": "2024-12-01",
    "end_date": "2024-12-31",
    "github": "https://github.com/technex-iit-bhu/events/event1.md"
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
func GetAllEvents(c *fiber.Ctx) error
```
- Fetches all the events from the MongoDB collection.
- Returns the list of all events.

### Response:
```json
{
  "events": [
    {
      "id": "ObjectID",
      "title": "Event 1",
      "description": "Event 1 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event1.md"
    },
    {
      "id": "ObjectID",
      "title": "Event 2",
      "description": "Event 2 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event2.md"
    }
  ]
}
```

---

## Update Event

### Route:
`PATCH /events/updateEvent`

### Description:
This handler updates an existing event by its MongoDB ObjectID. Only the fields provided in the request body will be updated, leaving other fields unaffected.

### Parameters:
- `id`: The ObjectID of the event to update (passed as a parameter in the URL).

### Request Body:
```json
{
  "title": "Updated Event Title",
  "description": "Updated Event Description",
  "sub-description": "Updated Information",
  "start_date": "2024-12-01",
  "end_date": "2024-12-31",
  "github": "https://github.com/technex-iit-bhu/events/event2.md"
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
`DELETE /events/deleteEvent`

### Description:
This handler deletes an event by its name.

### Request Body:
```json
{
  "name": "Event Title"
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
  "name": "Event Title"
}
```

---

## Bulk Insert Events

### Route:
`POST /events/insertEvents`

### Description:
This handler allows inserting multiple events in bulk.

### Request Body:
```json
{
  "events": [
    {
      "title": "Event 1",
      "description": "Event 1 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event1.md"
    },
    {
      "title": "Event 2",
      "description": "Event 2 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event2.md"
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
