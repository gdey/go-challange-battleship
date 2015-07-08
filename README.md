# Gopher Go Challange

## Let's play battleship.

Your challenge is to build a [battleship game][wiki_page] that a user can play against the computer.

The game will have the following rules.

1. The game will be build for the terminal.
2. A user will play against the computer. 
3. The computer will place five (5) ships ( out of the six (6) ships available — choosen at random ) on a sixteen by sixteen (16x16) board. The ships can only be place horizontally or vertically; ships can not be placed diagonally. Ships may not intersect each other. 
The ships are:
  1. Aircraft Carrier — which will take up five  (5) contiguous spots.
  2. Battleship       — which will take up four  (4) contiguous spots.
  3. Submarine        — which will take up three (3) contiguous spots.
  4. Destroyer        — which will take up three (3) contiguous spots.
  5. Cruiser          — which will take up three (3) contiguous spots.
  6. Patrol Boat      — which will take up two   (2) contiguous spots.
4. The player will be given six (6) rounds to sink all the computers ships.
  1. Each round will consist of the player entering in five (5) different spots.
  2. After all five (5) locations have been entered the computer will let the player know which shots (if any) hit ships, and which ships the player managed to sink. Granting the points associated with each ship the player managed to sink.
5. If the player manages to sink all the computers ship before the end of the sixth (6th) round, the player wins. Other wise the computer wins. Tally up the score and display it.

<table border=1 cellpadding=2>
<caption>Ship Attributes</caption>
<tr><th>Name</th><th>Size</th><th>Points</th></tr>
<tr><td>Aircraft Carrier</td><td>5</td><td>20</td></tr>
<tr><td>Battleship</td><td>4</td><td>12</td></tr>
<tr><td>Submarine</td><td>3</td><td>6</td></tr>
<tr><td>Destroyer</td><td>3</td><td>6</td></tr>
<tr><td>Crusier</td><td>3</td><td>6</td></tr>
<tr><td>Patrol Boat</td><td>2</td><td>2</td></tr>
</table>

The system should contain a leader board, that is displayed at the beginning of 
the game. At the start of the game, the user should be asked their name. 
This should be used to display the scores of the past 10 players. The table displaying
the name should be aligned nicely, with scores right aligned.

Make sure to store the leader board, so that the next time the game is started it
will be shown as well.


### Extra Credit (You will not be graded on this)

Using [termbox][termboxgo] add some cool ansi graphics to the game.

1. Draw the board, and refresh it instead of redrawing it everytime.
2. *Bonus* Allow the user to enter in the coordinates using the mouse.

[wiki_page]: https://en.wikipedia.org/wiki/Battleship_(game) "Battleship the board game"
[gosqlite]: https://github.com/mattn/go-sqlite3 "SQLite3 driver that supports the go/dbs interface."
[termboxgo]: https://github.com/nsf/termbox-go "Pure Go Termbox Implementation"

