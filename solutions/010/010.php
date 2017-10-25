<?php

$limit = 2000000;
$sievebound = floor(($limit - 1) / 2);
$crosslimit = floor((sqrt($limit) - 1) / 2);
$sieve = array_fill(0, $sievebound, false);

for ($i = 1; $i < $crosslimit; $i++)
    if (!$sieve[$i])
        for ($j = 2 * $i * ($i + 1); $j < $sievebound; $j += 2 * $i + 1)
            $sieve[$j] = true;

$sum = 2;
for ($i = 1; $i < $sievebound; $i++)
    if (!$sieve[$i])
        $sum += 2 * $i + 1;

print($sum . "\n");
