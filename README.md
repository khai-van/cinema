# **Cinema Service**

A service for managing reservations, and cancellations with specific seating distancing requirements. This service uses **gRPC** and **PostgreSQL** as the backend.

---

## **Features**

- **Configure Cinema**: Add a cinema with defined rows and columns.
- **Query Reserved Seats**: Retrieve all reserved seats in a specific cinema.
- **Reserve Seats**: Reserve seats for a specific cinema with distance validation.
- **Cancel Seats**: Cancel specific seat reservations.

---

## **Technology Stack**

- **gRPC**: Communication protocol for efficient service calls.
- **Go**: The primary language for backend development.
- **PostgreSQL**: A relational database for storing cinema and reservation data.

---

## **Requirements**

- **Go**: Version 1.22
- **Docker**: For running the service and database in containers.

---

## **Setup Instructions**

### **1. Clone the Repository**
Clone this repository to your local machine:
```bash
git clone https://github.com/khai-van/cinema.git
cd cinema
```

### **2. Docker Setup**

The project includes a `docker-compose.yml` file to run the service and PostgreSQL in Docker containers.

#### **To run with Docker:**
1. Build and start the containers:
   ```bash
   docker-compose up --build
   ```

2. The **PostgreSQL database** will be available on `localhost:5432`.
3. The **Cinema Service** (gRPC server) will be available on `localhost:8080`.

#### **To stop the containers:**
```bash
docker-compose down
```

---

## **Service Endpoints (gRPC)**

### **1. Configure Cinema**

- **Method**: `ConfigureCinema`
- **Request**:
  ```protobuf
  message CinemaConfig {
    string name = 1;  // Cinema name
    int32 rows = 2;   // Number of rows
    int32 columns = 3; // Number of columns
  }
  ```

- **Response**:
  ```protobuf
  message StatusResponse {
    Status status = 1;
    string message = 2;
  }

  enum Status {
    SUCCESS = 0;
    FAILURE = 1;
  }
  ```

### **2. Query Reserved Seats**

- **Method**: `QueryReservedSeats`
- **Request**:
  ```protobuf
  message QueryReservedSeatsRequest {
    int32 cinema_id = 1; // Cinema ID
  }
  ```

- **Response**:
  ```protobuf
  message QueryReservedSeatsResponse {
    repeated Seat seats = 1;   // List of reserved seats
    int32 max_rows = 2;        // Maximum number of rows
    int32 max_columns = 3;     // Maximum number of columns
    int32 min_distance = 4;    // Minimum distance between seats
  }

  message Seat {
    int32 row = 1;
    int32 column = 2;
  }
  ```

### **3. Reserve Seats**

- **Method**: `ReserveSeats`
- **Request**:
  ```protobuf
  message SeatReservationRequest {
    int32 cinema_id = 1;   // Cinema ID
    repeated Seat seats = 2; // List of seats to reserve
  }
  ```

- **Response**:
  ```protobuf
  message StatusResponse {
    Status status = 1;
    string message = 2;
  }
  ```

### **4. Cancel Seats**

- **Method**: `CancelSeats`
- **Request**:
  ```protobuf
  message SeatCancellationRequest {
    int32 cinema_id = 1;   // Cinema ID
    repeated Seat seats = 2; // List of seats to cancel
  }
  ```

- **Response**:
  ```protobuf
  message StatusResponse {
    Status status = 1;
    string message = 2;
  }
  ```
---

## **PostgreSQL Schema**

Run the following SQL script to create the necessary tables in PostgreSQL:

#### **`setup.sql`**:

```sql
-- Create table for storing cinema details
CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    rows INT NOT NULL,
    columns INT NOT NULL
);

-- Create table for storing reserved seats
CREATE TABLE reserved_seats (
    id SERIAL PRIMARY KEY,
    cinema_id INT REFERENCES cinemas(id) ON DELETE CASCADE,
    row INT NOT NULL,
    col INT NOT NULL,
    status VARCHAR(10) CHECK (status IN ('reserved', 'canceled')),
    UNIQUE (cinema_id, row, col)
);
```

Run this SQL script in your PostgreSQL database to set up the schema.
