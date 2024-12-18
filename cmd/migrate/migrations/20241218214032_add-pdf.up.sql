CREATE TABLE IF NOT EXISTS pdfs (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  filename VARCHAR(255) NOT NULL,
  uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  status VARCHAR(50) CHECK (status IN ('pending', 'processed', 'failed')) DEFAULT 'pending',
  text_content TEXT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);