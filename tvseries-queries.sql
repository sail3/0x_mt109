-- 1. Which actors play in the series Big Sister?

SELECT * FROM employee tb1
INNER JOIN actor_serie tb2 ON tb1.id = tb2.actor_id
INNER JOIN serie tb3 ON tb2.serie_id = tb3.id
WHERE tb3.name = "Big Sister"

-- 2. Which director has directed the greatest number of episodes?

SELECT tb2.id, count(*) as directed FROM episodes tb1
INNER JOIN employee tb2 ON tb1.director_id = tb2.id
GROUP BY tb2.id
ORDER BY directed DESC
LIMIT 1