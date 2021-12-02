{-# OPTIONS_GHC -Wno-incomplete-patterns #-}

type Command = (String, Int)
type Aim = Int
type HorizontalPosition = Int
type Depth = Int
type Position = (HorizontalPosition, Depth, Aim)

main = interact main'

main' :: String -> String
main' input = show $ multiply $ calculatePosition (0, 0, 0) commands
    where commands = inputToCommands input
          multiply (x, depth, _) = x*depth

inputToCommands :: String -> [Command]
inputToCommands input = map (toCommand . words) (lines input)
    where toCommand [a, b] = (a, read b :: Int)

calculatePosition :: Position -> [Command] -> Position
calculatePosition pos [] = pos
calculatePosition (x, depth, aim) (c@("forward", count):cs) = calculatePosition (x + count, depth + aim * count, aim) cs
calculatePosition (x, depth, aim) (c@("up", count):cs) = calculatePosition (x, depth, aim - count) cs
calculatePosition (x, depth, aim) (c@("down", count):cs) = calculatePosition (x, depth, aim + count) cs
