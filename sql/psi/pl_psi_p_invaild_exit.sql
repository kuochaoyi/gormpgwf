-- 作廢
-- 減少庫存量
CREATE OR REPLACE FUNCTION p_quantity_minus(prouduct_id uuid, quantity int) 
    RETURNS void
AS $$
DECLARE
    q int := 0;
BEGIN
    SELECT INTO q psi_p_stock.quantity FROM psi_p_stock 
		WHERE prouduct_id = psi_p_stock.uuid_id;
    q = q - quantity;
    UPDATE psi_p_stock SET quantity = q, updated_at = now()
		WHERE prouduct_id = psi_p_stock.uuid_id;
END;
$$ LANGUAGE plpgsql;

-- SELECT public.quantity_add('01159a58-eb2e-40e1-8f70-119b9a249c7a'::UUID, 50);

-- 作廢單
CREATE OR REPLACE FUNCTION p_invalid(order_no text)
    RETURNS integer AS $$
DECLARE
    r psi_p_logs%ROWTYPE;
	n RECORD; -- IS NULL
BEGIN
	-- order_no: No match or deleted.
	SELECT * INTO n FROM psi_p_logs 
        WHERE order_id = order_no AND deleted_at IS NULL;
	IF n IS NULL THEN RETURN 0; END IF;
	
    FOR r IN (SELECT * FROM psi_p_logs 
        WHERE order_id = order_no AND deleted_at IS NULL)
    LOOP
        PERFORM p_quantity_minus(r.product_id, r.quantity);
    END LOOP;
    
    UPDATE psi_p_logs SET deleted_at = now() WHERE order_id = order_no;
    RETURN 1;
END;
$$ LANGUAGE plpgsql;

-- CREATE OR REPLACE FUNCTION p_invalid(order_no text)
--     RETURNS void AS $$
-- DECLARE
--     c CURSOR;
--     r psi_p_logs%ROWTYPE;
-- BEGIN
--     OPEN curs FOR SELECT * FROM psi_p_logs WHERE order_id = order_no;
--     FETCH curs INTO r;
--         SELECT p_quantity_minus(r.product_id, r.quantity);
--         -- SELECT prouduct_id, quantity INTO r.prouduct_id, r.quantity FROM psi_p_logs WHERE order_no = order_id
--         -- SELECT product_id, quantity INTO r FROM psi_p_logs WHERE order_no = order_id
--     --CLOSE curs;
    
--     UPDATE psi_p_logs SET deleted_at = now() FROM psi_p_logs WHERE order_id = order_no;
--     -- RETURN 1;
-- END;
-- $$ LANGUAGE plpgsql;