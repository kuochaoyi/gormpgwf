 -- 累積進貨量(all)
 SELECT psi_p_logs.order_id,
    psi_p_logs.product_id,
    psi_p_logs.quantity
   FROM psi_p_logs
  WHERE psi_p_logs.deleted_at IS NULL
  ORDER BY psi_p_logs.order_id DESC;