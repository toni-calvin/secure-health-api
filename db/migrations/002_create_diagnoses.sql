CREATE TABLE diagnoses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    diagnosis TEXT NOT NULL,
    prescription TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
