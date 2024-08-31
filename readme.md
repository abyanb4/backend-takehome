
# How To Run
1. `make run` for development and `make run-prod` for production. There is no big difference actually, I just separate it to be production-ready.
2. Voila, the server run on port `8080`. The migration invoked before app run, at this point the table should be ready
3. You can test the endpoint using test script that I provided on zip `backend-takehome-testscript`. Make sure you read the readme first because there is some manual step need to be done.
<br>

# Indexing
Here I use index only on several query that I predict will have a big traffic
Query:
1. User Login, register -> check email and username
```
SELECT EXISTS(SELECT 1 FROM users WHERE email = "");
SELECT EXISTS(SELECT 1 FROM users WHERE username = "");
```

2. Get Comment By Post ID
``` 
SELECT id, post_id, author_name, content, created_at, updated_at FROM comments WHERE post_id = ? ORDER BY created_at DESC
```

3. Get Post By Author ID
```
SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE author_id = ? ORDER BY created_at DESC
```

Index:
```
CREATE INDEX idx_posts_author_id_created_at ON posts;
CREATE INDEX idx_comments_post_id_created_at ON comments;
CREATE INDEX idx_users_username ON users;
CREATE INDEX idx_users_email ON users;
```

# Common Error
1.  "error": "failed to check email uniquenessError 1146 (42S02): Table 'blog_service.users' doesn't exist"

Just re run `make run`. This usually Occur on first build. The migration script failed because MySQL usually take more time to accept connection on first build.


# Improvement Items
- Air hot reloading
- Migration Down
- Better Error Trace and error logging
- Better Response Format
- Token Invalidation using cache -> Already implemented the cache, only need to work on the invalidation