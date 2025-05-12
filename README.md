# Simple RESTful API in Go (Golang)

This is a basic RESTful API built with Golang using the Gorilla Mux router. It manages a list of blog posts, allowing users to perform full CRUD operations.

## Features

- List all posts (`GET /posts`)
- Get a specific post by ID (`GET /posts/{id}`)
- Add a new post (`POST /posts`)
- Update an existing post (`PUT /posts/{id}`)
- Partially update a post (`PATCH /posts/{id}`)
- Delete a post (`DELETE /posts/{id}`)

## Data Structure

### User

```json
{
  "fullName": "John Doe",
  "userName": "johndoe123",
  "email": "john@example.com"
}
````

### Post

```json
{
  "title": "My First Post",
  "body": "This is the content of the post",
  "author": {
    "fullName": "John Doe",
    "userName": "johndoe123",
    "email": "john@example.com"
  }
}
```

## Getting Started

### Prerequisites

* Go (version 1.16 or above)
* Git

### Installation

1. Clone the repository:

```bash
git clone https://github.com/sharukh010/go-gorilla-mux-CRUD
cd go-gorilla-mux-CRUD
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the application:

```bash
go run main.go
```

Server will start on `http://localhost:8081`.

## API Endpoints

| Method | Endpoint      | Description           |
| ------ | ------------- | --------------------- |
| GET    | `/posts`      | Get all posts         |
| GET    | `/posts/{id}` | Get post by ID        |
| POST   | `/posts`      | Create a new post     |
| PUT    | `/posts/{id}` | Update post by ID     |
| PATCH  | `/posts/{id}` | Partially update post |
| DELETE | `/posts/{id}` | Delete post by ID     |

## Example `POST` Request

```http
POST /posts
Content-Type: application/json

{
  "title": "New Blog Post",
  "body": "This is an awesome blog post.",
  "author": {
    "fullName": "Jane Doe",
    "userName": "janedoe01",
    "email": "jane@example.com"
  }
}
```

## Notes

* This project uses in-memory storage (a slice). All data will be lost once the server stops.
* IDs for posts are based on their index in the slice.

## License

This project is licensed under the MIT License.

