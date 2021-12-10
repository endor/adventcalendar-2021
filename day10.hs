import Data.Maybe (mapMaybe)
import Data.List (sort)

opening = ['{', '[', '<', '(']
closing = ['}', ']', '>', ')']

main = interact part2'

part2' :: String -> String
part2' input = show $ middleValue $ sort $ filter (/= 0) $ map (part2points . reverse . (`findIncompleteStack` [])) (lines input)

middleValue :: [Int] -> Int
middleValue values = values !! idx
  where idx = length values `div` 2

findIncompleteStack :: String -> [Char] -> [Char]
findIncompleteStack "" stack = stack
findIncompleteStack (x:xs) [] = findIncompleteStack xs [x]
findIncompleteStack (x:xs) stack | x `elem` opening = findIncompleteStack xs (x:stack)
findIncompleteStack (x:xs) (c:cs) | x `elem` closing && matches c x = findIncompleteStack xs cs
findIncompleteStack (x:xs) (c:cs) = []

part2points :: [Char] -> Int
part2points [] = 0
part2points (x:xs) | x == '(' = 1 + 5 * part2points xs
part2points (x:xs) | x == '[' = 2 + 5 * part2points xs
part2points (x:xs) | x == '{' = 3 + 5 * part2points xs
part2points (x:xs) | x == '<' = 4 + 5 * part2points xs
part2points _ = 0

part1' :: String -> String
part1' input = show $ sum $ map part1points $ mapMaybe (`findFirstIllegalCharacter` []) (lines input)

findFirstIllegalCharacter :: String -> [Char] -> Maybe Char
findFirstIllegalCharacter "" stack = Nothing
findFirstIllegalCharacter (x:xs) [] = findFirstIllegalCharacter xs [x]
findFirstIllegalCharacter (x:xs) stack | x `elem` opening = findFirstIllegalCharacter xs (x:stack)
findFirstIllegalCharacter (x:xs) (c:cs) | x `elem` closing && matches c x = findFirstIllegalCharacter xs cs
findFirstIllegalCharacter (x:xs) (c:cs) = Just x

matches :: Char -> Char -> Bool
matches '{' '}' = True
matches '[' ']' = True
matches '(' ')' = True
matches '<' '>' = True
matches _ _     = False

part1points :: Char -> Int
part1points ')' = 3
part1points ']' = 57
part1points '}' = 1197
part1points '>' = 25137
part1points _   = 0
