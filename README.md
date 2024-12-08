# dating-apps-be

# **Project Details and Instructions**

## **Service Structure**

The project follows a **microservices architecture**. Below is an overview of the structure and responsibilities of each service:

### **Services**

1. **`api-gateway`**

   - Entry point for all API requests.
   - Routes requests to the appropriate backend services (e.g., `profiles-service`, `users-service`).
   - May handle authentication, rate-limiting, and request validation.

2. **`users-service`**

   - Manages user accounts, including registration, login, and profile data.
   - Handles user authentication and authorization.

3. **`profiles-service`**

   - Manages user profiles, including age, bio, photos, and preferences.
   - Provides APIs to update or retrieve profile information.

4. **`date-service`**

   - Handles swipes, matches, and other dating-related actions.
   - Contains logic for recording swipes and checking mutual matches.

5. **`logs-service`**

   - Records activity logs for auditing and tracking user actions.
   - Supports fetching logs for a specific user or event type.

6. **`payment-service`**
   - Manages payments for premium features.
   - Handles transactions, payment methods, and payment statuses.

### **Additional Components**

- **Docker Compose (`compose.yml`)**

  - Orchestrates all microservices, including their dependencies (e.g., databases, ports, networks).
  - Defines service configurations, ports, and environment variables.

- **Makefile**

  - Automates project tasks like building, testing, or running services.

- **Database**
  - Each service interacts with a shared database.
  - Tables include `users`, `profiles`, `swipes`, `matches`, `logs`, and `payments`.

---

## **Instructions to Run the Project**

### **1. Prerequisites**

- **Install Required Tools**:

  - Docker and Docker Compose.
  - Go (for local development of individual services).
  - Protobuf compiler (`protoc`) for generating gRPC code.

- **Clone the Repository**:
  ```bash
  git clone <repository-url>
  cd dating-apps-be
  ```

---

### **2. Running the Project**

#### **Using Docker Compose (Recommended for Full Setup)**

1. **Start All Services**:

   ```bash
   make compose
   ```

   - Builds and starts all services defined in `compose.yml`.
   - Services will start on their defined ports (refer to `compose.yml`).

2. **Stop All Services**:
   ```bash
   ctrl + C
   ```

#### **Running a Single Service Locally (For Development)**

1. Navigate to the specific service directory:

   ```bash
   cd users-service
   ```

2. Build and run the service:

   ```bash
   go build -o users-service .
   ./users-service
   ```

3. Use `curl`, Postman, or gRPC clients to test the service.

---

### **3. Generating Protobuf Code**

If you make changes to `.proto` files, regenerate the gRPC code:

```bash
make protoc
```

---

### **4. Environment Variables**

Check `.env` files in each service or the `compose.yml` file for required environment variables, such as:

- **Database connection strings**.
- **Service-specific configurations** (e.g., API keys, secret tokens).

---

#### **Integration Tests**

Use Postman, gRPC clients, or custom scripts to test API requests.

---

### **6. Common Ports**

| **Service**      | **Port** |
| ---------------- | -------- |
| API Gateway      | `8080`   |
| Users Service    | `50001`  |
| Profiles Service | `50002`  |
| Date Service     | `50003`  |
| Logs Service     | `50004`  |
| Payment Service  | `50005`  |

### DEPLOYMENT TO GCP - CLOUD RUN

**DEPLOYED GATEWAY URL** <br>
API-Gateway: [https://api-gateway-611320088750.asia-southeast2.run.app] <br>

**_SERVICES API URL_** <br>
Users: [https://users-service-611320088750.asia-southeast2.run.app] <br>
Profiles: [https://profiles-service-611320088750.asia-southeast2.run.app] <br>
Payments: [https://payment-service-611320088750.asia-southeast2.run.app]<br>
Logs: [https://logs-service-611320088750.asia-southeast2.run.app]<br>
Date: [https://date-service-611320088750.asia-southeast2.run.app]<br>
