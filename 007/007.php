<?php

$n = 10001;
$primes = array(2);
for ($x = 3; count($primes) < $n; $x += 2)
    for ($i = 0; $i < count($primes); $i++)
        if ($x % $primes[$i] === 0)
            break;
        else if ($primes[$i] * $primes[$i] > $x) {
            array_push($primes, $x);
            break;
        }
print($primes[count($primes) - 1] . "\n");
