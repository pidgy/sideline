# sideline
A fantasy football prediction engine written in Go. 

A simple engine to consume RESTful endpoints and make predictions for multiple lineups as well as predict weekly winners.

Version 0.1 -> Currently only supporting https://www.fantasyfootballnerd.com/fantasy-football-api

```
go run coach.go -api-key <your FFN api key>
```

The prediction model is based on the Academic paper: 
`Becker, A. and X.A. Sun. 2016. An analytical approach for fantasy football draft and lineup management`

Focusing on the two core methodologies presented by Becker and X.A. Sun.:

```
1. A holistic optimization model which manages a team
through draft construction and weekly management.
2. The analysis of a player’s historical statistical performance
on a weekly basis in the context of the player’s
opponents; and the ability to make predictions on this
analysis.
```

* Parameters and decision variables
1. Parameters:
```
N: The set of NFL players and defensive teams.
M: The set of positions M = {QB, RB, WR, TE, K, Def}.
T: The set of weeks in the NFL regular and playoff seasons.
Pos(i): The position of player i, e.g. Pos(PeytonManning)
= QB.
PosLimit(j): The upper bound on the number of starting
players for position j ∈ M during the weekly play and
playoff phases, e.g. PosLimit(QB) = 1.
nk: The overall pick number of the DM’s k-th draft pick.
DMPlayer(k): The set of players that the DM has drafted
by her k-th pick.
OppPlayer(k): The set of players that the opponents have
drafted by the DM’s k-th pick.
Rk(i): Anticipated ranking of unselected player i at the
DM’s k-th draft pick.
f(i, t): An estimate of the number of fantasy points player i
will score in week t of the NFL regular season.
β(t): The predicted amount of fantasy points the DM
needs so to be reasonably confident to win in week t ∈ T
against the DM’s matchup opponent.
γj: The number of players the DM must draft at position j.
```

2. Decision variables:
```
–– yi ∈ {0, 1}: yi = 1 if the DM picks player i in the draft phase, and yi = 0 otherwise.
–– x{_i}{^t} ∈ {0, 1} : x{_i}{^t} = 1 if the DM starts player i in week t, and x{_i}{^t} = 0 otherwise.
zt ∈ {0, 1}: zt = 1 if the total estimated fantasy points if the DM’s line-up in week t is greater than β(t) 
```
