# Receipt Processor / Fetch Assessment

This project is an implementation of the [Fetch Rewards Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge).
It is designed to be **enterprise-ready**, demonstrating software engineering best practices, maintainability, and scalability beyond just fulfilling the API requirements.

## Running the Application

This application is built in **Go** and does not require a database. It stores all data in memory, as specified in the challenge.

### Prerequisites

- Go 1.22 or later
- (Optional) `make` installed for simplified commands

### Quick Start

You can start the application by running:

```sh
go run ./cmd/fetch-assessment
```

Alternatively, using `make`:

```sh
make run
```

The server will start on port **8080** by default.

### Running on a Different Port

To run on a different port, set the `PORT` environment variable:

```sh
make run PORT=9090
```

or

```sh
PORT=9090 go run ./cmd/fetch-assessment
```

### Running Tests

Unit tests and integration tests are included to validate the core functionality.

```sh
make test
```

### Building the Binary

To build the binary:

```sh
make build
```

The compiled binary will be in the `bin/` directory.

## API Specification

The API follows the OpenAPI definition provided in [`api/api.yml`](./api/api.yml).

### **Process Receipts**

- **Path**: `/receipts/process`
- **Method**: `POST`
- **Payload**: Receipt JSON
- **Response**: JSON containing an ID for the receipt

### **Get Points**

- **Path**: `/receipts/{id}/points`
- **Method**: `GET`
- **Response**: JSON containing the number of points awarded

### **Health Check Endpoint**

A standard `/healthz` endpoint is provided for checking service status.

- **Path**: `/healthz`
- **Method**: `GET`
- **Response**: `"OK"` with HTTP `200` status

---

## Architectural Decisions

This project was structured with an **enterprise mindset**, ensuring that future enhancements, scalability, and maintainability are achievable with minimal refactoring.

**Repository Pattern** *(repository.MemoryRepository)*
   - Encapsulates data access, making it easy to swap the in-memory storage for a database-backed solution in the future.
   - Promotes **single responsibility** for handling persistence logic.

**Data Adapter Pattern** *(mapper.ApiReceiptMapper)*
   - Decouples API models (`api.Receipt`) from business models (`rules.Receipt`), ensuring **API evolution does not impact core business logic**.

**Factory Pattern** *(mapper instantiation)*
   - Standardizes how different **mappers are created and managed**.

**Strategy Pattern** *(rules/ validation engine)*
   - Each **rule/validator** is encapsulated as an independent object, allowing easy **modification and extension**.

**Pipeline Pattern / Composite Pattern** *(rules/validation engine)*
   - Each rule/validator is executed sequentially, aggregating results dynamically.
   - Rules/validators are structured **hierarchically**, enabling seamless rule evaluation and composition.

**Extensible Rule/Validator Engine**
   - New rules/validators can be **easily added** with minimal refactoring.

**Separation of Concerns**
   - Validation occurs on api objects, business logic lives in `rules`, data logic in `repository`, and API transformation in `mapper`.

**Minimal Dependencies**
   - Only **essential libraries** are included, ensuring lightweight performance.

**Auto-Generated API Code** *(oapi-codegen)*
   - Ensures **API consistency** while allowing future schema updates.

---

## Engineering Decisions

This project was intentionally **over-engineered** to showcase software engineering maturity, with a focus on:

**Thread-Safe Memory Storage**
   - Ensures **concurrent read access** while preventing race conditions.

**Avoids Floating-Point Math for Money**
   - Prices are stored as **integers (cents)** to eliminate precision errors.
   - Could've used a 3rd party library like **shopspring/decimal** but architectural decision was made to minimize dependencies.

**Unit & Integration Testing**
   - **Critical components** (mappers, rules, repository) have tests to **catch regressions**.
   - An **integration test** ensures that the full API workflow functions correctly.

**Makefile & Tools Isolation**
   - The `Makefile` provides a streamlined **build, test, and run** process.
   - The `tools` directory isolates code generation from core application logic.

**Graceful Shutdown Implementation & Signal Handling**
   - Uses **context.WithTimeout** to ensure all requests complete before shutting down.
   - Ensures **clean termination** of the server with **proper resource cleanup**.

**Override Statuses**
   - Override **500** and return **400** incase of panic/failure. Interpreted spec/requirements as saying there are only 3 valid status codes: **200** + **400** *(POST receipt)* and **200** + **404** *(GET receipt)*
   - Interpreted spec as suggesting to return the description in the response body for a **400** and **404** as *"The receipt is invalid."* and *"No receipt found for that ID."*, respectively, so overrode the generated response objects *(server/status.go)*

---

## Possible Future Enhancements

While this implementation is complete, here are some areas that could be improved in a production setting:

1. **Database Integration**: Replace `MemoryRepository` with PostgreSQL or MongoDB for persistence.
2. **Caching**: Use Redis to store frequently accessed receipts and computed points.
3. **Authentication & Authorization**: Implement JWT-based authentication.
4. **Load Testing**: Ensure the server handles concurrent requests efficiently.
5. **Distributed Deployment**: Add Docker & Kubernetes configuration.

---

## Final Thoughts

This project is **more than a coding exercise**—it was structured as if it were going into production, balancing correctness, maintainability, and extensibility.

Thank you for taking the time to review this submission! I truly appreciate the effort involved in evaluating candidates, and I hope this project demonstrates not just my ability to deliver functional code, but also to think like an **engineer and architect**.

---

For any questions or discussions, feel free to reach out. Looking forward to your feedback!

