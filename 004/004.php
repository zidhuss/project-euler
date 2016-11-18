<?php

function is_palindrome($orig) {
    $rev = 0;
    $n = $orig;
    while ($n > 0) {
        $rev = $rev * 10 + $n % 10;
        $n = intdiv($n, 10);
    }
    return $orig === $rev;
}

$largest = 0;
$last = 999;
$first = 100;
for ($a = $last; $a >= $first; $a--)
    for ($b = $last; $b >= $first; $b--) {
        $calculated = $a * $b;
        if ($calculated > $largest && is_palindrome($calculated))
            $largest = $calculated;
    }
echo "$largest\n";
