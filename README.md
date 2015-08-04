# Markov headlines

This is a redis-backed Go webservice that can provide real and fake Hacker News headlines. 

It uses Markov chains to generate the fake ones.

### Redis store

Redis is used to store the corpus stories, the word-pairs and an associated set of follow-up words, and a set of generated titles.

The `generated_story:n` set contains a `fools` attribute that corresponds to the number of incorrect guesses in the corresponding website (forthcoming).

``` Redis
INCR next_corpus_id => 1000
HMSET corpus:1000 post_id 1234 title "Here is a sample title" score 235
LPUSH markov:word1/word2 "word"

HMSET generated_story:500 title "I am a fake title" fools 0
```
