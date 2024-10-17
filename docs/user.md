## API Documentation

### Authentication
Most endpoints require authentication via a Bearer token in the Authorization header:
```
Authorization: Bearer <token>
```

### Endpoints

#### Health Check
```http
GET /api/
```
Check if the API is running.

**Response**
```json
{
  "message": "Api is running",
  "data": null
}
```

#### User Registration
```http
POST /api/user/register
```
Create a new user account.

**Request Body**
```json
{
  "username": "string",
  "password": "string",
  "email": "string",
  "name": "string",
  "institute": "string",
  "city": "string",
  "year": "number",
  "branch": "string",
  "phone": "string"
}
```

**Response**
```json
{
  "id": "string" // MongoDB ObjectID
}
```

#### Get User Profile
```http
GET /api/user/profile
```
Get the current user's profile information.

**Headers**
```
Authorization: Bearer <token>
```

**Response**
```json
{
  "data": {
    "username": "string",
    "email": "string",
    "name": "string",
    "institute": "string",
    "city": "string",
    "year": "number",
    "branch": "string",
    "phone": "string",
    "createdAt": "timestamp",
    "updatedAt": "timestamp"
  }
}
```

#### Password Login
```http
POST /api/user/login/password
```
Authenticate user with username and password.

**Request Body**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response**
```json
{
  "token": "string"
}
```

#### Google Login
```http
POST /api/user/login/google
```
Authenticate user with Google token.

**Request Body**
```json
{
  "google_token": "string"
}
```

**Response**
```json
{
  "token": "string"
}
```

#### GitHub Login
```http
POST /api/user/login/github
```
Authenticate user with GitHub token.

**Request Body**
```json
{
  "github_token": "string"
}
```

**Response**
```json
{
  "token": "string"
}
```

#### Delete Account
```http
DELETE /api/user/delete
```
Delete the current user's account.

**Headers**
```
Authorization: Bearer <token>
```

**Response**
```json
{
  "message": "user yeeted successfully"
}
```

#### Update Profile
```http
PATCH /api/user/update
```
Update current user's profile information.

**Headers**
```
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "name": "string",
  "old_password": "string",
  "new_password": "string",  // Optional
  "institute": "string",
  "city": "string",
  "year": "number",
  "branch": "string",
  "phone": "string"
}
```

**Response**
```json
{
  "message": "user updated successfully"
}
```

#### Password Recovery Request
```http
GET /api/user/recovery/:username
```
Send password recovery email to user.

**Parameters**
- username: User's username

**Response**
```json
{
  "message": "recovery email sent successfully"
}
```

#### Update Password with Recovery Token
```http
POST /api/user/verify_recovery_and_update_password
```
Update password using recovery token.

**Request Body**
```json
{
  "recovery_token": "string",
  "new_password": "string"
}
```

**Response**
```json
{
  "message": "password updated successfully"
}
```

### Status Codes
- 200: Success
- 201: Created
- 400: Bad Request
- 404: Not Found
- 500: Server Error

### Error Responses
All endpoints may return error responses in the following format:
```json
{
  "message": "string" // Error description
}
```