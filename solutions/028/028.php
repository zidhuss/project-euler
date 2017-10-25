<?php

function a($n) {
    return floor($n * ($n + 2) / 4) + floor(($n % 4) / 3) + 1;
}

$n = 1001;
$sum = 0;
for ($i = 1; a($i) <= $n*$n; $i++)
    $sum += a($i);
print($sum . "\n");
