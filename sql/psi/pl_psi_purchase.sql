-- 增加庫存量
CREATE OR REPLACE FUNCTION p_quantity_add(prouduct_id uuid, quantity int) 
    RETURNS void
AS $$
DECLARE
    q int := 0;
BEGIN
    SELECT INTO q psi_p_stock.quantity FROM psi_p_stock 
		WHERE prouduct_id = psi_p_stock.uuid_id;
    q = q + quantity;
    UPDATE psi_p_stock SET quantity = q, updated_at = now()
		WHERE prouduct_id = psi_p_stock.uuid_id;
END;
$$ LANGUAGE plpgsql;

-- SELECT public.quantity_add('01159a58-eb2e-40e1-8f70-119b9a249c7a'::UUID, 50);

-- 單號+1 20200214001
CREATE OR REPLACE FUNCTION p_order_no_new()
    RETURNS text
    LANGUAGE 'plpgsql'

AS $BODY$
DECLARE
    perfix_data text;
    v text;
	i int8;
BEGIN
    SELECT INTO perfix_data to_char(CURRENT_DATE, 'YYYYMMDD');
	-- RAISE NOTICE '1 %', perfix_data;
	
    SELECT INTO v max(order_id) FROM psi_p_logs WHERE order_id LIKE perfix_data || '%';
    IF v IS NULL THEN 
		v = to_char(CURRENT_DATE, 'YYYYMMDD001');
		-- RAISE NOTICE '2 %', v;
		RETURN v;
		-- RAISE NOTICE '{order_no, %}', v;
		-- v = '{order_no, ' || v || '}';
		--RAISE NOTICE 'v %', v;
		-- RETURN json_object(v::text[]);
    END IF;

	i = v::int8 + 1;
	v = i::text;
	-- v = '{order_no, ' || v || '}';
	-- RAISE NOTICE 'v %', v;
	-- RETURN json_object(v::text[]);
    RETURN v;
END;
$BODY$;

-- 新增進貨單記錄
CREATE OR REPLACE FUNCTION p_insert(IN in_array text[])
    RETURNS text
AS $$
DECLARE
    t text[];
	pno_new text;
	t1 uuid;
	t2 integer;
BEGIN
	SELECT INTO pno_new p_order_no_new();
    FOREACH t SLICE 1 IN ARRAY in_array LOOP
        -- raise notice 't: %', t[1];
		-- raise notice 'b: %, %', t[1], t[2];
		t1 = t[1]::uuid;
		t2 = t[2]::integer;
		INSERT INTO psi_p_logs(order_id, product_id, quantity, unit_price)
	        -- VALUES (t[1], t2, t3, t[4]::numeric);
			VALUES (pno_new, t1, t2, t[3]::numeric);
		-- INSERT INTO psi_p_logs (order_id, quantity) 
			-- VALUES (t[2], t[3]::integer);
		PERFORM p_quantity_add(t1, t2);
    END LOOP;
    RETURN pno_new;
END;
$$ LANGUAGE plpgsql;

-- TODO: exception.

/*
SELECT p_insert('{
    {bcd16bb9-1f72-47af-90af-2e8ce2cf0668, 15, 300},
    {444d77a1-cb11-4444-a82c-dd71645f4858, 30, 100}
    }');

SELECT p_stock_insert('{
    {9f6734f6-5806-4973-8baf-f2282b52da83, 15, 300},
    {5f3d2e0d-3cc8-41e1-a17a-803c4975b0d3, 30, 100}
    }');

SELECT sync_insert_stock('{{a,b,c},{d,e,f},{g,h,i}}');

SELECT public.sync_insert_stock('{
    {20210215003, 9f6734f6-5806-4973-8baf-f2282b52da83, 15, 300},
    {20210215003, 5f3d2e0d-3cc8-41e1-a17a-803c4975b0d3, 30, 100}
    }');
*/


-- 作廢進單用(減少庫存量)
CREATE OR REPLACE FUNCTION quantity_minus(prouduct_id uuid, quantity int) 
    RETURNS void
AS $$
DECLARE
    q int :=0;
BEGIN
    SELECT INTO q psi_p_stock.quantity FROM psi_p_stock 
		WHERE prouduct_id = psi_p_stock.uuid_id;
    q = q - quantity;
    UPDATE psi_p_stock SET quantity = q, updated_at = now()
		WHERE prouduct_id = psi_p_stock.uuid_id;
END;
$$ LANGUAGE plpgsql;