# API Rate Limiter with Go Fiber

A demonstration project showing how to implement rate limiting for API endpoints using Go Fiber.

## Features

- Global rate limiting based on IP address
- User-specific rate limiting based on user ID
- Admin-specific rate limiting based on API key
- Configurable rate limits via environment variables
- Different rate limits for different API endpoints
- IP address information endpoint with User-Agent detection

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/1mr-newton/api-rate-limits-golang.git
cd api-rate-limits-golang
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The application can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 3000 |
| GLOBAL_RATE_LIMIT | Global rate limit (requests per minute) | 3 (for testing) |
| USER_RATE_LIMIT | User-specific rate limit (requests per minute) | 3 (for testing) |
| ADMIN_RATE_LIMIT | Admin-specific rate limit (requests per minute) | 3 (for testing) |
| RATE_LIMIT_EXPIRATION | Rate limit expiration time in seconds | 60 |

> **Note:** The default rate limits are set to 3 requests per minute for testing purposes. For production use, you may want to increase these values.

## Running the Application

```bash
go run main.go
```

Or build and run:

```bash
go build
./api-rate-limits-golang
```

## API Endpoints

### Public Endpoint (Global Rate Limit)

```
GET /api/public/
```

This endpoint is subject to the global rate limit.

### IP Information Endpoint

```
GET /api/ip
```

This endpoint returns information about your IP address, including:
- Your IP address
- Any forwarded IPs (if behind a proxy)
- Your hostname
- The rate limit key used for your requests
- Your User-Agent (browser or application making the request)

### User Endpoint (User-specific Rate Limit)

```
GET /api/user/?user_id=123
```

This endpoint is subject to user-specific rate limiting. The rate limit is applied per user ID.
You can also provide the user ID via the `X-User-ID` header.

### Admin Endpoint (Admin-specific Rate Limit)

```
GET /api/admin/
Header: X-API-Key: your-api-key
```

This endpoint is subject to admin-specific rate limiting. The rate limit is applied per API key.

## Testing Rate Limits

You can test the rate limits using tools like curl or Postman:

```bash
# Test public endpoint
curl http://localhost:3000/api/public/

# Get your IP information
curl http://localhost:3000/api/ip

# Test user endpoint
curl http://localhost:3000/api/user/?user_id=123

# Test admin endpoint
curl -H "X-API-Key: your-api-key" http://localhost:3000/api/admin/
```

Try making more than 3 requests within a minute to any endpoint to see the rate limiting in action.

## License

MIT

## Author

[1mr-newton](https://github.com/1mr-newton) 