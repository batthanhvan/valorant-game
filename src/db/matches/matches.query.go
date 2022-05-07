package matches

var QueryString string = `
SELECT 
    m.matchID, m.matchServer, m.mapName, m.modeName, m.startTime, m.endTime, m.recordLink
FROM
    matches m
WHERE p.username LIKE ?
LIMIT ?,?;`
