package players

var QueryString string = `
SELECT 
    p.username,
    p.playerName,
    p.playerTagline,
    p.playerRank,
    p.playerStatus,
    COUNT(pm.username) wins,
    round.kills,
    round.assists,
    round.killPerRound,
    round.firstbloods,
    round.aces,
    round.clutches,
    round.mostkill
FROM
    players p
        JOIN
    (SELECT 
        username, status
    FROM
        playerinmatch
    WHERE
        status = 'Victory') pm ON p.username = pm.username
        JOIN
    (SELECT 
        rs.username,
            r.kills,
            r.assists,
            r.kills / r.roundCount killPerRound,
            r.firstbloods,
            r.aces,
            r.clutches,
            MAX(rmostkill.killPerMatch) mostkill
    FROM
        rounds rs
    JOIN (SELECT 
        SUM(killCount) kills,
            SUM(assistCount) assists,
            COUNT(*) roundCount,
            SUM(firstBlood) firstbloods,
            COUNT(CASE
                WHEN killCount = 5 THEN 1
            END) aces,
            SUM(clutch) clutches,
            username
    FROM
        rounds
    GROUP BY username) r ON rs.username = r.username
    JOIN (SELECT 
        SUM(killCount) killPerMatch, username
    FROM
        rounds
    GROUP BY matchID , username) rmostkill ON rmostkill.username = rs.username
    GROUP BY rs.username) round ON p.username = round.username
WHERE p.username LIKE ?
GROUP BY p.username
LIMIT ?,?;`
