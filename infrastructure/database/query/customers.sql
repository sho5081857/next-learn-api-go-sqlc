-- name: GetAllCustomers :many
SELECT id,
    name
FROM customers
ORDER BY name ASC;

-- name: GetFilteredCustomers :many
SELECT customers.id,
    customers.name,
    customers.email,
    customers.image_url,
    COUNT(invoices.id) AS total_invoices,
    SUM(
        CASE
            WHEN invoices.status = 'pending' THEN invoices.amount
            ELSE 0
        END
    ) AS total_pending,
    SUM(
        CASE
            WHEN invoices.status = 'paid' THEN invoices.amount
            ELSE 0
        END
    ) AS total_paid
FROM customers
    LEFT JOIN invoices ON customers.id = invoices.customer_id
WHERE customers.name ILIKE $1
    OR customers.email ILIKE $1
GROUP BY customers.id,
    customers.name,
    customers.email,
    customers.image_url
ORDER BY customers.name ASC;

-- name: GetCustomerCount :one
SELECT COUNT(*) FROM customers;
