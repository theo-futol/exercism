module Pangram (isPangram) where

import Data.Char (toLower, isLetter)
import Data.List (nub)

isPangram :: String -> Bool
isPangram text = 
  let letters = [toLower c | c <- text, isLetter c]
      uniqueLetters = nub letters
  in length uniqueLetters == 26
