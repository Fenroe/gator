-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
),
user_name AS (
    SELECT id, name
    FROM users
    WHERE id = $4
),
feed_name AS (
    SELECT id, name
    FROM feeds
    WHERE id = $5
)
SELECT
    inserted_feed_follow.*,
    user_name.name AS user_name,
    feed_name.name AS feed_name
FROM inserted_feed_follow
JOIN user_name ON inserted_feed_follow.user_id = user_name.id
JOIN feed_name ON inserted_feed_follow.feed_id = feed_name.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.id,ff.created_at,ff.updated_at,ff.user_id,ff.feed_id,u.name AS user_name,f.name AS feed_name
FROM feed_follows ff
INNER JOIN users u
ON ff.user_id = u.id
INNER JOIN feeds f
ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.user_id = $1
AND feed_follows.feed_id = feeds.id
AND feeds.url = $2;
