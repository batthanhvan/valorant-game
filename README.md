# SQL project

- Register: /register
- Login: /login (=> Got **token** here)
- Get user info: /user

- Show all players report: /admin/reports/show
- Get reports by username: /admin/reports/search?username=... (add **token** to validate)

- Get player detail by username: /players/:username
- Change playername and tagline: /modify/:playername/:tagline (POST method) (add **token** to validate)

- Get match by username: /matches/search?username=...