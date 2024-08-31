CREATE INDEX idx_posts_author_id_created_at ON posts(author_id, created_at DESC);
CREATE INDEX idx_comments_post_id_created_at ON comments(post_id, created_at DESC);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
