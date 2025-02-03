# devops1.hng
This project provides an API to classify numbers and return useful properties about them. It checks if a number is prime, perfect, Armstrong, and also calculates the sum of its digits. Additionally, it fetches a fun mathematical fact for the number.

## Endpoints
`GET /api/classify-number`
- Query Parameters: number (integer)
- Response: JSON object with the number's properties (prime, perfect, armstrong, digit sum, and fun fact).
**Success Response:**
```json
{
  "number": 153,
  "is_prime": false,
  "is_perfect": false,
  "properties": ["armstrong", "odd"],
  "digit_sum": 9,
  "fun_fact": "153 is a narcissistic number."
}
```
**Error Response**
```json
{
  "number": "abc",
  "error": true
}
```

## Setup
1. Clone the repository
2. Create a '.env' file at the root directory with the following content:
```
PORT=8080
```
3. Install dependencies
```bash
go mod tidy
```
4. Run the project
```bash
go run *.go
```
