# Karteikarten

Basic backend system to CRUD index cards (Karteikarten)

# Usage
Build an executable via `go build -o kk`   
Then run the executable via `./kk`
Then the API is running on port `8080` and can be accessed via `http://localhost:8080/api`

# API Documentation

## Indexcards

### Get all Indexcards
- **Description**: Retrieve a list of all index cards.
- **URL**: `GET /api/indexcards`
- **Response**:
  - Status Code: `200`
  - Content: `[{id: 1, title: "Title", description: "Description"}, ...]`

### Get Indexcard by ID
- **Description**: Retrieve a specific index card by its ID.
- **URL**: `GET /api/indexcard?id={id}`
- **Parameters**:
  - `id` (int) - The unique identifier of the index card. 
- **Response**:
    - Status Code: `200`
    - Content: `{id: 1, question: "Question", answer: "Description"}`
  
### Create Indexcard
- **Description**: Create a new index card.
- **URL**: `POST /api/indexcard`
- **Body**:
  - `question` (string) - The question of the index card.
  - `answer` (string) - The answer of the index card.
- **Example Body**:
  ```json
  {
    "question": "Question",
    "answer": "Answer"
  }
  ```
- **Response**:
    - Status Code: `200`
    - Content: `{id: 1, question: "Question", answer: "Description"}`

### Delete Indexcard
- **Description**: Delete a specific index card by its ID.
- **URL**: `DELETE /api/indexcard?id={id}`
- **Parameters**:
  - `id` (int) - The unique identifier of the index card to be deleted.
- **Response**:
  - Status Code: `200 OK`
  - Content: `{"message": "Index card deleted successfully"}`
  - Status Code: `404 Not Found`
  - Content: `{"error": "Index card not found"}`

### Update Indexcard
- **Description**: Update the answer or question of a specific index card by its ID.
- **URL**: `PATCH /api/indexcard?id={id}`
- **Parameters**:
  - `id` (int) - The unique identifier of the index card to be updated.
- **Body**: JSON object representing the index card with the fields to be updated.
- **Example Body**:
  ```json
  {
    "question": "New question",
    "answer": "New answer"
  }
  ```
- **Response**:
  - Status Code: `200 OK`
  - Content: Updated JSON representation of the index card with the modified fields.
  - Status Code: `400 Bad Request`
  - Content: `{"error": "Invalid request data"}`
  - Status Code: `404 Not Found`
  - Content: `{"error": "Index card not found"}`
