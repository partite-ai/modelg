-- name: dropTable
DROP TABLE IF EXISTS things;

-- name: createTable
CREATE TABLE things (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  status INTEGER
);

-- name: CreateThing
INSERT INTO things (
  name, status
) VALUES (
  :name, :status
)
RETURNING *;

-- name: FindBy
SELECT * FROM things
WHERE --<when :filter.HasValues
  :filter!
--endwhen
--when :order.HasValue
ORDER BY name :order!
--endwhen

-- name: FindCount
SELECT COUNT(*) FROM things
WHERE --<when :filter.HasValues
  :filter!
--endwhen


-- name: FindByNameAndStatus
SELECT * FROM things
WHERE --<when :name.HasValue
  name = :name
AND --+when :status.HasValue
  status = :status
--endwhen

-- name: Delete
DELETE FROM things WHERE id = :thing_id

-- name: FindActiveByNames
SELECT * FROM things
WHERE status = 1
AND --<when :names.HasValues
  name IN (:names!)
--endwhen
