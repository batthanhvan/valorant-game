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
ORDER BY m.startTime DESC
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
ORDER BY m.startTime DESC
LIMIT ?,?;`

var PostNewReport string = `
INSERT INTO reports(username, reportCategory, matchID, reportDetail)
values(?,?,?,?)
`
var CheckPlayerInMatch string = `
SELECT 1
FROM
    valorant.playerinmatch
WHERE
    username = ?
        AND matchID = ?
`
