ALTER TABLE users
ADD COLUMN role VARCHAR(20) NOT NULL DEFAULT 'free';

ALTER TABLE users
ADD CONSTRAINT chk_user_role 
CHECK (role IN ('free', 'mid', 'unlimited'));
