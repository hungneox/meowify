50,000 Meows _rewritten in go_
-----

![nyan-cat](https://media.giphy.com/media/VpnipAx5LVNTy/giphy.gif)

Inspired by [meow.py](https://github.com/hugovk/meow.py) project. It was an entry for [NaNoGenMo 2014](https://github.com/dariusk/NaNoGenMo-2014/)

> Spend the month of November writing code that generates a novel of 50k+ words.

The rules says:

> The "novel" is defined however you want. It could be 50,000 repetitions of the word "meow". It could literally grab a random novel from Project Gutenberg. It doesn't matter, as long as it's 50k+ words.


Build
-----

```
go build meowify.go
```

Usage
-----

```
./meowify -t -f input.txt > output.txt
```

# Sample

## Input

```
Roses are red,
Violets are blue,
Sugar is sweet,
And so are you.
```

## Ouput

```
Meeow meo meo,
Meeeoow meo meow,
Meeow me meeow,
Meo me meo meo.
```


