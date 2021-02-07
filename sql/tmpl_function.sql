CREATE OR REPLACE FUNCTION [預存函式名稱] (_param1 int, param2 int) RETURNS [回傳型態] AS
$$
    DECLARE --變數宣告型態區域 (變數一定要事先宣告)
        _id                  BIGINT;
        _screens             JSONB;
 
    BEGIN --宣告型態之後
         
        SELECT x INTO _screens FROM table WHERE 1=1; --注意, 這是把搜尋的資料指派進去 _screen 這個變數。
 
 
        IF _screens IS NULL THEN -- 簡易判斷 _screen 這個變數是不是沒有查到資料
            RETURN jsonb_build_object('status', 'ERROR', 'message', 'UUID_NOT_FOUND');
        END IF;
 
        FOR i IN 0..jsonb_array_length(_screens) - 1 LOOP --迴圈
            _screen_id = _screens -> i ->> 'screen_id';
            _times := '[]';
            FOR
                _id
            IN SELECT
                id
            FROM reservation WHERE 1=1
 
            LOOP --迭代迴圈區域
                _times = jsonb_set(_times, '{2147483647}', jsonb_build_object(
                    'id', _id,
                    )
                );
 
            END LOOP;
 
            _data_array = array_append(_data_array, jsonb_build_object('screen_id', _screen_id, 'times', _times));
 
        END LOOP;
 
        RETURN json_build_object('status', 'SUCCESS', 'screens', _data_array);
 
        EXCEPTION --例外處理
            WHEN OTHERS THEN
                RETURN jsonb_build_object('status', 'ERROR', 'message', 'UN_KNOWN');
    END; --宣告結束
$$
LANGUAGE 'plpgsql';