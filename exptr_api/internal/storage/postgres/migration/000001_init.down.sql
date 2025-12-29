-- Drop indexes on operations
DROP INDEX IF EXISTS idx_operations_deleted_at;
DROP INDEX IF EXISTS idx_operations_created_at;
DROP INDEX IF EXISTS idx_operations_category_id;
DROP INDEX IF EXISTS idx_operations_user_id;

-- Drop indexes on categories
DROP INDEX IF EXISTS idx_categories_deleted_at;
DROP INDEX IF EXISTS idx_categories_created_at;
DROP INDEX IF EXISTS idx_categories_user_id;

-- Drop indexes on users_sessions
DROP INDEX IF EXISTS idx_users_sessions_deleted_at;
DROP INDEX IF EXISTS idx_users_sessions_created_at;
DROP INDEX IF EXISTS idx_users_sessions_user_id;
DROP INDEX IF EXISTS idx_users_sessions_token;

-- Drop indexes on users
DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_email;

-- Drop tables in reverse order of creation (respecting foreign key constraints)
DROP TABLE IF EXISTS operations;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users_sessions;
DROP TABLE IF EXISTS users;

-- Drop the UUID extension (optional, may be used by other databases)
-- DROP EXTENSION IF EXISTS "uuid-ossp";
