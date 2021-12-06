import Debug.Trace (trace)

main = interact main'

wordsWhen     :: (Char -> Bool) -> String -> [String]
wordsWhen p s =  case dropWhile p s of
                      "" -> []
                      s' -> w : wordsWhen p s''
                            where (w, s'') = break p s'

main' :: String -> String
main' input = show $ countFish' 0 a b c d e f g h i
  where a = numTimesFound 0 list
        b = numTimesFound 1 list
        c = numTimesFound 2 list
        d = numTimesFound 3 list
        e = numTimesFound 4 list
        f = numTimesFound 5 list
        g = numTimesFound 6 list
        h = numTimesFound 7 list
        i = numTimesFound 8 list
        list = map (\s -> read s :: Int) (wordsWhen (==',') input)

countFish :: Int -> [Int] -> Int
countFish day fish | trace ("countFish " ++ show day ++ " " ++ show (length fish)) False = undefined
countFish 80 fish = length fish
countFish day fish = countFish (day+1) $ concatMap changeTimer fish

changeTimer :: Int -> [Int]
changeTimer 0 = [6, 8]
changeTimer n = [n-1]

countFish' :: Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int
countFish' day a b c d e f g h i | trace ("countFish " ++ show day) False = undefined
countFish' 256 a b c d e f g h i = a + b + c + d + e + f + g + h + i
countFish' day a b c d e f g h i = countFish' (day+1) b c d e f g (h+a) i a

numTimesFound :: Int -> [Int] -> Int
numTimesFound x = length . filter (== x)
