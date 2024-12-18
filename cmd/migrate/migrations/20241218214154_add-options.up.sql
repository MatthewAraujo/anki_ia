CREATE TABLE IF NOT EXISTS options (
  id SERIAL PRIMARY KEY,
  question_id INT NOT NULL,
  option_key CHAR(1) NOT NULL, 
  option_text TEXT NOT NULL,
  is_correct BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);
