SELECT ord.*
FROM imarket_schema.orders ord,
     imarket_schema.products prod
WHERE prod.id = ord.product_id