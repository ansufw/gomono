-- name: GetUser :one
SELECT
    u.*,
    COALESCE(
        JSON_AGG(
            JSON_BUILD_OBJECT(
                'id', r.id,
                'name', r.name,
                'description', r.description
            )
        ) FILTER (WHERE r.id IS NOT NULL),
        '[]'::json
    ) AS roles
FROM users u
LEFT JOIN user_roles ur ON u.id = ur.user_id
LEFT JOIN roles r ON ur.role_id = r.id
WHERE u.id = $1
GROUP BY u.id
LIMIT 1;

-- name: ListUsers :many
SELECT
    u.*,
    COALESCE(
        JSON_AGG(
            JSON_BUILD_OBJECT(
                'id', r.id,
                'name', r.name,
                'description', r.description
            )
        ) FILTER (WHERE r.id IS NOT NULL),
        '[]'::json
    ) AS roles
FROM users u
LEFT JOIN user_roles ur ON u.id = ur.user_id
LEFT JOIN roles r ON ur.role_id = r.id
GROUP BY u.id
ORDER BY u.created_at DESC
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    username,
    first_name,
    last_name,
    phone_number,
    timezone,
    language
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;

