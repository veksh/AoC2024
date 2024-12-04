SELECT count(*) FROM
  (SELECT max(diff) BETWEEN 1 AND 3 AND max(sign) == min(sign) safe FROM 
    (SELECT
      l1.val l1,
      l2.val l2,
      abs(l2.val - l1.val) diff,
      sign(l2.val - l1.val) sign,
      COUNT(CASE WHEN l1.val IS NULL THEN 1 END) OVER (ORDER BY l1.id) AS line
      FROM list l1 LEFT JOIN list l2 ON l1.id + 1 = l2.id)
    GROUP BY line)
  WHERE safe
