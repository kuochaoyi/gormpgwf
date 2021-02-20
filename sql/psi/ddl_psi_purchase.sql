-- 進貨/庫存
-- 
CREATE TABLE psi_p_stock
(
    uuid_id uuid NOT NULL DEFAULT gen_random_uuid(),
    product_no varchar,
    quantity integer,
    -- 
    created_date date DEFAULT CURRENT_DATE,
    created_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone,
    PRIMARY KEY (uuid_id)
)

-- 初始化庫存及量
INSERT INTO psi_p_stock(product_no, quantity) VALUES ('A書', 5);
INSERT INTO psi_p_stock(product_no, quantity) VALUES ('飛機', 15);

-- 進貨記錄
CREATE TABLE psi_p_logs
(
    uuid_id uuid NOT NULL DEFAULT gen_random_uuid(),
    order_id varchar,
    product_id uuid, -- FK: psi_p_stock.uuid_id
    quantity integer,
    unit_price numeric,
    --
    created_date date DEFAULT CURRENT_DATE,
    created_at timestamp with time zone DEFAULT now(), -- type: timestampz
    deleted_at timestamp with time zone,
    PRIMARY KEY (uuid_id)
)

/* 同步用參考
INSERT INTO public.psi_p_logs(order_id, product_id, quantity, unit_price)
	VALUES ('20200214001', '9f6734f6-5806-4973-8baf-f2282b52da83', 5, 100);
    5f3d2e0d-3cc8-41e1-a17a-803c4975b0d3
*/