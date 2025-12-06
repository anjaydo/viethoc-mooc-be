# HocViet (HọcViệt) 🇻🇳

**"Learn the Job. Master the Craft."**

HocViet is a scalable, API-first MOOC platform built for the Vietnamese market. The name is a play on words: "Học Việc" (Apprenticeship/Learning on the job) and "Học Việt" (Learning for Vietnam).

Unlike generic video platforms, HocViet focuses on deep knowledge sharing where learners connect directly with industry experts through structured courses, manual assessments, and mentorship.

## 🏗 System Architecture

HocViet allows multiple Academies and Universities to host their own isolated learning environments (Multi-tenancy) while sharing a robust core infrastructure.

```mermaid
graph TD
    %% Styling
    classDef user fill:#f9f,stroke:#333,stroke-width:2px;
    classDef go fill:#add8e6,stroke:#333,stroke-width:2px;
    classDef db fill:#ff9,stroke:#333,stroke-width:2px;
    classDef ext fill:#f96,stroke:#333,stroke-width:2px;

    %% Users
    Admin[👩‍🏫 Expert / Mentor]
    Student[👨‍🎓 Learner (Apprentice)]

    %% The Application (Modular Monolith)
    subgraph "HocViet Core (Golang)"
        Router[⚡ Gin Router]
        
        subgraph "Modules"
            AuthMod[Auth & Tenants]
            ContentMod[Course Management]
            GradingMod[Assessment & Feedback]
        end
        
        PaymentLogic[💰 VietQR Manual Handler]
    end

    %% Frontend
    subgraph "Presentation Layer"
        NextCMS[CMS (Mentor Dashboard)]
        NextLearn[Learning App]
    end

    %% External
    Supabase[(Supabase Postgres + RLS)]
    Mux[Mux Video API]

    %% Connections
    Admin -->|Manage Content| NextCMS
    Student -->|Learn & Submit| NextLearn
    
    NextCMS -->|JSON API| Router
    NextLearn -->|JSON API| Router

    Router --> AuthMod
    Router --> ContentMod
    Router --> GradingMod
    Router --> PaymentLogic

    %% Logic
    ContentMod -->|Uploads| Mux
    ContentMod -->|Persist| Supabase
    GradingMod -->|Save Scores| Supabase
    PaymentLogic -->|Verify Transfer| Supabase

    %% Class
    class Admin,Student user;
    class Router,AuthMod,ContentMod,GradingMod,PaymentLogic go;
    class Supabase db;
    class Mux ext;
```

## 🌟 Key Features

### 1. Expert-Led Mentorship

- **Manual Grading Workflow**: Beyond multiple-choice quizzes, HocViet supports essay and file submissions where Experts manually review and grade work. This mimics real-world "code review" or "design critique" processes.
- **Direct Feedback**: Mentors can provide specific feedback on student submissions.

### 2. Built for Vietnam 🇻🇳

- **VietQR & Manual Payments**: Integrated "Wizard of Oz" payment flow allowing learners to pay via Bank Transfer (VietQR) or Momo, with manual admin verification.
- **Mobile Optimized**: Video delivery via Mux ensures low-latency streaming even on 4G networks in remote areas.

### 3. Technical Excellence

- **Multi-Tenancy**: Data isolation is enforced at the database level using Postgres Row-Level Security (RLS). An expert from "Academy A" cannot see students from "University B".
- **Golang Performance**: Built on a Modular Monolith architecture for high concurrency and easy deployment.

## 🛠 Tech Stack

- **Backend**: Golang 1.21+ (Gin Framework)
- **Database**: Supabase (PostgreSQL)
- **Authentication**: Supabase Auth (JWT)
- **Video Engine**: Mux (Streaming & Encoding)
- **Frontend**: Next.js (App Router)

## 🚀 Getting Started

### Prerequisites

- Go 1.21+ installed
- A Supabase project
- A Mux account

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/your-username/hocviet-backend.git
   cd hocviet-backend
   ```

2. Install Dependencies
   ```bash
   go mod download
   ```

3. Environment Setup

   Create a `.env` file in the root directory:
   ```env
   PORT=8080
   DATABASE_URL="postgres://postgres:[YOUR-PASSWORD]@db.[YOUR-ID].supabase.co:5432/postgres"
   MUX_TOKEN_ID="your-mux-id"
   MUX_TOKEN_SECRET="your-mux-secret"
   ```

4. Run the Server
   ```bash
   go run cmd/server/main.go
   ```

5. Verify

   Visit `http://localhost:8080/api/v1/health` to confirm the server is running.

## 📂 Project Layout

The project follows the standard Golang layout for Modular Monoliths:

```
/
├── cmd/server/         # Application entry point (main.go)
├── internal/
│   ├── config/         # Environment configuration
│   ├── middleware/     # Tenant isolation & Auth middleware
│   └── modules/        # Domain logic
│       ├── auth/       # User identity
│       ├── content/    # Videos & Courses
│       └── grading/    # Assignments & Quizzes
├── pkg/database/       # Shared DB connection
└── go.mod              # Dependencies
```

## 🤝 Contributing

We welcome contributions from the Vietnamese tech community! Whether it's fixing a bug or adding a new payment integration (ZaloPay, ViettelMoney), feel free to open a PR.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
