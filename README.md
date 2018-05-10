Simple Tic Tac Toe with option for multiple players and board size from 3x3 to 10x10

Player signs and board size defined in config.txt

##To run use
## go run main.go

##Test cases and results :
## go test ./...



=== RUN   TestReadArgs
=== RUN   TestReadArgs/Should_return_given_size_as_board_size_without_error
board size : 10
Let the game begin !!
--- PASS: TestReadArgs (0.00s)
    --- PASS: TestReadArgs/Should_return_given_size_as_board_size_without_error (0.00s)
PASS
ok  	learn/tic	(cached)
=== RUN   TestCheckRowColumnDiagWin
=== RUN   TestCheckRowColumnDiagWin/should_return_zero
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_wins
=== RUN   TestCheckRowColumnDiagWin/should_return_my_row_3_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_row_1_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_row_2_win
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_col_1_win
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_col_2_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_col_3_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_col_1_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_col_2_win
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_col_3_win
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_diag_right_win
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_diag_left_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_diag_right_win
=== RUN   TestCheckRowColumnDiagWin/should_return_my_diag_left_win
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_win_with_2_opponents_for_row
=== RUN   TestCheckRowColumnDiagWin/should_return_opponent_win_with_2_opponents_for_col
--- PASS: TestCheckRowColumnDiagWin (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_zero (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_wins (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_row_3_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_row_1_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_row_2_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_col_1_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_col_2_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_col_3_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_col_1_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_col_2_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_col_3_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_diag_right_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_diag_left_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_diag_right_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_my_diag_left_win (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_win_with_2_opponents_for_row (0.00s)
    --- PASS: TestCheckRowColumnDiagWin/should_return_opponent_win_with_2_opponents_for_col (0.00s)
=== RUN   TestWeightMove
=== RUN   TestWeightMove/no_one_wins
=== RUN   TestWeightMove/my_win_by_row
=== RUN   TestWeightMove/opposite_win_by_row
=== RUN   TestWeightMove/my_win_by_column
=== RUN   TestWeightMove/opposite_win_by_column
=== RUN   TestWeightMove/my_win_by_diag
--- PASS: TestWeightMove (0.00s)
    --- PASS: TestWeightMove/no_one_wins (0.00s)
    --- PASS: TestWeightMove/my_win_by_row (0.00s)
    --- PASS: TestWeightMove/opposite_win_by_row (0.00s)
    --- PASS: TestWeightMove/my_win_by_column (0.00s)
    --- PASS: TestWeightMove/opposite_win_by_column (0.00s)
    --- PASS: TestWeightMove/my_win_by_diag (0.00s)
=== RUN   TestIsRemaining
=== RUN   TestIsRemaining/no_moves_remaining_when_all_are_filled
=== RUN   TestIsRemaining/remaining_moves_when_multiple_are_free
=== RUN   TestIsRemaining/remaining_moves_when_all_are_filled_except_1
=== RUN   TestIsRemaining/remaining_moves_when_complete_board_is_empty
--- PASS: TestIsRemaining (0.00s)
    --- PASS: TestIsRemaining/no_moves_remaining_when_all_are_filled (0.00s)
    --- PASS: TestIsRemaining/remaining_moves_when_multiple_are_free (0.00s)
    --- PASS: TestIsRemaining/remaining_moves_when_all_are_filled_except_1 (0.00s)
    --- PASS: TestIsRemaining/remaining_moves_when_complete_board_is_empty (0.00s)
=== RUN   TestMinMax
=== RUN   TestMinMax/return_max_and_min_of_two_numbers
--- PASS: TestMinMax (0.00s)
    --- PASS: TestMinMax/return_max_and_min_of_two_numbers (0.00s)
=== RUN   TestFinder
=== RUN   TestFinder/should_return_opponents_win_for_the_given_board
=== RUN   TestFinder/should_return_my_win_for_the_given_diag_board
=== RUN   TestFinder/should_return_my_win_for_the_given_board
=== RUN   TestFinder/should_return_oponnents_win_for_the_given_row_board
=== RUN   TestFinder/should_return_oponnents_win_for_the_given_diag_board
=== RUN   TestFinder/should_return_oponnents_win_for_the_given_diag_part_empty_board
=== RUN   TestFinder/should_return_my_win_for_the_given_row_part_empty_board
=== RUN   TestFinder/should_return_draw_for_the_given_non_empty_board
=== RUN   TestFinder/should_return_my_possible_win_(row-column)_when_its_my_turn
=== RUN   TestFinder/should_return_my_possible_win_(column)_when_its_my_turn
=== RUN   TestFinder/should_return_opponents_win_(column)_when_its_opponents_turn
=== RUN   TestFinder/should_return_opponents_win_(column)_when_its_opponents_turn#01
--- PASS: TestFinder (0.00s)
    --- PASS: TestFinder/should_return_opponents_win_for_the_given_board (0.00s)
    --- PASS: TestFinder/should_return_my_win_for_the_given_diag_board (0.00s)
    --- PASS: TestFinder/should_return_my_win_for_the_given_board (0.00s)
    --- PASS: TestFinder/should_return_oponnents_win_for_the_given_row_board (0.00s)
    --- PASS: TestFinder/should_return_oponnents_win_for_the_given_diag_board (0.00s)
    --- PASS: TestFinder/should_return_oponnents_win_for_the_given_diag_part_empty_board (0.00s)
    --- PASS: TestFinder/should_return_my_win_for_the_given_row_part_empty_board (0.00s)
    --- PASS: TestFinder/should_return_draw_for_the_given_non_empty_board (0.00s)
    --- PASS: TestFinder/should_return_my_possible_win_(row-column)_when_its_my_turn (0.00s)
    --- PASS: TestFinder/should_return_my_possible_win_(column)_when_its_my_turn (0.00s)
    --- PASS: TestFinder/should_return_opponents_win_(column)_when_its_opponents_turn (0.00s)
    --- PASS: TestFinder/should_return_opponents_win_(column)_when_its_opponents_turn#01 (0.00s)
=== RUN   TestPlay
=== RUN   TestPlay/should_return_the_winning_move_for_computer
[x x o]
[o x _]
[o x _]
=== RUN   TestPlay/should_return_the_winning_move_for_computer#01
[x x x]
[o o _]
[o x _]
=== RUN   TestPlay/should_return_the_move_not_letting_opponent_win
[x _ _]
[o o x]
[_ _ x]
=== RUN   TestPlay/should_return_a_move_for_the_given_board
[x _ _]
[_ o _]
[_ _ _]
--- PASS: TestPlay (0.01s)
    --- PASS: TestPlay/should_return_the_winning_move_for_computer (0.00s)
    --- PASS: TestPlay/should_return_the_winning_move_for_computer#01 (0.00s)
    --- PASS: TestPlay/should_return_the_move_not_letting_opponent_win (0.00s)
    --- PASS: TestPlay/should_return_a_move_for_the_given_board (0.01s)
PASS
ok  	learn/tic/players	(cached)

Process finished with exit code 0
