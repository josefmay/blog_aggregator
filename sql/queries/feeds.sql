-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, 
        $2, 
        $3, 
        $4, 
        $5,
        $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * from feeds;

-- name: GetFeedFollows :many
SELECT * from feed_follows WHERE user_id = $1;

-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, 
        $2, 
        $3, 
        $4, 
        $5
)
RETURNING *;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows WHERE id = $1;