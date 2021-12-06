{-# LANGUAGE TupleSections #-}
{-# LANGUAGE TupleSections #-}
{-# OPTIONS_GHC -Wno-incomplete-patterns #-}

import Data.List(find, sort, intersect, subsequences, (\\))
import Debug.Trace (trace)
import qualified Data.Map.Strict as M
import qualified Data.Set as S
import Data.Map.Strict (empty, keys, union, (!))
import Data.Time.Format.ISO8601 (yearFormat)
import Data.Maybe

main = interact main'

wordsWhen     :: (Char -> Bool) -> String -> [String]
wordsWhen p s =  case dropWhile p s of
                      "" -> []
                      s' -> w : wordsWhen p s''
                            where (w, s'') = break p s'

main' :: String -> String
main' input = show output
    where (list, rowToBoard) = parseInput $ lines input
          rowToMarked = M.map (const 5) rowToBoard
          output = findLastToWin list rowToBoard [] rowToMarked (quot (length rowToBoard) 10)

findFirstToWin :: [Int] -> M.Map [Int] [Int] -> [Int] -> Int
findFirstToWin [] m (x:xs) = x
findFirstToWin (x:xs) m list | allMarked /= [] = x * sum (board \\ (x:list))
  where board = m ! head allMarked
        allMarked = k `intersect` subs
        k = keys m
        subs = map sort (filter (\s -> length s == 5) (subsequences (x:list)))
findFirstToWin (x:xs) m list = findFirstToWin xs m (x:list)

findLastToWin :: [Int] -> M.Map [Int] [Int] -> [Int] -> M.Map [Int] Int -> Int -> Int
findLastToWin [] rowToBoard (x:xs) rowToMarked boards = x
findLastToWin (x:xs) rowToBoard list rowToMarked boards | boards > 1 && matchedRows /= [] = findLastToWin xs rowToBoard (x:list) u (boards-(quot (length rows) 10))
  where u = M.filterWithKey (\k v -> k `notElem` rows) unmatchedRows
        rows = M.keys $ M.filter (`elem` b) rowToBoard
        b = M.elems $ M.filterWithKey (\k v -> k `elem` matchedRows) rowToBoard
        matchedRows = M.keys matchedRowsMap
        (matchedRowsMap, unmatchedRows) = M.partition (== 0) $ M.mapWithKey (\k v -> if x `elem` k then v-1 else v) rowToMarked
findLastToWin (x:xs) rowToBoard list rowToMarked boards | boards == 1 && matchedRows /= [] = x * sum (board \\ (x:list))
  where board = rowToBoard ! head matchedRows
        matchedRows = M.keys $ M.filter (== 0) $ M.mapWithKey (\k v -> if x `elem` k then v-1 else v) rowToMarked
findLastToWin (x:xs) rowToBoard list rowToMarked boards = findLastToWin xs rowToBoard (x:list) unmatchedRows boards
  where (_, unmatchedRows) = M.partition (== 0) $ M.mapWithKey (\k v -> if x `elem` k then v-1 else v) rowToMarked

parseInput :: [String] -> ([Int], M.Map [Int] [Int])
parseInput [] = ([], empty)
parseInput (x:x2:xs) = (randomNumbers, boards)
    where randomNumbers = map (\s -> read s :: Int) $ wordsWhen (==',') x
          boards = parseBoards xs [] [] [] [] [] [] [] empty
parseInput (x:xs) = ([], empty)

parseBoards :: [String] -> [Int] -> [Int] -> [Int] -> [Int] -> [Int] -> [Int] -> [[Int]] -> M.Map [Int] [Int] -> M.Map [Int] [Int]
parseBoards [] a b c d e board lists m = m
parseBoards (x:xs) a b c d e board lists m | x == "" = parseBoards xs [] [] [] [] [] [] [] newMap
  where newMap = m `union` M.fromList (map (, board) ([sort a, sort b, sort c, sort d, sort e] ++ lists))
parseBoards (x:xs) a b c d e board lists m = parseBoards xs (head row:a) (row!!1:b) (row!!2:c) (row!!3:d) (row!!4:e) (row ++ board) (sort row:lists) m
  where row = map (\s -> read s :: Int) $ wordsWhen (==' ') x
