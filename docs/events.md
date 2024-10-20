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

To support the new functionality of fetching multiple events by their IDs, you can add a new route in your router. Since the request requires passing multiple IDs, a `POST` route would be ideal because the list of IDs will be sent in the request body.

Here is the suggested route:

### New Route for Fetching Multiple Events by IDs:
Add the following line in your `Route` function to define a new route for the `GetEventsByID` handler:

```go
events.Post("/getEventsByIds", event_handler.GetEventsByID)
```

### Updated Router:
```go
events := api.Group("/events")
events.Get("/", event_handler.GetAllEvents)
events.Get("/:name", event_handler.GetEventDetails)
events.Get("/:id", event_handler.GetEventByID)
events.Post("/insertEvent", event_handler.InsertEvent)
events.Post("/insertEvents", event_handler.BulkInsertEvents)
events.Post("/getEventsByIds", event_handler.GetEventsByID) // New route for fetching multiple events by IDs
events.Delete("/deleteEvent", event_handler.DeleteEvent)
events.Patch("/updateEvent", event_handler.UpdateEvent)
```

### Documentation to be Appended:

---

## Get Events by IDs

### Route:
`POST /events/getEventsByIds`

### Description:
This handler retrieves multiple events based on their MongoDB ObjectIDs provided in the request body.

### Request Body:
```json
{
  "ids": ["ObjectID1", "ObjectID2", "ObjectID3"]
}
```

### Structure:
```go
func GetEventsByID(c *fiber.Ctx) error
```
- Parses the request body to get the list of event IDs.
- Converts the IDs from string format to MongoDB ObjectIDs.
- Queries the MongoDB collection for the events matching the provided IDs.
- Returns the list of events found.
- Body Format of Ids is as follow : `{"ids": ["id1", "id2", "id3"]}`

### Response:
```json
{
  "events": [
    {
      "id": "ObjectID1",
      "title": "Event 1",
      "description": "Event 1 Description",
      "sub-description": "Additional Information",
      "start_date": "2024-12-01",
      "end_date": "2024-12-31",
      "github": "https://github.com/technex-iit-bhu/events/event1.md"
    },
    {
      "id": "ObjectID2",
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
