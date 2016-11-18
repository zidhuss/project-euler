<?php

$sumSquare = 0;
$squareSum = 0;
for ($i = 0; $i <= 100; $i++) {
    $squareSum += $i * $i;
    $sumSquare += $i;
}

echo $sumSquare * $sumSquare - $squareSum . "\n";
