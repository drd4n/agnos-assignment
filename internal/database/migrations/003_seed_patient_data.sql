-- Migration: 003_seed_patient_data.sql
-- Seed sample patient data for Bangkok Medical Center

INSERT INTO patients (first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender)
VALUES
    ('สมชาย', '', 'ใจสว่าง', 'Somchai', '', 'Jaisawang', '1985-03-15', 'Bangkok Medical Center', '1234567890123', 'N12345678', '0812345678', 'somchai@example.com', 'M'),
    ('สุจริตา', 'มิ่ง', 'ชาตรี', 'Sujarita', 'Ming', 'Chatri', '1990-07-22', 'HN000002', '9876543210987', 'N98765432', '0898765432', 'sujarita@example.com', 'F'),
    ('สมศักดิ์', '', 'วิทยา', 'Somsak', '', 'Withaya', '1978-11-08', 'HN000003', '5555555555555', NULL, '0865555555', 'somsak@example.com', 'M'),
    ('สายใจ', '', 'คุณธรรม', 'Saijai', '', 'Kuntharam', '1992-05-30', 'HN000004', '4444444444444', 'N44444444', '0884444444', 'saijai@example.com', 'F')
ON CONFLICT DO NOTHING;
