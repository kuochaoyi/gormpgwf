/*

*/
CREATE TABLE template
(
    uuid_id      UUID        DEFAULT gen_random_uuid(),
    data         JSONB,

/* somethings
    qty             int,            --  quantity
    price           ,
   //
   -- gorm Price decimal.Decimal `json:"price" gorm:"type:numeric"`
    amount          money,          -- gorm -> float64 `gorm:"type:money"`
    subtotal        int,            -- 小計
    total           int,
 */

-- logs gorm: *time.Time
    created_date DATE        DEFAULT CURRENT_DATE,
--  created_at    TIMESTAMP       DEFAULT now(),
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    deleted_at   TIMESTAMPTZ,
--     is_deleted      BOOL            DEFAULT false,

    PRIMARY KEY (uuid_id)
);

-- INSERT INTO template(data) VALUES ('{"title": "Sleeping Beauties", "genres": ["Fiction", "Thriller", "Horror"], "published": false}');