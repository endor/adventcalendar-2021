import qualified Data.Map.Strict as M
import Data.Map (empty, insertWith, findWithDefault)
import Data.List.Split (splitOn)
import Data.Char (isLower)

main = interact main'

main' :: String -> String
main' input = show $ length $ findPathsPart2 "start" [] False [] $ parsePaths (lines input) empty

parsePaths :: [String] -> M.Map String [String] -> M.Map String [String]
parsePaths [] map = map
parsePaths (x:xs) map = parsePaths xs $ insertWith prepend src [dest] $ insertWith prepend dest [src] map
  where prepend d paths = head d:paths
        src = head parts
        dest = parts !! 1
        parts = splitOn "-" x

findPathsPart1 :: String -> [String] -> [[String]] -> M.Map String [String] -> [[String]]
findPathsPart1 src visited completed paths | src == "end" = (src:visited):completed
findPathsPart1 src visited completed paths | isLower (head src) && src `elem` visited = completed
findPathsPart1 src visited completed paths = concatMap (\d -> findPathsPart1 d (src:visited) completed paths) dests
  where dests = findWithDefault ["end"] src paths

findPathsPart2 :: String -> [String] -> Bool -> [[String]] -> M.Map String [String] -> [[String]]
findPathsPart2 src visited cannotVisitTwice completed paths | src == "end" = (src:visited):completed
findPathsPart2 src visited cannotVisitTwice completed paths | isLower (head src) && src `elem` visited && stop = completed
  where stop = cannotVisitTwice || src == "start" || src == "end"
findPathsPart2 src visited cannotVisitTwice completed paths = concatMap (\d -> findPathsPart2 d (src:visited) stop completed paths) dests
  where dests = findWithDefault ["end"] src paths
        stop = cannotVisitTwice || (isLower (head src) && src `elem` visited)
