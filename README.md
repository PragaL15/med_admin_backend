## 🏥 Medical Record Backend - Secure API with JWT & Role-Based Authentication 🚀

Welcome to the **Medical Record Backend**, a secure and scalable system designed to **fetch, store, and manage** patient details, doctor details, and doctor-patient conversations. It provides **role-based access control (RBAC) with JWT authentication** to ensure data privacy and security. 🔐  

---

## 🌟 **Key Features**
✅ **JWT Authentication** - Secure login and token-based access control 🔑  
✅ **Role-Based Access (RBAC)** - API access controlled by user roles (Admin, Doctor, Patient) 🛂  
✅ **Secure API Endpoints** - Restricted access to medical records based on role 🛡️  
✅ **Complete CRUD Operations** - Supports `GET`, `POST`, `PUT`, `DELETE`, and `UPDATE` methods 🔄  
✅ **Doctor-Patient Interaction Tracking** - Stores records of conversations securely 📜  
✅ **Scalable & Optimized** - Built with **Go (Golang)** and **Fiber** for high performance ⚡  

---

## 🛠️ **Tech Stack**
- 🚀 **Golang** - High-performance backend  
- ⚡ **Fiber** - Lightweight & fast web framework  
- 🗄️ **PostgreSQL** - Reliable database  
- 🔐 **JWT Authentication** - Secure API access  
- 🛂 **Role-Based Access Control (RBAC)** - Fine-grained permissions  
- 📡 **Docker (Future Scope)** - Containerized deployment  

---

## 🚀 **How to Set Up Locally**
### **🔹 Prerequisites**
1️⃣ Install **Go** (v1.19+) 🛠️  
2️⃣ Install **PostgreSQL** 🗄️  
3️⃣ Clone the repository 🔽  
```sh
git clone https://github.com/PragaL15/medical_record.git
cd medical_record
```
4️⃣ Install dependencies 📦  
```sh
go mod tidy
```
5️⃣ Create a **.env** file for environment variables 🌍  
```sh
PORT=4000
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=medical_db
JWT_SECRET=your_secret_key
```
6️⃣ Run the server 🚀  
```sh
go run main.go
```
Now, the backend is running on **`http://localhost:4000`** 🎉  

---

## 📌 **Project Structure**
```sh
medical_record/
│── handlers/        # API handlers (Patients, Doctors, Conversations)
│── middleware/      # JWT Authentication & Role-Based Access Control
│── routes/          # API Route definitions
│── db/              # Database connection setup
│── models/          # Database models & schemas
│── main.go          # Main entry point of the server
│── go.mod           # Go module dependencies
│── .env             # Environment variables (Port, DB config, JWT secret)
```

---

## 🔑 **Authentication & Role-Based Access**
### **👥 User Roles**
- **Admin** 🏛️ - Full access to all APIs (Manage doctors, patients, and records)  
- **Doctor** 🩺 - Can fetch and update patient details, add medical records  
- **Patient** 🧑‍⚕️ - Can only view their own records  

### **🔒 Secured API Paths**
| **Role**    | **Accessible Endpoints** | **Methods** |
|------------|-------------------------|------------|
| **Admin** 🏛️ | `/users` `/patients` `/doctors` `/records` | `GET, POST, PUT, DELETE` |
| **Doctor** 🩺 | `/patients` `/records` | `GET, POST, UPDATE` |
| **Patient** 🧑‍⚕️ | `/records/{patient_id}` | `GET` |

🚀 **JWT Authentication is required for all API calls**. Every request must include a valid token in the header:  
```http
Authorization: Bearer <your-jwt-token>
```
---

### 🔥 **Upcoming Features**

- **Audit Logs** - Track changes to medical records 📊  
- **Two-Factor Authentication (2FA)** - Extra layer of security 🔐  
- **Email & SMS Notifications** - Appointment reminders 📩  
- **Docker & Kubernetes Deployment** - Scalable containerized setup 🐳  

---

## 👨‍💻 **Contributors**
💡 **Pragalya Kanakaraj** - Backend Developer 🚀  

---

## 📝 **License**
📜 MIT License - Use it freely, modify it responsibly!  
