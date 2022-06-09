package players

var GetPlayerDetailQuery string = `
SELECT 
    p.username,
    p.playerName,
    p.playerTagline,
    p.playerRank,
    p.playerStatus,
    pm.wins,
    r.kills,
    r.assists,
    (SELECT r.kills / r.rCount) killPerRound,
    r.firstbloods,
    r.aces,
    r.clutches,
    MAX(rmostkill.killPerMatch) mostkill
FROM
    players p
        JOIN
    (SELECT 
        COUNT(username) wins, status, username
    FROM
        playerinmatch
    WHERE
        status = 'Victory'
    GROUP BY username) pm ON p.username = pm.username
        JOIN
    (SELECT 
        SUM(killCount) kills,
            SUM(assistCount) assists,
            COUNT(*) rCount,
            SUM(firstBlood) firstbloods,
            COUNT(CASE
                WHEN killCount = 5 THEN 1
            END) aces,
            SUM(clutch) clutches,
            username
    FROM
        rounds
    GROUP BY username) r ON p.username = r.username
        JOIN
    (SELECT 
        SUM(killCount) killPerMatch, username
    FROM
        rounds
    GROUP BY matchID , username) rmostkill ON rmostkill.username = p.username
WHERE p.username LIKE ?`

var ModifyPlayerNameQuery string = `
UPDATE players 
SET 
    playerName = ?
WHERE
    username = ?`

var ModifyTaglineQuery string = `
UPDATE players 
SET 
    playerTagline = ?
WHERE
    username = ?`
