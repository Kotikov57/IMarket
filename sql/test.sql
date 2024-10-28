SELECT ord.*
FROM imarket_db.public.users us,
     imarket_db.public.orders ord
WHERE  us.id = ord.user_id
ORDER BY user_id