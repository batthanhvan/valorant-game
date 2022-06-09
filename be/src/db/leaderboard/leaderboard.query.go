package leaderboard

var LeaderboardDetail string = `
SELECT 
    p.username,
    p.playerName,
    p.playerTagline,
    p.playerRank,
    ROUND((SELECT round.kills / round.rCount) * 1000) rating
FROM
    players p
        JOIN
    (SELECT 
        rs.username, SUM(killCount) kills, COUNT(*) rCount, time.startTime
    FROM
        rounds rs
    JOIN (SELECT 
        startTime, pm.username
    FROM
        matches m
    JOIN (SELECT 
        username, matchID
    FROM
        playerinmatch) pm ON pm.matchID = m.matchID
    GROUP BY username) time ON time.username = rs.username
    GROUP BY rs.username) round ON p.username = round.username
WHERE
    YEAR(round.startTime) = (SELECT 
            YEAR(MAX(round.startTime))
        FROM
            matches)
        OR YEAR(round.startTime) = (SELECT 
            YEAR(MAX(round.startTime))
        FROM
            matches) - 1
GROUP BY p.username
ORDER BY rating DESC
LIMIT ?,?
`
var CountPlayerLeaderboard string = `
SELECT 
    COUNT(p.username)
FROM
    players p
        JOIN
    (SELECT 
        rs.username, SUM(killCount) kills, COUNT(*) rCount, time.startTime
    FROM
        rounds rs
    JOIN (SELECT 
        startTime, pm.username
    FROM
        matches m
    JOIN (SELECT 
        username, matchID
    FROM
        playerinmatch) pm ON pm.matchID = m.matchID
    GROUP BY username) time ON time.username = rs.username
    GROUP BY rs.username) round ON p.username = round.username
WHERE
    YEAR(round.startTime) = (SELECT 
            YEAR(MAX(round.startTime))
        FROM
            matches)
        OR YEAR(round.startTime) = (SELECT 
            YEAR(MAX(round.startTime))
        FROM
            matches) - 1
`
