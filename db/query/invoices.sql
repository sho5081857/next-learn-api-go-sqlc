-- name: GetLatestInvoices :many
SELECT invoices.amount,
    customers.name,
    customers.image_url,
    customers.email,
    invoices.id
FROM invoices
    JOIN customers ON invoices.customer_id = customers.id
ORDER BY invoices.date DESC
LIMIT 5;

-- name: GetFilteredInvoices :many
SELECT invoices.id,
    invoices.amount,
    invoices.date,
    invoices.status,
    customers.name,
    customers.email,
    customers.image_url
FROM invoices
    JOIN customers ON invoices.customer_id = customers.id
WHERE customers.name ILIKE $1
    OR customers.email ILIKE $1
    OR invoices.amount::text ILIKE $1
    OR invoices.date::text ILIKE $1
    OR invoices.status ILIKE $1
ORDER BY invoices.date DESC
LIMIT $2 OFFSET $3;

-- name: GetInvoicesPages :one
SELECT COUNT(*)
FROM invoices
    JOIN customers ON invoices.customer_id = customers.id
WHERE customers.name ILIKE $1
    OR customers.email ILIKE $1
    OR invoices.amount::text ILIKE $1
    OR invoices.date::text ILIKE $1
    OR invoices.status ILIKE $1;

-- name: GetInvoiceById :one
SELECT invoices.id,
    invoices.customer_id,
    invoices.amount,
    invoices.status
FROM invoices
WHERE invoices.id = $1;

-- name: GetInvoiceCount :one
SELECT COUNT(*)
FROM invoices;

-- name: GetInvoiceStatusCount :one
SELECT SUM(
        CASE
            WHEN status = 'paid' THEN amount
            ELSE 0
        END
    ) AS "paid",
    SUM(
        CASE
            WHEN status = 'pending' THEN amount
            ELSE 0
        END
    ) AS "pending"
FROM invoices;

-- name: CreateInvoice :exec
INSERT INTO invoices (customer_id, amount, status, date)
VALUES ($1, $2, $3, $4);
-- name: UpdateInvoice :exec
UPDATE invoices
SET customer_id = $2,
    amount = $3,
    status = $4
WHERE id = $1;

-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1;
