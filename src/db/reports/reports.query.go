package reports

var QueryAllString string = `
SELECT 
    r.reportCategory, r.reportDetail, DATE(m.startTime), m.recordLink, r.username
FROM
    reports r
        JOIN
    matches m ON m.matchID = r.matchID
LIMIT ?,?;`

var QueryString string = `
SELECT 
    r.reportCategory, r.reportDetail, DATE(m.startTime), m.recordLink, r.username
FROM
    reports r
        JOIN
    matches m ON m.matchID = r.matchID
WHERE r.username LIKE ?
LIMIT ?,?;`
