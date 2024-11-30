CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL, -- e.g., 'internal', 'external', 'admin'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
