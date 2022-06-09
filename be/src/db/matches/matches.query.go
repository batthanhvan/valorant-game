package matches

var MatchesDetailQuery string = `
SELECT 
    m.matchID,
    m.matchServer,
    m.mapName,
    m.modeName,
    m.startTime,
    m.endTime,
    m.recordLink
FROM
    matches m
        JOIN
    (SELECT 
        username, matchID
    FROM
        playerinmatch
    WHERE
        username LIKE ?) pm ON m.matchID = pm.matchID
ORDER BY m.startTime DESC
LIMIT ?,?;`
