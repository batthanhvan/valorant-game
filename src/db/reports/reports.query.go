package reports

var CountAllReport string = `
SELECT 
    COUNT(*)
FROM
    reports r`

var GetAllReport string = `
SELECT 
    r.reportCategory, r.reportDetail, DATE(m.startTime), m.recordLink, r.username
FROM
    reports r
        JOIN
    matches m ON m.matchID = r.matchID
ORDER BY (SELECT 
        startTime
    FROM
        matches m
    WHERE
        m.matchID = r.matchID)
LIMIT ?,?`

var CountReport string = `
SELECT 
    COUNT(*)
FROM
    reports r
WHERE r.username LIKE ?`

var GetReportByUsername string = `
SELECT 
    r.reportCategory, r.reportDetail, DATE(m.startTime), m.recordLink, r.username
FROM
    reports r
        JOIN
    matches m ON m.matchID = r.matchID
WHERE r.username LIKE ?
ORDER BY (SELECT 
        startTime
    FROM
        matches m
    WHERE
        m.matchID = r.matchID)
LIMIT ?,?;`
