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

**Response Codes**

- `200` - API is running successfully
- `500` - Server error

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

**Response Codes**

- `201` - User created successfully
- `400` - Invalid request body
- `500` - Server error

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

**Response Codes**

- `200` - Success
- `404` - Invalid token or user not found
- `500` - Server error

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

**Response Codes**

- `200` - Success
- `404` - Invalid credentials
- `500` - Server error

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

**Response Codes**

- `200` - Success
- `404` - Invalid token or email not found
- `500` - Server error

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

**Response Codes**

- `200` - Success
- `404` - Invalid token or GitHub account not found
- `500` - Server error

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

**Response Codes**

- `200` - Success
- `404` - User not found or invalid token
- `500` - Server error

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

**Response Codes**

- `200` - Success
- `404` - User not found or invalid credentials/password
- `500` - Server error

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

**Response Codes**

- `200` - Success, recovery email sent
- `404` - User not found
- `500` - Server error or unable to send email

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

**Response Codes**

- `200` - Success
- `400` - Invalid recovery token
- `404` - User not found
- `500` - Server error

### Error Response Format

All error responses follow this format:

```json
{
  "message": "string" // Error description
}
```
